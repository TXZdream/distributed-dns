package dns

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	grpc "github.com/txzdream/distributed-dns/grpc"
	"github.com/txzdream/distributed-dns/logger"
)

// Ping 一个节点
func (d *DistributeDNS) Ping(ctx context.Context, req *grpc.Empty) (*grpc.Node, error) {
	logger.Logger.Sugar().Infow("收到Ping请求")
	return &grpc.Node{
		NodeID: toString(d.id),
		Access: d.access,
	}, nil
}

// FindNode 找到接收者离请求id更近的K个节点
func (d *DistributeDNS) FindNode(ctx context.Context, req *grpc.FindNodesRequest) (*grpc.FindNodesResponse, error) {
	logger.Logger.Sugar().Infow("收到查找节点请求",
		"target", ToBitArr(req.GetNodeID()).Bytes(),
		"from", ToBitArr(req.GetFromNodeID()).Bytes(),
		"fromAccess", req.GetFromAccess(),
	)
	d.RtLock.Lock()
	defer d.RtLock.Unlock()
	d.DataLock.Lock()
	defer d.DataLock.Unlock()
	// 被动添加请求节点到k桶中
	d.AddNode(ToBitArr(req.GetFromNodeID()), req.GetFromAccess())
	var ret grpc.FindNodesResponse
	nodes, err := d.GetNodes(req.GetNodeID())
	if err != nil {
		log.Println(err)
		return &grpc.FindNodesResponse{}, errors.New("服务器内部错误")
	}
	for k, v := range nodes {
		ret.Nodes = append(ret.Nodes, &grpc.Node{
			NodeID: k,
			Access: v,
		})
	}
	return &ret, nil
}

// FindValue 查询key值
func (d *DistributeDNS) FindValue(ctx context.Context, req *grpc.FindValueRequest) (*grpc.FindValueResponse, error) {
	logger.Logger.Sugar().Infow("收到查找值请求",
		"key", req.GetKey(),
		"from", ToBitArr(req.GetFromNodeID()).Bytes(),
		"fromAccess", req.GetFromAccess(),
	)
	d.RtLock.Lock()
	defer d.RtLock.Unlock()
	d.DataLock.Lock()
	defer d.DataLock.Unlock()
	// 被动添加请求节点到k桶中
	d.AddNode(ToBitArr(req.GetFromNodeID()), req.GetFromAccess())
	has, v := d.GetData(req.GetKey())
	ret := grpc.FindValueResponse{
		Has:   has,
		Value: v,
	}
	// 返回最近的K个节点
	if !has {
		keyID, err := CalculateHash(req.GetKey())
		if err != nil {
			log.Println(err)
			return &grpc.FindValueResponse{}, errors.New("服务器内部错误")
		}
		lcp, err := d.GetLCP(keyID)
		if err != nil {
			log.Println(err)
			return &grpc.FindValueResponse{}, errors.New("服务器内部错误")
		}
		for k, v := range d.routeTable[lcp] {
			ret.Nodes = append(ret.Nodes, &grpc.Node{
				NodeID: k,
				Access: v,
			})
		}
	}
	return &ret, nil
}

// Store 在该节点上存储数据
func (d *DistributeDNS) Store(ctx context.Context, req *grpc.StoreRequest) (*grpc.Empty, error) {
	logger.Logger.Sugar().Infow("收到存储请求",
		"key", req.GetKey(),
		"value", req.GetValue(),
	)
	d.RtLock.Lock()
	defer d.RtLock.Unlock()
	d.DataLock.Lock()
	defer d.DataLock.Unlock()
	// 如果value为空代表删除操作
	if req.GetValue() == "" {
		d.DeleteData(req.GetKey())
		return &grpc.Empty{}, nil
	}
	t := time.Now().Unix()
	s := strconv.FormatInt(t, 10)
	value := req.GetValue() + "@" + s
	// AddData根据时间戳来操作
	d.AddData(req.GetKey(), value)
	return &grpc.Empty{}, nil
}
