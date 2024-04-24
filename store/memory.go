package store

import (
	"context"
	"fmt"
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
	m.Lock()
	td, ok := m.data[uuid]
	if !ok || td.userID != userID {
		return nil, fmt.Errorf("not found")
	}
	m.Unlock()

	return todo.FromStore(td.uuid, td.userID, td.text, td.done), nil
}

func (m *MemoryStore) CreateTodo(_ context.Context, todo *todo.Todo) error {
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
