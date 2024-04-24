// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/todos.proto

package todos

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

// TodosClient is the client API for Todos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodosClient interface {
	GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error)
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*TodoCreateResponse, error)
}

type todosClient struct {
	cc grpc.ClientConnInterface
}

func NewTodosClient(cc grpc.ClientConnInterface) TodosClient {
	return &todosClient{cc}
}

func (c *todosClient) GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*TodoResponse, error) {
	out := new(TodoResponse)
	err := c.cc.Invoke(ctx, "/todos.Todos/GetTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*TodoCreateResponse, error) {
	out := new(TodoCreateResponse)
	err := c.cc.Invoke(ctx, "/todos.Todos/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodosServer is the server API for Todos service.
// All implementations must embed UnimplementedTodosServer
// for forward compatibility
type TodosServer interface {
	GetTodo(context.Context, *GetTodoRequest) (*TodoResponse, error)
	CreateTodo(context.Context, *CreateTodoRequest) (*TodoCreateResponse, error)
	mustEmbedUnimplementedTodosServer()
}

// UnimplementedTodosServer must be embedded to have forward compatible implementations.
type UnimplementedTodosServer struct {
}

func (UnimplementedTodosServer) GetTodo(context.Context, *GetTodoRequest) (*TodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}
func (UnimplementedTodosServer) CreateTodo(context.Context, *CreateTodoRequest) (*TodoCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodosServer) mustEmbedUnimplementedTodosServer() {}

// UnsafeTodosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodosServer will
// result in compilation errors.
type UnsafeTodosServer interface {
	mustEmbedUnimplementedTodosServer()
}

func RegisterTodosServer(s grpc.ServiceRegistrar, srv TodosServer) {
	s.RegisterService(&Todos_ServiceDesc, srv)
}

func _Todos_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.Todos/GetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).GetTodo(ctx, req.(*GetTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todos.Todos/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Todos_ServiceDesc is the grpc.ServiceDesc for Todos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Todos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todos.Todos",
	HandlerType: (*TodosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTodo",
			Handler:    _Todos_GetTodo_Handler,
		},
		{
			MethodName: "CreateTodo",
			Handler:    _Todos_CreateTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/todos.proto",
}
