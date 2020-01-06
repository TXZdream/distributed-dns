package dns

import (
	"context"

	grpc "distributed-dns/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Ping 一个节点
func (d DistributeDNS) Ping(ctx context.Context, req *grpc.Empty) (*grpc.Empty, error) {
	return &grpc.Empty{}, nil
}

// FindNode 找到接收者离请求id更近的K个节点
func (d DistributeDNS) FindNode(ctx context.Context, req *grpc.FindNodesRequest) (*grpc.FindNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNode not implemented")
}

// FindValue 查询key值
func (d DistributeDNS) FindValue(ctx context.Context, req *grpc.FindNodesRequest) (*grpc.Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindValue not implemented")
}

// Store 在该节点上存储数据
func (d DistributeDNS) Store(ctx context.Context, req *grpc.StoreRequest) (*grpc.Empty, error) {
	d.AddData(req.GetKey(), req.GetValue())
	return &grpc.Empty{}, nil
}
