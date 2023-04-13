// Package todo holds the repository and the implementations for a TodoRepository.
package todo

import (
	"errors"

	"github.com/google/uuid"
)

var (
	//ErrListNotFound is returned when a list is not found.
	ErrListNotFound = errors.New("the list was not found")
	//ErrListAlreadyExist is returned when trying to add a list that already exists.
	ErrListAlreadyExist = errors.New("the list already exists")
)

// Repository is the repository interface to fulfill to use the todo list aggregate.
type Repository interface {
	GetAll() ([]List, error)
	GetByID(uuid.UUID) (List, error)
	Add(list List) error
	Delete(uuid.UUID)

	AddItem(id uuid.UUID, item Item) error
	ListItem(id uuid.UUID) []Item
	MarkItemDone(id uuid.UUID, itemID uuid.UUID)
}
