package kademlia

import "github.com/willf/bitset"

// Kademlia 代表一个实现了Kademlia算法所必须的函数
type Kademlia interface {
	// AddNode 添加一个节点到K桶中，data表示这个节点的访问方式
	AddNode(id *bitset.BitSet, data string) error
	// DeleteNode 删除一个节点
	DeleteNode(id string) error
	// GetLCP 获取当前节点与目标节点的最长公共前缀
	GetLCP(target *bitset.BitSet) (uint8, error)
	// AddData 添加一组数据到当前节点
	AddData(key, value string)
	// GetData 在集群中获取指定key的值
	GetData(key string) (bool, string)
	// GetNodes 返回当前节点的路由表中距离给定id最近的k个id
	GetNodes(id string) (map[string]string, error)
	// Update 定期将自己的key-value对复制到其他的列表上
	Update()
	// GetDataRecursive 递归查询key所对应的value值
	GetDataRecursive(key string) (bool, string, error)
	// GetNodesRecursive 递归获取距离给定id最近的k个节点
	GetNodesRecursive(id string) (map[string]string, error)
}
