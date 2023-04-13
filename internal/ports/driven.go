package ports

import (
	"github.com/anuragaryan/ddd-with-hex-go/internal/application/domain/todo"
)

//go:generate mockgen -source=driven.go -destination=mocks/driven.go StoragePort
// StoragePort is technology-agnostic and can be implemented using any kind of data storage.
type StoragePort interface {
	GetAll() ([]todo.List, error)
	GetByID(id string) (todo.List, error)
	Add(list todo.List) error
	Delete(id string)

	AddItem(id string, item todo.Item) error
	ListItem(id string) ([]todo.Item, error)
	MarkItemDone(id string, itemID string)
}
