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
	ID        string
	Name      string
	Todos     []Item
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewList will create a new list of todos.
func NewList(n string) (*List, error) {
	if n == "" {
		return nil, errors.New("list needs a name")
	}

	return &List{
		ID:        uuid.New().String(),
		Name:      n,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// AddItem adds item to the todo list.
func (l *List) AddItem(item Item) error {
	for _, i := range l.Todos {
		if i.ID == item.ID {
			return ErrItemAlreadyExists
		}
	}
	l.Todos = append(l.Todos, item)
	return nil
}

// ListItems lists all the items in the todo list.
func (l *List) ListItems() ([]Item, error) {
	return l.Todos, nil
}

// RemoveItem removes item from the todo list.
func (l *List) RemoveItem(item Item) {
	for idx, i := range l.Todos {
		if i.ID == item.ID {
			l.Todos = append(l.Todos[:idx], l.Todos[idx+1:]...)
			break
		}
	}
}

// MarkItemDone marks item in the todo list as done.
func (l *List) MarkItemDone(id string) {
	// TALK: Mutation on aggregates and not entities, entities are not exposed to the outisde world directly.
	for _, i := range l.Todos {
		if i.ID == id {
			i.Status = done
		}
	}
}
