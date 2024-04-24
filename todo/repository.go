package todo

import "context"

type Repository interface {
	GetTodo(ctx context.Context, uuid string, userID string) (*Todo, error)
	CreateTodo(ctx context.Context, todo *Todo) error
}
