package dns

import (
	"context"
	"errors"
	"log"

	grpc "distributed-dns/grpc"
)

// Ping 一个节点
func (d DistributeDNS) Ping(ctx context.Context, req *grpc.Empty) (*grpc.Empty, error) {
	return &grpc.Empty{}, nil
}

// FindNode 找到接收者离请求id更近的K个节点
func (d DistributeDNS) FindNode(ctx context.Context, req *grpc.FindNodesRequest) (*grpc.FindNodesResponse, error) {
	// 被动添加请求节点到k桶中
	d.AddNode(toBitArr(req.GetFromNodeID()), req.GetFromAccess())
	var ret grpc.FindNodesResponse
	lcp, err := d.GetLCP(toBitArr(req.GetNodeID()))
	if err != nil {
		log.Println(err)
		return &ret, errors.New("服务器内部错误")
	}
	for k, v := range d.routeTable[lcp] {
		ret.Nodes = append(ret.Nodes, &grpc.Node{
			NodeID: k,
			Access: v,
		})
	}
	return &ret, nil
}

// FindValue 查询key值
func (d DistributeDNS) FindValue(ctx context.Context, req *grpc.FindValueRequest) (*grpc.FindValueResponse, error) {
	// 被动添加请求节点到k桶中
	d.AddNode(toBitArr(req.GetFromNodeID()), req.GetFromAccess())
	has, v := d.GetData(req.GetKey())
	ret := grpc.FindValueResponse{
		Has:   has,
		Value: v,
	}
	// 返回最近的K个节点
	if !has {
		keyID, err := calculateHash(req.GetKey())
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
func (d DistributeDNS) Store(ctx context.Context, req *grpc.StoreRequest) (*grpc.Empty, error) {
	d.AddData(req.GetKey(), req.GetValue())
	return &grpc.Empty{}, nil
}
