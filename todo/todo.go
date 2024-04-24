package todo

import "fmt"

type Todo struct {
	uuid   string
	userID string
	text   string
	done   bool
}

func NewTodo(uuid, userID, text string) (*Todo, error) {
	if uuid == "" {
		return nil, fmt.Errorf("todo's must have a valid uuid")
	}
	if userID == "" {
		return nil, fmt.Errorf("todo's must have a valid auth id")
	}
	if text == "" {
		return nil, fmt.Errorf("todo's must be provided with a text")
	}

	return &Todo{
		uuid:   uuid,
		userID: userID,
		text:   text,
	}, nil
}

// FromStore is meant to be used to create the domain object from the persisted object
// This is useful to avoid exposing the domain object properties for potentially adverse mutations
func FromStore(uuid, userID, text string, done bool) *Todo {
	return &Todo{
		uuid:   uuid,
		userID: userID,
		text:   text,
		done:   done,
	}
}

func (t Todo) UUID() string {
	return t.uuid
}

func (t Todo) UserID() string {
	return t.userID
}

func (t Todo) Text() string {
	return t.text
}

func (t Todo) Done() bool {
	return t.done
}

func (t *Todo) Finish() {
	t.done = true
}
