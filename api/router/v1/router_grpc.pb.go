// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: router/v1/router.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Router_ExactInputSingle_FullMethodName  = "/api.router.v1.Router/exactInputSingle"
	Router_ExactOutputSingle_FullMethodName = "/api.router.v1.Router/exactOutputSingle"
)

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouterClient interface {
	ExactInputSingle(ctx context.Context, in *ExactInputSingleRequest, opts ...grpc.CallOption) (*ExactInputSingleReply, error)
	ExactOutputSingle(ctx context.Context, in *ExactOutputSingleRequest, opts ...grpc.CallOption) (*ExactOutputSingleReply, error)
}

type routerClient struct {
	cc grpc.ClientConnInterface
}

func NewRouterClient(cc grpc.ClientConnInterface) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) ExactInputSingle(ctx context.Context, in *ExactInputSingleRequest, opts ...grpc.CallOption) (*ExactInputSingleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExactInputSingleReply)
	err := c.cc.Invoke(ctx, Router_ExactInputSingle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) ExactOutputSingle(ctx context.Context, in *ExactOutputSingleRequest, opts ...grpc.CallOption) (*ExactOutputSingleReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExactOutputSingleReply)
	err := c.cc.Invoke(ctx, Router_ExactOutputSingle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
// All implementations must embed UnimplementedRouterServer
// for forward compatibility
type RouterServer interface {
	ExactInputSingle(context.Context, *ExactInputSingleRequest) (*ExactInputSingleReply, error)
	ExactOutputSingle(context.Context, *ExactOutputSingleRequest) (*ExactOutputSingleReply, error)
	mustEmbedUnimplementedRouterServer()
}

// UnimplementedRouterServer must be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (UnimplementedRouterServer) ExactInputSingle(context.Context, *ExactInputSingleRequest) (*ExactInputSingleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExactInputSingle not implemented")
}
func (UnimplementedRouterServer) ExactOutputSingle(context.Context, *ExactOutputSingleRequest) (*ExactOutputSingleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExactOutputSingle not implemented")
}
func (UnimplementedRouterServer) mustEmbedUnimplementedRouterServer() {}

// UnsafeRouterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouterServer will
// result in compilation errors.
type UnsafeRouterServer interface {
	mustEmbedUnimplementedRouterServer()
}

func RegisterRouterServer(s grpc.ServiceRegistrar, srv RouterServer) {
	s.RegisterService(&Router_ServiceDesc, srv)
}

func _Router_ExactInputSingle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExactInputSingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).ExactInputSingle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_ExactInputSingle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).ExactInputSingle(ctx, req.(*ExactInputSingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_ExactOutputSingle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExactOutputSingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).ExactOutputSingle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_ExactOutputSingle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).ExactOutputSingle(ctx, req.(*ExactOutputSingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Router_ServiceDesc is the grpc.ServiceDesc for Router service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Router_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.router.v1.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "exactInputSingle",
			Handler:    _Router_ExactInputSingle_Handler,
		},
		{
			MethodName: "exactOutputSingle",
			Handler:    _Router_ExactOutputSingle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router/v1/router.proto",
}
