package todo

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	// ErrItemAlreadyExists is returned when a todo item is added to a list it already exists in.
	ErrItemAlreadyExists = errors.New("item already exists")
)

// List is an aggregate.
type List struct {
	id        uuid.UUID
	todos     []Item
	createdAt time.Time
	updatedAt time.Time
}

// NewList will create a new list of todos.
func NewList() (List, error) {
	return List{
		id:        uuid.New(),
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func (l *List) GetID() uuid.UUID {
	return l.id
}

// AddItem adds item to the todo list.
func (l *List) AddItem(item Item) error {
	for _, i := range l.todos {
		if i.id == item.id {
			return ErrItemAlreadyExists
		}
	}
	l.todos = append(l.todos, item)
	return nil
}

// RemoveItem removes item from the todo list.
func (l *List) RemoveItem(item Item) {
	for idx, i := range l.todos {
		if i.id == item.id {
			l.todos = append(l.todos[:idx], l.todos[idx+1:]...)
			break
		}
	}
}

// MarkItemDone marks item in the todo list as done.
func (l *List) MarkItemDone(id uuid.UUID) {
	for _, i := range l.todos {
		if i.id == id {
			i.status = done
		}
	}
}
