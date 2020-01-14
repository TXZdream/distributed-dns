package dns

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	kademila "github.com/txzdream/distributed-dns/grpc"
	"github.com/txzdream/distributed-dns/logger"
	"github.com/willf/bitset"
)

// AddNode 添加一个节点到K桶中，data表示这个节点的访问方式
func (d *DistributeDNS) AddNode(id *bitset.BitSet, data string) error {
	logger.Logger.Sugar().Infow("添加节点",
		"id", id.Bytes(),
		"access", data,
	)
	if toString(id) == toString(d.id) {
		return nil
	}
	lcp, err := d.GetLCP(id)
	if err != nil {
		return err
	}
	strID := toString(id)
	// 更新accessQueue
	if d.accessQueue[lcp].Len() >= int(d.k) {
		items, err := d.accessQueue[lcp].Get(1)
		if err != nil {
			return err
		}
		// 若最早的节点无响应则删除
		client, err := dialGrpc(d.routeTable[lcp][items[0].(Item).ID])
		if err != nil {
			d.DeleteNode(items[0].(Item).ID)
		}
		_, err = client.Ping(context.Background(), &kademila.Empty{})
		if err != nil {
			d.DeleteNode(items[0].(Item).ID)
		}
		// 若有响应则忽视添加节点的请求
		d.accessQueue[lcp].Put(Item{
			ID:        items[0].(Item).ID,
			Timestamp: time.Now().Unix(),
		})
	} else {
		d.accessQueue[lcp].Put(Item{
			ID:        strID,
			Timestamp: time.Now().Unix(),
		})
		if d.routeTable[lcp] == nil {
			d.routeTable[lcp] = make(map[string]string)
		}
		d.routeTable[lcp][strID] = data
	}
	return nil
}

