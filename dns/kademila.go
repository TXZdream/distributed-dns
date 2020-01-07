package dns

import (
	"errors"
	"time"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/willf/bitset"
)

// Init 初始化当前节点，并通过已知信息加入/创建一个已知集群
// id代表当前节点的id号
// other代表另外一个已知集群中某个节点的访问地址，若为空，则不加入其他集群
func Init(k uint16, id *bitset.BitSet, other string) (d DistributeDNS) {
	d.k = k
	d.id = id
	for i := uint(0); i < id.Len(); i++ {
		d.accessQueue = append(d.accessQueue, *queue.NewPriorityQueue(int(k), false))
	}
	d.routeTable = make([]map[string]string, k)
	d.data = make(map[string]string)
	// 加入已知节点
	if other != "" {

	}
	return
}

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

// Update 定期将自己的key-value对复制到其他的列表上
// 发出ping请求，将过期节点下线
func (d DistributeDNS) Update() {
}
