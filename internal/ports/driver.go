package ports

import (
	"github.com/anuragaryan/ddd-with-hex-go/internal/application/domain/todo"
)

//go:generate mockgen -source=driver.go -destination=mocks/driver.go APIPort

// APIPort is technology-agnostic and can be used by any of http, grpc, etc.
type APIPort interface {
	CreateList(name string) error
	GetList(id string) (*todo.List, error)
	GetLists() ([]todo.List, error)
	AddItemToList(id string, item string) error
	GetAllItemsFromAList(id string) ([]todo.Item, error)
}
