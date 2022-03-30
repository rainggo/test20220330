// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package c

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DClient is the client API for D service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type dClient struct {
	cc grpc.ClientConnInterface
}

func NewDClient(cc grpc.ClientConnInterface) DClient {
	return &dClient{cc}
}

func (c *dClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/b.c.D/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DServer is the server API for D service.
// All implementations must embed UnimplementedDServer
// for forward compatibility
type DServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedDServer()
}

// UnimplementedDServer must be embedded to have forward compatible implementations.
type UnimplementedDServer struct {
}

func (UnimplementedDServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedDServer) mustEmbedUnimplementedDServer() {}

// UnsafeDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DServer will
// result in compilation errors.
type UnsafeDServer interface {
	mustEmbedUnimplementedDServer()
}

func RegisterDServer(s grpc.ServiceRegistrar, srv DServer) {
	s.RegisterService(&D_ServiceDesc, srv)
}

func _D_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/b.c.D/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// D_ServiceDesc is the grpc.ServiceDesc for D service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var D_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "b.c.D",
	HandlerType: (*DServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _D_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "D.proto",
}
