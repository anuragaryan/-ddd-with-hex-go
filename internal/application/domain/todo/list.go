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
	name      string
	id        string
	todos     []Item
	createdAt time.Time
	updatedAt time.Time
}

// NewList will create a new list of todos.
func NewList(n string) (*List, error) {
	if n == "" {
		return nil, errors.New("list needs a name")
	}

	return &List{
		id:        uuid.New().String(),
		name:      n,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

// GetID ...
func (l *List) GetID() string {
	return l.id
}

// GetName ...
func (l *List) GetName() string {
	return l.name
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

// ListItems lists all the items in the todo list.
func (l *List) ListItems() ([]Item, error) {
	return l.todos, nil
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
func (l *List) MarkItemDone(id string) {
	for _, i := range l.todos {
		if i.id == id {
			i.status = done
		}
	}
}

// Mutation on aggregates and bnot entities - call out
