package dns

import (
	"errors"
	"time"

	"github.com/Workiva/go-datastructures/bitarray"
	"github.com/Workiva/go-datastructures/queue"
)

// Init 初始化当前节点
func (d DistributeDNS) Init(k uint16, id bitarray.BitArray) error {
	d.k = k
	d.id = id
	d.accessQueue = *queue.NewPriorityQueue(int(k), false)
	d.routeTable = make([]map[string]string, k)
	d.data = make(map[string]string)
	return nil
}

// AddNode 添加一个节点到K桶中，data表示这个节点的访问方式
func (d DistributeDNS) AddNode(id bitarray.BitArray, data string) error {
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
	if d.accessQueue.Len() >= int(d.k) {
		items, err := d.accessQueue.Get(1)
		if err != nil {
			return err
		}
		remove = items[0].(Item).ID
	}
	d.accessQueue.Put(Item{
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
func (d DistributeDNS) GetLCP(target bitarray.BitArray) (uint8, error) {
	if d.id.Capacity() != target.Capacity() {
		return 0, errors.New("a和b的长度不一致")
	}
	cnt := uint8(0)
	for i := uint64(0); i < d.id.Capacity(); i++ {
		aval, err := d.id.GetBit(i)
		if err != nil {
			return 0, errors.New("获取bitset内容失败")
		}
		bval, err := target.GetBit(i)
		if err != nil {
			return 0, errors.New("获取bitset内容失败")
		}
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
