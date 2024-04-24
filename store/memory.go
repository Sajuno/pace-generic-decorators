package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/sajuno/pace-generic-decorators/todo"
	"log/slog"
	"sync"
	"time"
)

type Todo struct {
	// some props that might be useful to just the store
	id        int
	createdAt time.Time

	uuid   string
	text   string
	userID string
	done   bool
}

type MemoryStore struct {
	data map[string]Todo

	sync.Mutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]Todo),
	}
}

func (m *MemoryStore) GetTodo(_ context.Context, uuid string, userID string) (*todo.Todo, error) {
	slog.Debug("getting todo", "user_id", userID)

	m.Lock()
	td, ok := m.data[uuid]
	if !ok || td.userID != userID {
		slog.Error("failed to get todo")
		sentry.CaptureException(errors.New("some sentry error"))
		return nil, fmt.Errorf("not found")
	}
	m.Unlock()

	return todo.FromStore(td.uuid, td.userID, td.text, td.done), nil
}

func (m *MemoryStore) CreateTodo(_ context.Context, todo *todo.Todo) error {
	slog.Debug("creating new todo", "todo", todo)

	m.Lock()
	if _, ok := m.data[todo.UUID()]; ok {
		slog.Error("")
		return fmt.Errorf("already exists")
	}

	m.data[todo.UUID()] = Todo{
		uuid:   todo.UUID(),
		text:   todo.Text(),
		userID: todo.UserID(),
		done:   todo.Done(),
	}
	m.Unlock()

	return nil
}
