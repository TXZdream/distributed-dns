// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kademila.proto

package kademila

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_4de05f9272f955c8, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type FindNodesRequest struct {
	NodeID               string   `protobuf:"bytes,1,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	FromNodeID           string   `protobuf:"bytes,2,opt,name=FromNodeID,proto3" json:"FromNodeID,omitempty"`
	FromAccess           string   `protobuf:"bytes,3,opt,name=FromAccess,proto3" json:"FromAccess,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindNodesRequest) Reset()         { *m = FindNodesRequest{} }
func (m *FindNodesRequest) String() string { return proto.CompactTextString(m) }
func (*FindNodesRequest) ProtoMessage()    {}
func (*FindNodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4de05f9272f955c8, []int{1}
}

func (m *FindNodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNodesRequest.Unmarshal(m, b)
}
func (m *FindNodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNodesRequest.Marshal(b, m, deterministic)
}
func (m *FindNodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNodesRequest.Merge(m, src)
}
func (m *FindNodesRequest) XXX_Size() int {
	return xxx_messageInfo_FindNodesRequest.Size(m)
}
func (m *FindNodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindNodesRequest proto.InternalMessageInfo

func (m *FindNodesRequest) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *FindNodesRequest) GetFromNodeID() string {
	if m != nil {
		return m.FromNodeID
	}
	return ""
}

func (m *FindNodesRequest) GetFromAccess() string {
	if m != nil {
		return m.FromAccess
	}
	return ""
}

type FindNodesResponse struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=Nodes,proto3" json:"Nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindNodesResponse) Reset()         { *m = FindNodesResponse{} }
func (m *FindNodesResponse) String() string { return proto.CompactTextString(m) }
func (*FindNodesResponse) ProtoMessage()    {}
func (*FindNodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4de05f9272f955c8, []int{2}
}

