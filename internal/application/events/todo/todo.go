package todo

import "github.com/anuragaryan/ddd-with-hex-go/internal/application/events"

// ToDoEvent interface for describing todo relevant Domain Event.
type ToDoEvent interface {
	events.Event
	ListID() string
	ItemID() string
}

// ListCreated event.
type ListCreated struct {
	listID string
}

// NewListCreatedEvent creates a list created event.
func NewListCreatedEvent(id string) ListCreated {
	return ListCreated{
		listID: id,
	}
}

func (e ListCreated) Name() string {
	return "event.list.created"
}

func (e ListCreated) ListID() string {
	return e.listID
}

func (e ListCreated) ItemID() string {
	return ""
}

type ListItemCreated struct {
	listID string
	itemID string
}

func (e ListItemCreated) Name() string {
	return "event.list.item.created"
}

func (e ListItemCreated) ListID() string {
	return e.listID
}

func (e ListItemCreated) ItemID() string {
	return e.itemID
}

type ListItemDone struct {
	listID string
	itemID string
}

func (e ListItemDone) Name() string {
	return "event.list.item.done"
}

func (e ListItemDone) ListID() string {
	return e.listID
}

func (e ListItemDone) ItemID() string {
	return e.itemID
}
