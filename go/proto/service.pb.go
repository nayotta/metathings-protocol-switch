// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package _switch

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type StateResponse struct {
	State                string   `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateResponse) Reset()         { *m = StateResponse{} }
func (m *StateResponse) String() string { return proto.CompactTextString(m) }
func (*StateResponse) ProtoMessage()    {}
func (*StateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *StateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateResponse.Unmarshal(m, b)
}
func (m *StateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateResponse.Marshal(b, m, deterministic)
}
func (m *StateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateResponse.Merge(m, src)
}
func (m *StateResponse) XXX_Size() int {
	return xxx_messageInfo_StateResponse.Size(m)
}
func (m *StateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StateResponse proto.InternalMessageInfo

func (m *StateResponse) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*StateResponse)(nil), "ai.metathings.protocol.service.switch.StateResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x4d, 0xcc, 0xd4, 0xcb, 0x4d,
	0x2d, 0x49, 0x2c, 0xc9, 0xc8, 0xcc, 0x4b, 0x2f, 0x86, 0x08, 0x26, 0xe7, 0xe7, 0xe8, 0xc1, 0x54,
	0x15, 0x97, 0x67, 0x96, 0x24, 0x67, 0x48, 0x49, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83,
	0xe5, 0x93, 0x4a, 0xd3, 0xf4, 0x53, 0x73, 0x0b, 0x4a, 0x2a, 0x21, 0xca, 0x95, 0x54, 0xb9, 0x78,
	0x83, 0x4b, 0x12, 0x4b, 0x52, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x44, 0xb8,
	0x58, 0x8b, 0x41, 0x02, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0xd1, 0x4d, 0x46,
	0x2e, 0xde, 0x60, 0xb0, 0x71, 0xc1, 0x10, 0xc3, 0x85, 0xcc, 0xb8, 0x98, 0xfc, 0xf3, 0x84, 0xc4,
	0xf4, 0x20, 0x86, 0xeb, 0xc1, 0x0c, 0xd7, 0x73, 0x05, 0x19, 0x2e, 0x85, 0x43, 0x5c, 0x89, 0x41,
	0xc8, 0x9c, 0x8b, 0xd9, 0x3f, 0x2d, 0x8d, 0x0c, 0x8d, 0xe1, 0x5c, 0xac, 0x60, 0x97, 0xe2, 0xd4,
	0x6a, 0xa2, 0x47, 0x54, 0x78, 0xe8, 0xa1, 0xf8, 0x57, 0x89, 0xc1, 0x89, 0x23, 0x8a, 0x0d, 0x22,
	0x93, 0xc4, 0x06, 0xd6, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x93, 0xc2, 0x7a, 0xaf, 0x68,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SwitchServiceClient is the client API for SwitchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SwitchServiceClient interface {
	// turn switch on
	On(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	// turn stitch off
	Off(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	// get switch state
	State(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StateResponse, error)
}

type switchServiceClient struct {
	cc *grpc.ClientConn
}

func NewSwitchServiceClient(cc *grpc.ClientConn) SwitchServiceClient {
	return &switchServiceClient{cc}
}

func (c *switchServiceClient) On(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ai.metathings.protocol.service.switch.SwitchService/On", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *switchServiceClient) Off(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ai.metathings.protocol.service.switch.SwitchService/Off", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *switchServiceClient) State(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StateResponse, error) {
	out := new(StateResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.protocol.service.switch.SwitchService/State", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SwitchServiceServer is the server API for SwitchService service.
type SwitchServiceServer interface {
	// turn switch on
	On(context.Context, *empty.Empty) (*empty.Empty, error)
	// turn stitch off
	Off(context.Context, *empty.Empty) (*empty.Empty, error)
	// get switch state
	State(context.Context, *empty.Empty) (*StateResponse, error)
}

// UnimplementedSwitchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSwitchServiceServer struct {
}

func (*UnimplementedSwitchServiceServer) On(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method On not implemented")
}
func (*UnimplementedSwitchServiceServer) Off(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Off not implemented")
}
func (*UnimplementedSwitchServiceServer) State(ctx context.Context, req *empty.Empty) (*StateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method State not implemented")
}

func RegisterSwitchServiceServer(s *grpc.Server, srv SwitchServiceServer) {
	s.RegisterService(&_SwitchService_serviceDesc, srv)
}

func _SwitchService_On_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwitchServiceServer).On(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.protocol.service.switch.SwitchService/On",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwitchServiceServer).On(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwitchService_Off_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwitchServiceServer).Off(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.protocol.service.switch.SwitchService/Off",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwitchServiceServer).Off(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwitchService_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwitchServiceServer).State(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.protocol.service.switch.SwitchService/State",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwitchServiceServer).State(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SwitchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.protocol.service.switch.SwitchService",
	HandlerType: (*SwitchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "On",
			Handler:    _SwitchService_On_Handler,
		},
		{
			MethodName: "Off",
			Handler:    _SwitchService_Off_Handler,
		},
		{
			MethodName: "State",
			Handler:    _SwitchService_State_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
