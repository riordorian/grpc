// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: promo.proto

package grpc

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

const (
	Promos_PromoList_FullMethodName = "/grpc.Promos/PromoList"
	Promos_PromoInfo_FullMethodName = "/grpc.Promos/PromoInfo"
)

// PromosClient is the client API for Promos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PromosClient interface {
	PromoList(ctx context.Context, in *PromoListRequest, opts ...grpc.CallOption) (*PromosList, error)
	PromoInfo(ctx context.Context, in *PromoInfoRequest, opts ...grpc.CallOption) (*Promo, error)
}

type promosClient struct {
	cc grpc.ClientConnInterface
}

func NewPromosClient(cc grpc.ClientConnInterface) PromosClient {
	return &promosClient{cc}
}

func (c *promosClient) PromoList(ctx context.Context, in *PromoListRequest, opts ...grpc.CallOption) (*PromosList, error) {
	out := new(PromosList)
	err := c.cc.Invoke(ctx, Promos_PromoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promosClient) PromoInfo(ctx context.Context, in *PromoInfoRequest, opts ...grpc.CallOption) (*Promo, error) {
	out := new(Promo)
	err := c.cc.Invoke(ctx, Promos_PromoInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromosServer is the server API for Promos service.
// All implementations must embed UnimplementedPromosServer
// for forward compatibility
type PromosServer interface {
	PromoList(context.Context, *PromoListRequest) (*PromosList, error)
	PromoInfo(context.Context, *PromoInfoRequest) (*Promo, error)
	mustEmbedUnimplementedPromosServer()
}

// UnimplementedPromosServer must be embedded to have forward compatible implementations.
type UnimplementedPromosServer struct {
}

func (UnimplementedPromosServer) PromoList(context.Context, *PromoListRequest) (*PromosList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromoList not implemented")
}
func (UnimplementedPromosServer) PromoInfo(context.Context, *PromoInfoRequest) (*Promo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromoInfo not implemented")
}
func (UnimplementedPromosServer) mustEmbedUnimplementedPromosServer() {}

// UnsafePromosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PromosServer will
// result in compilation errors.
type UnsafePromosServer interface {
	mustEmbedUnimplementedPromosServer()
}

func RegisterPromosServer(s grpc.ServiceRegistrar, srv PromosServer) {
	s.RegisterService(&Promos_ServiceDesc, srv)
}

func _Promos_PromoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromosServer).PromoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Promos_PromoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromosServer).PromoList(ctx, req.(*PromoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Promos_PromoInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromosServer).PromoInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Promos_PromoInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromosServer).PromoInfo(ctx, req.(*PromoInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Promos_ServiceDesc is the grpc.ServiceDesc for Promos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Promos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Promos",
	HandlerType: (*PromosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PromoList",
			Handler:    _Promos_PromoList_Handler,
		},
		{
			MethodName: "PromoInfo",
			Handler:    _Promos_PromoInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "promo.proto",
}
