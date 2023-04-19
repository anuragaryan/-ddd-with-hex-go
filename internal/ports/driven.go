package ports

import (
	"github.com/anuragaryan/ddd-with-hex-go/internal/application/domain/todo"
	"github.com/anuragaryan/ddd-with-hex-go/internal/application/events"
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

//go:generate mockgen -source=driven.go -destination=mocks/driven.go EventHandlerPort
// EventHandlerPort interface defining the contracts for event handler.
type EventHandlerPort interface {
	Subscribe(handler events.EventHandler, events ...events.Event)
	Notify(event events.Event) error
}
