package dns

import (
	"github.com/Workiva/go-datastructures/queue"
	"github.com/willf/bitset"
)

// DistributeDNS 实现了kademlia接口并提供了一个DNS服务
type DistributeDNS struct {
	// accessQueue 存储每个节点的最新访问时间
	accessQueue []queue.PriorityQueue
	// id 存储当前节点的ID值
	id *bitset.BitSet
	// k 表示桶的最大容量
	k uint16
	// routeTable 即为k桶
	routeTable []map[string]string
	// data 存储路由信息
	data map[string]string
}