// DeleteNode 删除一个节点
func (d *DistributeDNS) DeleteNode(id string) error {
	logger.Logger.Sugar().Infow("删除节点",
		"id", ToBitArr(id).Bytes(),
	)
	lcp, err := d.GetLCP(ToBitArr(id))
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
func (d *DistributeDNS) GetLCP(target *bitset.BitSet) (uint8, error) {
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
func (d *DistributeDNS) AddData(key, value string) {
	logger.Logger.Sugar().Infow("向当前节点添加或更新值",
		"key", key,
		"value", value,
	)
	_, ok := d.data[key]
	if ok == false {
		d.data[key] = value
	} else {
		num := strings.Index(value, "@")
		reqTime := value[num+1 : len(value)]
		reqTimeInt64, _ := strconv.ParseInt(reqTime, 10, 64)

		num = strings.Index(d.data[key], "@")
		originTime := d.data[key][num+1 : len(value)]
		originTimeInt64, _ := strconv.ParseInt(originTime, 10, 64)
		if reqTimeInt64 > originTimeInt64 {
			d.data[key] = value
		} else {
			return
		}
	}
}

// DeleteData 删除当期节点的指定数据
func (d *DistributeDNS) DeleteData(key string) {
	logger.Logger.Sugar().Infow("删除当前节点指定数据",
		"key", key,
		"value", "",
	)
	_, ok := d.data[key]
	if ok == false {
		return
	}
	delete(d.data, key)
	return
}

// GetData 在集群中获取指定key的值
func (d *DistributeDNS) GetData(key string) (bool, string) {
	logger.Logger.Sugar().Infow("从当前节点获取值",
		"key", key,
	)
	value, ok := d.data[key]
	return ok, value
}

// GetNodes 返回当前节点的路由表中距离给定id最近的k个id
func (d *DistributeDNS) GetNodes(id string) (map[string]string, error) {
	lcp, err := d.GetLCP(ToBitArr(id))
	if err != nil {
		return nil, err
	}
	logger.Logger.Sugar().Infow("从当前节点获取与目标最接近的k个节点",
		"k", d.k,
		"target", ToBitArr(id).Bytes(),
		"lcp", lcp,
	)
	ret := make(map[string]string)
	around := 0
	if int(lcp) == int(d.id.Len()) {
		around = 1
	}
	for {
		if int(lcp)-around < 0 && int(lcp)+around >= int(d.id.Len()) || len(ret) >= int(d.k) {
			break
		}
		index := int(lcp) - around
		if index >= 0 {
			for k, v := range d.routeTable[index] {
				ret[k] = v
			}
		}
		index = int(lcp) + around
		if index < int(d.id.Len()) && around != 0 {
			for k, v := range d.routeTable[index] {
				ret[k] = v
			}
		}
		around++
	}
	return ret, nil
}

// Update 定期将自己的key-value对复制到其他的列表上
// 发出ping请求，将过期节点下线
func (d *DistributeDNS) Update() {
	logger.Logger.Sugar().Info("定时更新")
	d.RtLock.Lock()
	defer d.RtLock.Unlock()
	d.DataLock.Lock()
	defer d.DataLock.Unlock()
	// 找到过期节点
	for _, v := range d.routeTable {
		for k1, v1 := range v {
			client, err := dialGrpc(v1)
			if err != nil {
				d.DeleteNode(k1)
			}
			_, err = client.Ping(context.Background(), &kademila.Empty{})
			if err != nil {
				d.DeleteNode(k1)
			}
		}
	}
	// 将key-value对传播到其他节点
	for k, v := range d.data {
		hashKey, _ := CalculateHash(k)
		nodes, _ := d.GetNodesRecursive(toString(hashKey))
		for id, node := range nodes {
			client, err := dialGrpc(node)
			if err != nil {
				d.DeleteNode(node)
				continue
			}
			_, err = client.Store(context.Background(), &kademila.StoreRequest{
				Key:   k,
				Value: v,
			})
			if err != nil {
				d.DeleteNode(node)
			}
			logger.Logger.Sugar().Infow("传播键值对",
				"key", k,
				"value", v,
				"targetID", ToBitArr(id).Bytes(),
				"targetAccess", node,
			)
		}
	}
}

// GetDataRecursive 递归查询key所对应的value值
func (d *DistributeDNS) GetDataRecursive(key string) (bool, string, error) {
	logger.Logger.Sugar().Infow("从集群中获取值",
		"key", key,
	)
	if has, value := d.GetData(key); has {
		return true, value, nil
	}
	hashKey, err := CalculateHash(key)
	if err != nil {
		return false, "", err
	}
	strHashKey := toString(hashKey)
	nodes, err := d.GetNodesRecursive(strHashKey)
	if err != nil {
		return false, "", err
	}
	strID := toString(d.id)
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
func (d *DistributeDNS) GetNodesRecursive(id string) (map[string]string, error) {
	logger.Logger.Sugar().Infow("从集群中获取与目标最接近的k个节点",
		"k", d.k,
		"target", ToBitArr(id).Bytes(),
	)
	nodes, err := d.GetNodes(id)
	if err != nil {
		return nil, err
	}
	isVisited := make(map[string]bool)
	isVisited[toString(d.id)] = true
	strID := toString(d.id)
	for {
		isChanged := false
		raw := make(map[string]string)
		for k, v := range nodes {
			raw[k] = v
		}
		// 对本地获得的k个最近的节点，分别发送请求以获取值
		for k, v := range raw {
			if isVisited[k] {
				continue
			}
			client, err := dialGrpc(v)
			if err != nil {
				d.DeleteNode(k)
				continue
			}
			res, err := client.FindNode(context.Background(), &kademila.FindNodesRequest{
				NodeID:     id,
				FromNodeID: strID,
				FromAccess: d.access,
			})
			if err != nil {
				d.DeleteNode(k)
				continue
			}
			for _, node := range res.GetNodes() {
				nodes[node.GetNodeID()] = node.GetAccess()
			}
			isVisited[k] = true
			// 筛选出逻辑距离最小的前k个节点
			if len(nodes) <= int(d.k) {
				if len(nodes) > len(raw) {
					isChanged = true
				}
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
