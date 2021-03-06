// Code generated by protoc-gen-go. DO NOT EDIT.
// source: add-num.proto

/*
Package addition is a generated protocol buffer package.

It is generated from these files:
	add-num.proto

It has these top-level messages:
	Numbers
	Addition
*/
package addition

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Numbers struct {
	Number1 int64 `protobuf:"varint,1,opt,name=number1" json:"number1,omitempty"`
	Number2 int64 `protobuf:"varint,2,opt,name=number2" json:"number2,omitempty"`
}

func (m *Numbers) Reset()                    { *m = Numbers{} }
func (m *Numbers) String() string            { return proto.CompactTextString(m) }
func (*Numbers) ProtoMessage()               {}
func (*Numbers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Numbers) GetNumber1() int64 {
	if m != nil {
		return m.Number1
	}
	return 0
}

func (m *Numbers) GetNumber2() int64 {
	if m != nil {
		return m.Number2
	}
	return 0
}

type Addition struct {
	// addtion holds addition
	Addition int64 `protobuf:"varint,1,opt,name=addition" json:"addition,omitempty"`
}

func (m *Addition) Reset()                    { *m = Addition{} }
func (m *Addition) String() string            { return proto.CompactTextString(m) }
func (*Addition) ProtoMessage()               {}
func (*Addition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Addition) GetAddition() int64 {
	if m != nil {
		return m.Addition
	}
	return 0
}

func init() {
	proto.RegisterType((*Numbers)(nil), "addition.Numbers")
	proto.RegisterType((*Addition)(nil), "addition.Addition")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Adder service

type AdderClient interface {
	// Sends addition of two numbers
	Add(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Addition, error)
}

type adderClient struct {
	cc *grpc.ClientConn
}

func NewAdderClient(cc *grpc.ClientConn) AdderClient {
	return &adderClient{cc}
}

func (c *adderClient) Add(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Addition, error) {
	out := new(Addition)
	err := grpc.Invoke(ctx, "/addition.Adder/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Adder service

type AdderServer interface {
	// Sends addition of two numbers
	Add(context.Context, *Numbers) (*Addition, error)
}

func RegisterAdderServer(s *grpc.Server, srv AdderServer) {
	s.RegisterService(&_Adder_serviceDesc, srv)
}

func _Adder_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Numbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addition.Adder/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServer).Add(ctx, req.(*Numbers))
	}
	return interceptor(ctx, in, info, handler)
}

var _Adder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "addition.Adder",
	HandlerType: (*AdderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Adder_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "add-num.proto",
}

func init() { proto.RegisterFile("add-num.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4c, 0x49, 0xd1,
	0xcd, 0x2b, 0xcd, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x4c, 0x49, 0xc9, 0x2c,
	0xc9, 0xcc, 0xcf, 0x53, 0xb2, 0xe5, 0x62, 0xf7, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0x2a, 0x16, 0x92,
	0xe0, 0x62, 0xcf, 0x03, 0x33, 0x0d, 0x25, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x60, 0x5c, 0x84,
	0x8c, 0x91, 0x04, 0x13, 0xb2, 0x8c, 0x91, 0x92, 0x1a, 0x17, 0x87, 0x23, 0xd4, 0x28, 0x21, 0x29,
	0x2e, 0xb8, 0xb1, 0x50, 0x03, 0xe0, 0x7c, 0x23, 0x73, 0x2e, 0x56, 0xc7, 0x94, 0x94, 0xd4, 0x22,
	0x21, 0x3d, 0x2e, 0x66, 0xc7, 0x94, 0x14, 0x21, 0x41, 0x3d, 0x98, 0x94, 0x1e, 0xd4, 0x7a, 0x29,
	0x21, 0x84, 0x10, 0xcc, 0x48, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x83, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xbf, 0xdc, 0xf0, 0xff, 0xc1, 0x00, 0x00, 0x00,
}
