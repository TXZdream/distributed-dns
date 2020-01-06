package kademlia

import "github.com/Workiva/go-datastructures/bitarray"

// Kademlia 代表一个实现了Kademlia算法所必须的函数
type Kademlia interface {
	// Init 初始化当前节点
	Init(k uint16, id bitarray.BitArray) error
	// AddNode 添加一个节点到K桶中，data表示这个节点的访问方式
	AddNode(id bitarray.BitArray, data string) error
	// GetLCP 获取当前节点与目标节点的最长公共前缀
	GetLCP(target bitarray.BitArray) (uint8, error)
	// AddData 添加一组数据到当前节点
	AddData(key, value string)
	// GetData 在集群中获取指定key的值
	GetData(key string) (bool, string)
}
