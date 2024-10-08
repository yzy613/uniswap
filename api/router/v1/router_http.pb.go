// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.1
// source: router/v1/router.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationRouterexactInputSingle = "/api.router.v1.Router/exactInputSingle"
const OperationRouterexactOutputSingle = "/api.router.v1.Router/exactOutputSingle"

type RouterHTTPServer interface {
	ExactInputSingle(context.Context, *ExactInputSingleRequest) (*ExactInputSingleReply, error)
	ExactOutputSingle(context.Context, *ExactOutputSingleRequest) (*ExactOutputSingleReply, error)
}

func RegisterRouterHTTPServer(s *http.Server, srv RouterHTTPServer) {
	r := s.Route("/")
	r.POST("/router/exactInputSingle", _Router_ExactInputSingle0_HTTP_Handler(srv))
	r.POST("/router/exactOutputSingle", _Router_ExactOutputSingle0_HTTP_Handler(srv))
}

func _Router_ExactInputSingle0_HTTP_Handler(srv RouterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExactInputSingleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRouterexactInputSingle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExactInputSingle(ctx, req.(*ExactInputSingleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExactInputSingleReply)
		return ctx.Result(200, reply)
	}
}

func _Router_ExactOutputSingle0_HTTP_Handler(srv RouterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExactOutputSingleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRouterexactOutputSingle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExactOutputSingle(ctx, req.(*ExactOutputSingleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExactOutputSingleReply)
		return ctx.Result(200, reply)
	}
}

type RouterHTTPClient interface {
	ExactInputSingle(ctx context.Context, req *ExactInputSingleRequest, opts ...http.CallOption) (rsp *ExactInputSingleReply, err error)
	ExactOutputSingle(ctx context.Context, req *ExactOutputSingleRequest, opts ...http.CallOption) (rsp *ExactOutputSingleReply, err error)
}

type RouterHTTPClientImpl struct {
	cc *http.Client
}

func NewRouterHTTPClient(client *http.Client) RouterHTTPClient {
	return &RouterHTTPClientImpl{client}
}

func (c *RouterHTTPClientImpl) ExactInputSingle(ctx context.Context, in *ExactInputSingleRequest, opts ...http.CallOption) (*ExactInputSingleReply, error) {
	var out ExactInputSingleReply
	pattern := "/router/exactInputSingle"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRouterexactInputSingle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *RouterHTTPClientImpl) ExactOutputSingle(ctx context.Context, in *ExactOutputSingleRequest, opts ...http.CallOption) (*ExactOutputSingleReply, error) {
	var out ExactOutputSingleReply
	pattern := "/router/exactOutputSingle"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationRouterexactOutputSingle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
