/**
 * Copyright 2019 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nodeagent.proto

package nodeagent

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

type Node struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Iqns                 []string `protobuf:"bytes,2,rep,name=iqns,proto3" json:"iqns,omitempty"`
	Wwpns                []string `protobuf:"bytes,3,rep,name=wwpns,proto3" json:"wwpns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5331ae4b8115762, []int{0}
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

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetIqns() []string {
	if m != nil {
		return m.Iqns
	}
	return nil
}

func (m *Node) GetWwpns() []string {
	if m != nil {
		return m.Wwpns
	}
	return nil
}

type GetNodeInfoRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNodeInfoRequest) Reset()         { *m = GetNodeInfoRequest{} }
func (m *GetNodeInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetNodeInfoRequest) ProtoMessage()    {}
func (*GetNodeInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5331ae4b8115762, []int{1}
}

func (m *GetNodeInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNodeInfoRequest.Unmarshal(m, b)
}
func (m *GetNodeInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNodeInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetNodeInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeInfoRequest.Merge(m, src)
}
func (m *GetNodeInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetNodeInfoRequest.Size(m)
}
func (m *GetNodeInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeInfoRequest proto.InternalMessageInfo

func (m *GetNodeInfoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetNodeInfoReply struct {
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNodeInfoReply) Reset()         { *m = GetNodeInfoReply{} }
func (m *GetNodeInfoReply) String() string { return proto.CompactTextString(m) }
func (*GetNodeInfoReply) ProtoMessage()    {}
func (*GetNodeInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5331ae4b8115762, []int{2}
}

func (m *GetNodeInfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNodeInfoReply.Unmarshal(m, b)
}
func (m *GetNodeInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNodeInfoReply.Marshal(b, m, deterministic)
}
func (m *GetNodeInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeInfoReply.Merge(m, src)
}
func (m *GetNodeInfoReply) XXX_Size() int {
	return xxx_messageInfo_GetNodeInfoReply.Size(m)
}
func (m *GetNodeInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeInfoReply proto.InternalMessageInfo

func (m *GetNodeInfoReply) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type IscsiLoginRequest struct {
	Targets              []string `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IscsiLoginRequest) Reset()         { *m = IscsiLoginRequest{} }
func (m *IscsiLoginRequest) String() string { return proto.CompactTextString(m) }
func (*IscsiLoginRequest) ProtoMessage()    {}
func (*IscsiLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5331ae4b8115762, []int{3}
}

func (m *IscsiLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IscsiLoginRequest.Unmarshal(m, b)
}
func (m *IscsiLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IscsiLoginRequest.Marshal(b, m, deterministic)
}
func (m *IscsiLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IscsiLoginRequest.Merge(m, src)
}
func (m *IscsiLoginRequest) XXX_Size() int {
	return xxx_messageInfo_IscsiLoginRequest.Size(m)
}
func (m *IscsiLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IscsiLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IscsiLoginRequest proto.InternalMessageInfo

func (m *IscsiLoginRequest) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type IscsiLoginReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IscsiLoginReply) Reset()         { *m = IscsiLoginReply{} }
func (m *IscsiLoginReply) String() string { return proto.CompactTextString(m) }
func (*IscsiLoginReply) ProtoMessage()    {}
func (*IscsiLoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5331ae4b8115762, []int{4}
}

func (m *IscsiLoginReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IscsiLoginReply.Unmarshal(m, b)
}
func (m *IscsiLoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IscsiLoginReply.Marshal(b, m, deterministic)
}
func (m *IscsiLoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IscsiLoginReply.Merge(m, src)
}
func (m *IscsiLoginReply) XXX_Size() int {
	return xxx_messageInfo_IscsiLoginReply.Size(m)
}
func (m *IscsiLoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_IscsiLoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_IscsiLoginReply proto.InternalMessageInfo

type IscsiLogoutRequest struct {
	Targets              []string `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IscsiLogoutRequest) Reset()         { *m = IscsiLogoutRequest{} }
func (m *IscsiLogoutRequest) String() string { return proto.CompactTextString(m) }
func (*IscsiLogoutRequest) ProtoMessage()    {}
func (*IscsiLogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5331ae4b8115762, []int{5}
}

func (m *IscsiLogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IscsiLogoutRequest.Unmarshal(m, b)
}
func (m *IscsiLogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IscsiLogoutRequest.Marshal(b, m, deterministic)
}
func (m *IscsiLogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IscsiLogoutRequest.Merge(m, src)
}
func (m *IscsiLogoutRequest) XXX_Size() int {
	return xxx_messageInfo_IscsiLogoutRequest.Size(m)
}
func (m *IscsiLogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IscsiLogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IscsiLogoutRequest proto.InternalMessageInfo

func (m *IscsiLogoutRequest) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type IscsiLogoutReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IscsiLogoutReply) Reset()         { *m = IscsiLogoutReply{} }
func (m *IscsiLogoutReply) String() string { return proto.CompactTextString(m) }
func (*IscsiLogoutReply) ProtoMessage()    {}
func (*IscsiLogoutReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5331ae4b8115762, []int{6}
}

func (m *IscsiLogoutReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IscsiLogoutReply.Unmarshal(m, b)
}
func (m *IscsiLogoutReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IscsiLogoutReply.Marshal(b, m, deterministic)
}
func (m *IscsiLogoutReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IscsiLogoutReply.Merge(m, src)
}
func (m *IscsiLogoutReply) XXX_Size() int {
	return xxx_messageInfo_IscsiLogoutReply.Size(m)
}
func (m *IscsiLogoutReply) XXX_DiscardUnknown() {
	xxx_messageInfo_IscsiLogoutReply.DiscardUnknown(m)
}

var xxx_messageInfo_IscsiLogoutReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Node)(nil), "nodeagent.Node")
	proto.RegisterType((*GetNodeInfoRequest)(nil), "nodeagent.GetNodeInfoRequest")
	proto.RegisterType((*GetNodeInfoReply)(nil), "nodeagent.GetNodeInfoReply")
	proto.RegisterType((*IscsiLoginRequest)(nil), "nodeagent.IscsiLoginRequest")
	proto.RegisterType((*IscsiLoginReply)(nil), "nodeagent.IscsiLoginReply")
	proto.RegisterType((*IscsiLogoutRequest)(nil), "nodeagent.IscsiLogoutRequest")
	proto.RegisterType((*IscsiLogoutReply)(nil), "nodeagent.IscsiLogoutReply")
}

func init() { proto.RegisterFile("nodeagent.proto", fileDescriptor_b5331ae4b8115762) }

var fileDescriptor_b5331ae4b8115762 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x1b, 0x1b, 0x95, 0x4c, 0x0e, 0x69, 0x07, 0x0f, 0x4b, 0x54, 0x28, 0xeb, 0xa5, 0x17,
	0x73, 0xa8, 0x07, 0xcf, 0x82, 0xa0, 0xa5, 0xe2, 0x21, 0x6f, 0x10, 0xed, 0x1a, 0x02, 0x71, 0x77,
	0xdb, 0xdd, 0x52, 0xf2, 0xd4, 0xbe, 0x82, 0xcc, 0x96, 0xd8, 0x8d, 0x31, 0x78, 0x9b, 0xf9, 0xe7,
	0xe7, 0xdf, 0x6f, 0x86, 0x85, 0x44, 0xaa, 0xb5, 0x28, 0x4a, 0x21, 0x6d, 0xa6, 0xb7, 0xca, 0x2a,
	0x8c, 0x7e, 0x04, 0xfe, 0x08, 0xe1, 0xab, 0x5a, 0x0b, 0x44, 0x08, 0x65, 0xf1, 0x29, 0x58, 0x30,
	0x0b, 0xe6, 0x51, 0xee, 0x6a, 0xd2, 0xaa, 0x8d, 0x34, 0xec, 0x64, 0x36, 0x26, 0x8d, 0x6a, 0xbc,
	0x80, 0xd3, 0xfd, 0x5e, 0x4b, 0xc3, 0xc6, 0x4e, 0x3c, 0x34, 0x7c, 0x0e, 0xf8, 0x24, 0x2c, 0x05,
	0x2d, 0xe5, 0x87, 0xca, 0xc5, 0x66, 0x27, 0x8c, 0xfd, 0x2b, 0x93, 0xdf, 0xc3, 0xa4, 0xe3, 0xd4,
	0x75, 0x83, 0x37, 0x10, 0x12, 0x90, 0xf3, 0xc5, 0x8b, 0x24, 0x3b, 0xe2, 0x92, 0x2f, 0x77, 0x43,
	0x7e, 0x0b, 0xd3, 0xa5, 0x79, 0x37, 0xd5, 0x8b, 0x2a, 0x2b, 0xd9, 0xbe, 0xc0, 0xe0, 0xdc, 0x16,
	0xdb, 0x52, 0x58, 0xc3, 0x02, 0xc7, 0xd3, 0xb6, 0x7c, 0x0a, 0x89, 0x6f, 0xd7, 0x75, 0xc3, 0x33,
	0xc0, 0x56, 0x52, 0x3b, 0xfb, 0x7f, 0x04, 0xc2, 0xa4, 0xe3, 0xd7, 0x75, 0xb3, 0xf8, 0x0a, 0x20,
	0x22, 0xa8, 0x07, 0xc2, 0xc3, 0x15, 0xc4, 0xde, 0x32, 0x78, 0xed, 0x91, 0xf7, 0xcf, 0x91, 0x5e,
	0x0e, 0x8d, 0x09, 0x6e, 0x84, 0xcf, 0x00, 0x47, 0x62, 0xbc, 0xf2, 0xcc, 0xbd, 0xbd, 0xd3, 0x74,
	0x60, 0x7a, 0x48, 0x5a, 0x41, 0xec, 0x81, 0x77, 0xb0, 0xfa, 0x07, 0xe8, 0x60, 0xfd, 0xde, 0x97,
	0x8f, 0xde, 0xce, 0xdc, 0x97, 0xb9, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x00, 0x30, 0x9e, 0x45,
	0x45, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeAgentClient is the client API for NodeAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeAgentClient interface {
	GetNodeInfo(ctx context.Context, in *GetNodeInfoRequest, opts ...grpc.CallOption) (*GetNodeInfoReply, error)
	IscsiLogin(ctx context.Context, in *IscsiLoginRequest, opts ...grpc.CallOption) (*IscsiLoginReply, error)
	IscsiLogout(ctx context.Context, in *IscsiLogoutRequest, opts ...grpc.CallOption) (*IscsiLogoutReply, error)
}

type nodeAgentClient struct {
	cc *grpc.ClientConn
}

func NewNodeAgentClient(cc *grpc.ClientConn) NodeAgentClient {
	return &nodeAgentClient{cc}
}

func (c *nodeAgentClient) GetNodeInfo(ctx context.Context, in *GetNodeInfoRequest, opts ...grpc.CallOption) (*GetNodeInfoReply, error) {
	out := new(GetNodeInfoReply)
	err := c.cc.Invoke(ctx, "/nodeagent.NodeAgent/GetNodeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAgentClient) IscsiLogin(ctx context.Context, in *IscsiLoginRequest, opts ...grpc.CallOption) (*IscsiLoginReply, error) {
	out := new(IscsiLoginReply)
	err := c.cc.Invoke(ctx, "/nodeagent.NodeAgent/IscsiLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAgentClient) IscsiLogout(ctx context.Context, in *IscsiLogoutRequest, opts ...grpc.CallOption) (*IscsiLogoutReply, error) {
	out := new(IscsiLogoutReply)
	err := c.cc.Invoke(ctx, "/nodeagent.NodeAgent/IscsiLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeAgentServer is the server API for NodeAgent service.
type NodeAgentServer interface {
	GetNodeInfo(context.Context, *GetNodeInfoRequest) (*GetNodeInfoReply, error)
	IscsiLogin(context.Context, *IscsiLoginRequest) (*IscsiLoginReply, error)
	IscsiLogout(context.Context, *IscsiLogoutRequest) (*IscsiLogoutReply, error)
}

// UnimplementedNodeAgentServer can be embedded to have forward compatible implementations.
type UnimplementedNodeAgentServer struct {
}

func (*UnimplementedNodeAgentServer) GetNodeInfo(ctx context.Context, req *GetNodeInfoRequest) (*GetNodeInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeInfo not implemented")
}
func (*UnimplementedNodeAgentServer) IscsiLogin(ctx context.Context, req *IscsiLoginRequest) (*IscsiLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IscsiLogin not implemented")
}
func (*UnimplementedNodeAgentServer) IscsiLogout(ctx context.Context, req *IscsiLogoutRequest) (*IscsiLogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IscsiLogout not implemented")
}

func RegisterNodeAgentServer(s *grpc.Server, srv NodeAgentServer) {
	s.RegisterService(&_NodeAgent_serviceDesc, srv)
}

func _NodeAgent_GetNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAgentServer).GetNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeagent.NodeAgent/GetNodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAgentServer).GetNodeInfo(ctx, req.(*GetNodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAgent_IscsiLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IscsiLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAgentServer).IscsiLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeagent.NodeAgent/IscsiLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAgentServer).IscsiLogin(ctx, req.(*IscsiLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAgent_IscsiLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IscsiLogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAgentServer).IscsiLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeagent.NodeAgent/IscsiLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAgentServer).IscsiLogout(ctx, req.(*IscsiLogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodeagent.NodeAgent",
	HandlerType: (*NodeAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeInfo",
			Handler:    _NodeAgent_GetNodeInfo_Handler,
		},
		{
			MethodName: "IscsiLogin",
			Handler:    _NodeAgent_IscsiLogin_Handler,
		},
		{
			MethodName: "IscsiLogout",
			Handler:    _NodeAgent_IscsiLogout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodeagent.proto",
}
