package server

import (
	"context"
	"github.com/google/uuid"
	todos "github.com/sajuno/pace-generic-decorators/api/proto"
	"github.com/sajuno/pace-generic-decorators/auth"
	"github.com/sajuno/pace-generic-decorators/todo"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	todos.TodosServer

	repo todo.Repository
}

func NewGRPCServer(repo todo.Repository) *GRPCServer {
	return &GRPCServer{repo: repo}
}

func (s GRPCServer) GetTodo(ctx context.Context, req *todos.GetTodoRequest) (*todos.TodoResponse, error) {
	u, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	td, err := s.repo.GetTodo(ctx, req.GetUuid(), u.UUID)
	if err != nil {
		return nil, err
	}

	return &todos.TodoResponse{
		Todo: &todos.Todo{
			Uuid:   td.UUID(),
			Text:   td.Text(),
			Done:   td.Done(),
			UserId: td.UserID(),
		},
	}, nil
}

func (s GRPCServer) CreateTodo(ctx context.Context, req *todos.CreateTodoRequest) (*todos.TodoCreateResponse, error) {
	u, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	td, err := todo.NewTodo(uuid.NewString(), u.UUID, req.Todo.GetText())
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateTodo(ctx, td)
	if err != nil {
		return nil, err
	}

	return &todos.TodoCreateResponse{Uuid: td.UUID()}, nil
}

func GRPCAuthInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// normally done client side and potentially intercepted and checked server side
	return handler(context.WithValue(ctx, "user", auth.User{UUID: "dummy_user"}), req)
}
