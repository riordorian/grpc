// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: new.proto

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
	News_List_FullMethodName   = "/grpc.News/List"
	News_Create_FullMethodName = "/grpc.News/Create"
)

// NewsClient is the client API for News service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*NewsList, error)
	Create(ctx context.Context, opts ...grpc.CallOption) (News_CreateClient, error)
}

type newsClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsClient(cc grpc.ClientConnInterface) NewsClient {
	return &newsClient{cc}
}

func (c *newsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*NewsList, error) {
	out := new(NewsList)
	err := c.cc.Invoke(ctx, News_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) Create(ctx context.Context, opts ...grpc.CallOption) (News_CreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &News_ServiceDesc.Streams[0], News_Create_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &newsCreateClient{stream}
	return x, nil
}

type News_CreateClient interface {
	Send(*CreateRequest) error
	CloseAndRecv() (*CreateResponse, error)
	grpc.ClientStream
}

type newsCreateClient struct {
	grpc.ClientStream
}

func (x *newsCreateClient) Send(m *CreateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *newsCreateClient) CloseAndRecv() (*CreateResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NewsServer is the server API for News service.
// All implementations must embed UnimplementedNewsServer
// for forward compatibility
type NewsServer interface {
	List(context.Context, *ListRequest) (*NewsList, error)
	Create(News_CreateServer) error
	mustEmbedUnimplementedNewsServer()
}

// UnimplementedNewsServer must be embedded to have forward compatible implementations.
type UnimplementedNewsServer struct {
}

func (UnimplementedNewsServer) List(context.Context, *ListRequest) (*NewsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedNewsServer) Create(News_CreateServer) error {
	return status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNewsServer) mustEmbedUnimplementedNewsServer() {}

// UnsafeNewsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsServer will
// result in compilation errors.
type UnsafeNewsServer interface {
	mustEmbedUnimplementedNewsServer()
}

func RegisterNewsServer(s grpc.ServiceRegistrar, srv NewsServer) {
	s.RegisterService(&News_ServiceDesc, srv)
}

func _News_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_Create_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NewsServer).Create(&newsCreateServer{stream})
}

type News_CreateServer interface {
	SendAndClose(*CreateResponse) error
	Recv() (*CreateRequest, error)
	grpc.ServerStream
}

type newsCreateServer struct {
	grpc.ServerStream
}

func (x *newsCreateServer) SendAndClose(m *CreateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *newsCreateServer) Recv() (*CreateRequest, error) {
	m := new(CreateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// News_ServiceDesc is the grpc.ServiceDesc for News service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var News_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.News",
	HandlerType: (*NewsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _News_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Create",
			Handler:       _News_Create_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "new.proto",
}
