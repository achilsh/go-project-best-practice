// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: otal-tracer/demo_protocal.proto

package demo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GrpcCallDemo_Echo_FullMethodName  = "/grpc.demo.GrpcCallDemo/Echo"
	GrpcCallDemo_Hello_FullMethodName = "/grpc.demo.GrpcCallDemo/Hello"
)

// GrpcCallDemoClient is the client API for GrpcCallDemo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 使用注解 来设置 http 的url 和body
type GrpcCallDemoClient interface {
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
	Hello(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type grpcCallDemoClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcCallDemoClient(cc grpc.ClientConnInterface) GrpcCallDemoClient {
	return &grpcCallDemoClient{cc}
}

func (c *grpcCallDemoClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, GrpcCallDemo_Echo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcCallDemoClient) Hello(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, GrpcCallDemo_Hello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcCallDemoServer is the server API for GrpcCallDemo service.
// All implementations must embed UnimplementedGrpcCallDemoServer
// for forward compatibility.
//
// 使用注解 来设置 http 的url 和body
type GrpcCallDemoServer interface {
	Echo(context.Context, *StringMessage) (*StringMessage, error)
	Hello(context.Context, *StringMessage) (*StringMessage, error)
	mustEmbedUnimplementedGrpcCallDemoServer()
}

// UnimplementedGrpcCallDemoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGrpcCallDemoServer struct{}

func (UnimplementedGrpcCallDemoServer) Echo(context.Context, *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedGrpcCallDemoServer) Hello(context.Context, *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedGrpcCallDemoServer) mustEmbedUnimplementedGrpcCallDemoServer() {}
func (UnimplementedGrpcCallDemoServer) testEmbeddedByValue()                      {}

// UnsafeGrpcCallDemoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpcCallDemoServer will
// result in compilation errors.
type UnsafeGrpcCallDemoServer interface {
	mustEmbedUnimplementedGrpcCallDemoServer()
}

func RegisterGrpcCallDemoServer(s grpc.ServiceRegistrar, srv GrpcCallDemoServer) {
	// If the following call pancis, it indicates UnimplementedGrpcCallDemoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GrpcCallDemo_ServiceDesc, srv)
}

func _GrpcCallDemo_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcCallDemoServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GrpcCallDemo_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcCallDemoServer).Echo(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcCallDemo_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcCallDemoServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GrpcCallDemo_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcCallDemoServer).Hello(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// GrpcCallDemo_ServiceDesc is the grpc.ServiceDesc for GrpcCallDemo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GrpcCallDemo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.demo.GrpcCallDemo",
	HandlerType: (*GrpcCallDemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _GrpcCallDemo_Echo_Handler,
		},
		{
			MethodName: "Hello",
			Handler:    _GrpcCallDemo_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "otal-tracer/demo_protocal.proto",
}
