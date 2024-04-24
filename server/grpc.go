package server

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	todos "github.com/sajuno/pace-generic-decorators/api/proto"
	"github.com/sajuno/pace-generic-decorators/auth"
	"github.com/sajuno/pace-generic-decorators/todo"
	"log/slog"
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

	if !u.Authorized() {
		slog.Error("unauthorized", "user", u)
		sentry.CaptureException(fmt.Errorf("unauthorized user %s tried to gain access", u.UUID()))
		return nil, fmt.Errorf("unauthorized")
	}

	td, err := s.repo.GetTodo(ctx, req.GetUuid(), u.UUID())
	if err != nil {
		slog.Error("failed to get todo for user", "user", u)
		sentry.CaptureException(fmt.Errorf("failed to get todo"))
		return nil, err
	}

	slog.Debug("successfully fetched todo", "todo", td)

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

	if !u.Authorized() {
		slog.Error("unauthorized", "user", u)
		sentry.CaptureException(fmt.Errorf("unauthorized user %s tried to gain access", u.UUID()))
		return nil, fmt.Errorf("unauthorized")
	}

	td, err := todo.NewTodo(uuid.NewString(), u.UUID(), req.Todo.GetText())
	if err != nil {
		slog.Error("todo validation failed", "todo", td)
		return nil, err
	}

	err = s.repo.CreateTodo(ctx, td)
	if err != nil {
		slog.Error("failed to create todo for user", "user", u, "todo", td)
		sentry.CaptureException(fmt.Errorf("failed to create todo"))
		return nil, err
	}

	slog.Debug("todo created", "todo", td)

	return &todos.TodoCreateResponse{Uuid: td.UUID()}, nil
}

//func GRPCAuthInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
//	// normally done client side and potentially intercepted and checked server side
//	return handler(context.WithValue(ctx, "user", auth.User{UUID: "dummy_user"}), req)
//}
