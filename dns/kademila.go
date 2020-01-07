package dns

import (
	"context"
	kademila "distributed-dns/grpc"
	"errors"
	"time"

	"github.com/willf/bitset"
)

// AddNode 添加一个节点到K桶中，data表示这个节点的访问方式
func (d DistributeDNS) AddNode(id *bitset.BitSet, data string) error {
	lcp, err := d.GetLCP(id)
	if err != nil {
		return err
	}
	strID, err := toString(id)
	if err != nil {
		return err
	}
	// 更新accessQueue
	var remove string
	if d.accessQueue[lcp].Len() >= int(d.k) {
		items, err := d.accessQueue[lcp].Get(1)
		if err != nil {
			return err
		}
		remove = items[0].(Item).ID
	}
	d.accessQueue[lcp].Put(Item{
		ID:        strID,
		Timestamp: time.Now().Unix(),
	})
	// 添加路由表项
	if remove != "" {
		removedID := toBitArr(remove)
		if removeLCP, err := d.GetLCP(removedID); err != nil {
			return err
		} else {
			delete(d.routeTable[removeLCP], remove)
		}
	}
	if d.routeTable[lcp] == nil {
		d.routeTable[lcp] = make(map[string]string)
	}
	d.routeTable[lcp][strID] = data
	return nil
}

// DeleteNode 删除一个节点
func (d DistributeDNS) DeleteNode(id string) error {
	lcp, err := d.GetLCP(toBitArr(id))
	if err != nil {
		return err
	}
	err = d.accessQueue[lcp].Put(Item{
		ID:        id,
		Timestamp: -1,
	})
	if err != nil {
		return err
	}
	_, err = d.accessQueue[lcp].Get(1)
	if err != nil {
		return err
	}
	delete(d.routeTable[lcp], id)
	return nil
}

// GetLCP 获取a与b之间的最长公共前缀值
func (d DistributeDNS) GetLCP(target *bitset.BitSet) (uint8, error) {
	if d.id.Len() != target.Len() {
		return 0, errors.New("a和b的长度不一致")
	}
	cnt := uint8(0)
	for i := uint(0); i < d.id.Len(); i++ {
		aval := d.id.Test(i)
		bval := target.Test(i)
		if aval == bval {
			cnt++
		} else {
			break
		}
	}
	return cnt, nil
}

// AddData 添加一组数据到当前节点
func (d DistributeDNS) AddData(key, value string) {
	d.data[key] = value
}

// GetData 在集群中获取指定key的值
func (d DistributeDNS) GetData(key string) (bool, string) {
	value, ok := d.data[key]
	return ok, value
}

// GetNodes 返回当前节点的路由表中距离给定id最近的k个id
func (d DistributeDNS) GetNodes(id string) (map[string]string, error) {
	lcp, err := d.GetLCP(toBitArr(id))
	if err != nil {
		return nil, err
	}
	ret := make(map[string]string)
	index := lcp
	for {
		if index < 0 || index >= uint8(d.id.Len()) || len(ret) >= int(d.k) {
			break
		}
		count := int(d.k) - len(ret)
		for k, v := range d.routeTable[lcp] {
			ret[k] = v
			if count--; count == 0 {
				break
			}
		}
		// 计算下一个index的值
		if index > lcp {
			if 2*lcp-index < 0 {
				index++
			} else {
				index = 2*lcp - index
			}
		} else if index == lcp {
			if lcp == 0 {
				index++
			} else {
				index--
			}
		} else {
			if 2*lcp-index >= uint8(d.id.Len()) {
				index--
			} else {
				index = 2*lcp - index
			}
		}
	}
	return ret, nil
}

// Update 定期将自己的key-value对复制到其他的列表上
// 发出ping请求，将过期节点下线
func (d DistributeDNS) Update() {
}

// GetDataRecursive 递归查询key所对应的value值
func (d DistributeDNS) GetDataRecursive(key string) (bool, string, error) {
	if has, value := d.GetData(key); has {
		return true, value, nil
	}
	hashKey, err := calculateHash(key)
	if err != nil {
		return false, "", err
	}
	strHashKey, err := toString(hashKey)
	if err != nil {
		return false, "", err
	}
	nodes, err := d.GetNodesRecursive(strHashKey)
	if err != nil {
		return false, "", err
	}
	strID, err := toString(d.id)
	if err != nil {
		return false, "", err
	}
	// 从nodes中查找key
	for _, v := range nodes {
		client, err := dialGrpc(v)
		if err != nil {
			return false, "", err
		}
		res, _ := client.FindValue(context.Background(), &kademila.FindValueRequest{
			Key:        key,
			FromNodeID: strID,
			FromAccess: d.access,
		})
		if res.GetHas() {
			return true, res.GetValue(), nil
		}
	}
	return false, "", nil
}

// GetNodesRecursive 递归获取距离给定id最近的k个节点
func (d DistributeDNS) GetNodesRecursive(id string) (map[string]string, error) {
	nodes, err := d.GetNodes(id)
	if err != nil {
		return nil, err
	}
	isVisited := make(map[string]bool)
	strID, err := toString(d.id)
	if err != nil {
		return nil, err
	}
	for {
		isChanged := false
		raw := make(map[string]string)
		for k, v := range nodes {
			raw[k] = v
		}
		for k, v := range nodes {
			if isVisited[k] {
				continue
			}
			client, err := dialGrpc(v)
			if err != nil {
				d.DeleteNode(k)
			}
			res, err := client.FindNode(context.Background(), &kademila.FindNodesRequest{
				NodeID:     id,
				FromNodeID: strID,
				FromAccess: d.access,
			})
			if err != nil {
				d.DeleteNode(k)
			}
			for _, node := range res.GetNodes() {
				nodes[node.GetNodeID()] = node.GetAccess()
			}
			isVisited[k] = true
			// 筛选出最近的前k个节点
			if len(nodes) <= int(d.k) {
				continue
			}
			keys := make([]string, len(nodes))
			for k := range nodes {
				keys = append(keys, k)
			}
			for i := 0; i < int(d.k-1); i++ {
				for j := i + 1; j < len(keys); j++ {
					cmp := compare(id, keys[i], keys[j])
					if cmp > 0 {
						keys[i], keys[j] = keys[j], keys[i]
					}
				}
			}
			// 删除超过k长度的其他节点
			for i := int(d.k); i < len(keys); i++ {
				delete(nodes, keys[i])
			}
			for i := 0; i < len(keys); i++ {
				if _, has := raw[keys[i]]; !has {
					isChanged = true
					break
				}
			}
		}
		if !isChanged {
			break
		}
	}
	return nodes, err
}
