package dns

import (
	"context"
	kademila "github.com/txzdream/distributed-dns/grpc"
	"github.com/txzdream/distributed-dns/kademlia"
	"github.com/txzdream/distributed-dns/logger"
	"sync"

	"github.com/Workiva/go-datastructures/queue"
	"github.com/willf/bitset"
)

// DistributeDNS 实现了kademlia接口并提供了一个DNS服务
type DistributeDNS struct {
	// accessQueue 存储每个节点的最新访问时间
	accessQueue []queue.PriorityQueue
	// id 存储当前节点的ID值
	id *bitset.BitSet
	// access 表示当前节点的访问地址
	access string
	// k 表示桶的最大容量
	k uint16
	// routeTable 即为k桶
	routeTable []map[string]string
	RtLock     sync.Mutex
	// data 存储路由信息
	data     map[string]string
	DataLock sync.Mutex
}

var distributeDNS DistributeDNS
var once sync.Once

// Init 初始化当前节点，并通过已知信息加入/创建一个已知集群
// id代表当前节点的id号
// other代表另外一个已知集群中某个节点的访问地址，若为空，则不加入其他集群
func Init(k uint16, id *bitset.BitSet, access, other string) kademlia.Kademlia {
	once.Do(func() {
		distributeDNS.k = k
		distributeDNS.id = id
		for i := uint(0); i < distributeDNS.id.Len(); i++ {
			distributeDNS.accessQueue = append(distributeDNS.accessQueue, *queue.NewPriorityQueue(int(k), false))
		}
		distributeDNS.routeTable = make([]map[string]string, id.Len())
		for i := range distributeDNS.routeTable {
			distributeDNS.routeTable[i] = make(map[string]string)
		}
		distributeDNS.data = make(map[string]string)
		distributeDNS.access = access
		logger.Logger.Sugar().Infow("节点完成初始化",
			"id", id.Bytes(),
		)
		// 加入已知节点
		if other != "" {
			if client, err := dialGrpc(other); err != nil {
				logger.Logger.Sugar().Info("无法加入已知节点", err)
			} else {
				if res, err := client.Ping(context.Background(), &kademila.Empty{}); err != nil {
					logger.Logger.Sugar().Info("无法加入已知节点", err)
				} else {
					if err = distributeDNS.AddNode(ToBitArr(res.GetNodeID()), res.GetAccess()); err != nil {
						logger.Logger.Sugar().Info("无法加入已知节点", err)
					} else {
						if nodes, err := distributeDNS.GetNodesRecursive(toString(id)); err != nil {
							logger.Logger.Sugar().Info("无法加入已知节点", err)
						} else {
							for k, v := range nodes {
								distributeDNS.AddNode(ToBitArr(k), v)
							}
						}
					}
				}
			}
		}
	})
	return &distributeDNS
}