func (m *FindNodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNodesResponse.Unmarshal(m, b)
}
func (m *FindNodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNodesResponse.Marshal(b, m, deterministic)
}
func (m *FindNodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNodesResponse.Merge(m, src)
}
func (m *FindNodesResponse) XXX_Size() int {
	return xxx_messageInfo_FindNodesResponse.Size(m)
}
func (m *FindNodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindNodesResponse proto.InternalMessageInfo

func (m *FindNodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Node struct {
	NodeID string `protobuf:"bytes,1,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	// Access 以IP:PORT的方式返回
	Access               string   `protobuf:"bytes,2,opt,name=Access,proto3" json:"Access,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_4de05f9272f955c8, []int{3}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *Node) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

type FindValueRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	FromNodeID           string   `protobuf:"bytes,2,opt,name=FromNodeID,proto3" json:"FromNodeID,omitempty"`
	FromAccess           string   `protobuf:"bytes,3,opt,name=FromAccess,proto3" json:"FromAccess,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindValueRequest) Reset()         { *m = FindValueRequest{} }
func (m *FindValueRequest) String() string { return proto.CompactTextString(m) }
func (*FindValueRequest) ProtoMessage()    {}
func (*FindValueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4de05f9272f955c8, []int{4}
}

func (m *FindValueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindValueRequest.Unmarshal(m, b)
}
func (m *FindValueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindValueRequest.Marshal(b, m, deterministic)
}
func (m *FindValueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindValueRequest.Merge(m, src)
}
func (m *FindValueRequest) XXX_Size() int {
	return xxx_messageInfo_FindValueRequest.Size(m)
}
func (m *FindValueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindValueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindValueRequest proto.InternalMessageInfo

func (m *FindValueRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *FindValueRequest) GetFromNodeID() string {
	if m != nil {
		return m.FromNodeID
	}
	return ""
}

func (m *FindValueRequest) GetFromAccess() string {
	if m != nil {
		return m.FromAccess
	}
	return ""
}

type FindValueResponse struct {
	Has                  bool     `protobuf:"varint,1,opt,name=Has,proto3" json:"Has,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Nodes                []*Node  `protobuf:"bytes,3,rep,name=Nodes,proto3" json:"Nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindValueResponse) Reset()         { *m = FindValueResponse{} }
func (m *FindValueResponse) String() string { return proto.CompactTextString(m) }
func (*FindValueResponse) ProtoMessage()    {}
func (*FindValueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4de05f9272f955c8, []int{5}
}

func (m *FindValueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindValueResponse.Unmarshal(m, b)
}
func (m *FindValueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindValueResponse.Marshal(b, m, deterministic)
}
func (m *FindValueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindValueResponse.Merge(m, src)
}
func (m *FindValueResponse) XXX_Size() int {
	return xxx_messageInfo_FindValueResponse.Size(m)
}
func (m *FindValueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindValueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindValueResponse proto.InternalMessageInfo

func (m *FindValueResponse) GetHas() bool {
	if m != nil {
		return m.Has
	}
	return false
}

func (m *FindValueResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *FindValueResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type StoreRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4de05f9272f955c8, []int{6}
}

func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (m *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(m, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

func (m *StoreRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StoreRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "kademila.Empty")
	proto.RegisterType((*FindNodesRequest)(nil), "kademila.FindNodesRequest")
	proto.RegisterType((*FindNodesResponse)(nil), "kademila.FindNodesResponse")
	proto.RegisterType((*Node)(nil), "kademila.Node")
	proto.RegisterType((*FindValueRequest)(nil), "kademila.FindValueRequest")
	proto.RegisterType((*FindValueResponse)(nil), "kademila.FindValueResponse")
	proto.RegisterType((*StoreRequest)(nil), "kademila.StoreRequest")
}

func init() { proto.RegisterFile("kademila.proto", fileDescriptor_4de05f9272f955c8) }

var fileDescriptor_4de05f9272f955c8 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x25, 0xa6, 0x89, 0xe9, 0x28, 0x35, 0x0e, 0x12, 0x42, 0x04, 0x29, 0x41, 0xb0, 0xa7, 0x22,
	0x15, 0x04, 0x8f, 0x62, 0x2d, 0x4a, 0x41, 0x24, 0x82, 0xf7, 0xd8, 0x2c, 0x12, 0x6d, 0xb2, 0x31,
	0x9b, 0x1e, 0xfa, 0xd7, 0x7e, 0x82, 0x64, 0x77, 0xd2, 0xc4, 0xa5, 0xf6, 0xe2, 0x6d, 0x66, 0xde,
	0xf2, 0xde, 0xbc, 0xb7, 0x03, 0x83, 0xcf, 0x38, 0x61, 0x59, 0xba, 0x8c, 0xc7, 0x45, 0xc9, 0x2b,
	0x8e, 0x4e, 0xd3, 0x87, 0xfb, 0x60, 0xdd, 0x67, 0x45, 0xb5, 0x0e, 0x3f, 0xc0, 0x9d, 0xa5, 0x79,
	0xf2, 0xc4, 0x13, 0x26, 0x22, 0xf6, 0xb5, 0x62, 0xa2, 0x42, 0x0f, 0xec, 0xba, 0x7f, 0x9c, 0xfa,
	0xc6, 0xd0, 0x18, 0xf5, 0x23, 0xea, 0xf0, 0x0c, 0x60, 0x56, 0xf2, 0x8c, 0xb0, 0x3d, 0x89, 0x75,
	0x26, 0x0d, 0x7e, 0xbb, 0x58, 0x30, 0x21, 0x7c, 0xb3, 0xc5, 0xd5, 0x24, 0xbc, 0x81, 0xe3, 0x8e,
	0x96, 0x28, 0x78, 0x2e, 0x18, 0x9e, 0x83, 0x25, 0x07, 0xbe, 0x31, 0x34, 0x47, 0x07, 0x93, 0xc1,
	0x78, 0xb3, 0x73, 0x3d, 0x8e, 0x14, 0x18, 0x5e, 0x43, 0xaf, 0x2e, 0xfe, 0x5c, 0xcd, 0x03, 0x9b,
	0x64, 0xd5, 0x5a, 0xd4, 0x85, 0x89, 0xb2, 0xf7, 0x1a, 0x2f, 0x57, 0xac, 0xb1, 0xe7, 0x82, 0x39,
	0x67, 0x6b, 0x22, 0xa8, 0xcb, 0x7f, 0x1b, 0x8b, 0x95, 0x31, 0x52, 0x21, 0x63, 0x2e, 0x98, 0x0f,
	0xb1, 0x90, 0x32, 0x4e, 0x54, 0x97, 0x78, 0x02, 0x96, 0x7c, 0x42, 0x0a, 0xaa, 0x69, 0x03, 0x30,
	0x77, 0x07, 0x70, 0xf8, 0x52, 0xf1, 0x72, 0x87, 0x89, 0xad, 0xec, 0x93, 0x6f, 0x03, 0x9c, 0x39,
	0x11, 0xe2, 0x05, 0xf4, 0x9e, 0xd3, 0xfc, 0x1d, 0x8f, 0x5a, 0x0d, 0x79, 0x05, 0x81, 0x26, 0x8a,
	0x77, 0xe0, 0x34, 0x3f, 0x85, 0x41, 0x8b, 0xe9, 0x97, 0x12, 0x9c, 0x6e, 0xc5, 0x28, 0x80, 0x29,
	0xf4, 0x37, 0xa9, 0xe8, 0x2c, 0xdd, 0x0f, 0xd1, 0x59, 0x7e, 0xc7, 0x78, 0x09, 0x96, 0x34, 0x8e,
	0x5e, 0xfb, 0xaa, 0x9b, 0x44, 0xa0, 0x9b, 0x79, 0xb3, 0xe5, 0xb1, 0x5f, 0xfd, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xe8, 0x4e, 0xaf, 0xa5, 0xfe, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KademilaClient is the client API for Kademila service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KademilaClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Node, error)
	// 返回离id更近的K个节点
	FindNode(ctx context.Context, in *FindNodesRequest, opts ...grpc.CallOption) (*FindNodesResponse, error)
	// 查询key值，若找到直接返回，找不到则返回最近的K个节点
	FindValue(ctx context.Context, in *FindValueRequest, opts ...grpc.CallOption) (*FindValueResponse, error)
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*Empty, error)
}

type kademilaClient struct {
	cc *grpc.ClientConn
}

func NewKademilaClient(cc *grpc.ClientConn) KademilaClient {
	return &kademilaClient{cc}
}

func (c *kademilaClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/kademila.Kademila/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kademilaClient) FindNode(ctx context.Context, in *FindNodesRequest, opts ...grpc.CallOption) (*FindNodesResponse, error) {
	out := new(FindNodesResponse)
	err := c.cc.Invoke(ctx, "/kademila.Kademila/FindNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kademilaClient) FindValue(ctx context.Context, in *FindValueRequest, opts ...grpc.CallOption) (*FindValueResponse, error) {
	out := new(FindValueResponse)
	err := c.cc.Invoke(ctx, "/kademila.Kademila/FindValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kademilaClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/kademila.Kademila/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KademilaServer is the server API for Kademila service.
type KademilaServer interface {
	Ping(context.Context, *Empty) (*Node, error)
	// 返回离id更近的K个节点
	FindNode(context.Context, *FindNodesRequest) (*FindNodesResponse, error)
	// 查询key值，若找到直接返回，找不到则返回最近的K个节点
	FindValue(context.Context, *FindValueRequest) (*FindValueResponse, error)
	Store(context.Context, *StoreRequest) (*Empty, error)
}

// UnimplementedKademilaServer can be embedded to have forward compatible implementations.
type UnimplementedKademilaServer struct {
}

func (*UnimplementedKademilaServer) Ping(ctx context.Context, req *Empty) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedKademilaServer) FindNode(ctx context.Context, req *FindNodesRequest) (*FindNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNode not implemented")
}
func (*UnimplementedKademilaServer) FindValue(ctx context.Context, req *FindValueRequest) (*FindValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindValue not implemented")
}
func (*UnimplementedKademilaServer) Store(ctx context.Context, req *StoreRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}

func RegisterKademilaServer(s *grpc.Server, srv KademilaServer) {
	s.RegisterService(&_Kademila_serviceDesc, srv)
}

func _Kademila_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KademilaServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kademila.Kademila/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KademilaServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kademila_FindNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KademilaServer).FindNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kademila.Kademila/FindNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KademilaServer).FindNode(ctx, req.(*FindNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kademila_FindValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KademilaServer).FindValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kademila.Kademila/FindValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KademilaServer).FindValue(ctx, req.(*FindValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kademila_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KademilaServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kademila.Kademila/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KademilaServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Kademila_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kademila.Kademila",
	HandlerType: (*KademilaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Kademila_Ping_Handler,
		},
		{
			MethodName: "FindNode",
			Handler:    _Kademila_FindNode_Handler,
		},
		{
			MethodName: "FindValue",
			Handler:    _Kademila_FindValue_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _Kademila_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kademila.proto",
}
