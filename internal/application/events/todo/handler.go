package todo

import (
	"errors"
	"fmt"

	"github.com/anuragaryan/ddd-with-hex-go/internal/application/events"
)

type EventHandler struct{}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

func (teh *EventHandler) Notify(event events.Event) error {
	switch e := event.(type) {
	case ListCreated:
		fmt.Printf("Handled event: %s with List ID %s", e.Name(), e.ListID())
	case ListItemCreated:
		fmt.Printf("Handled event: %s with List ID %s and Item ID %s", e.Name(), e.ListID(), e.ItemID())
	case ListItemDone:
		fmt.Printf("Handled event: %s with List ID %s and Item ID %s", e.Name(), e.ListID(), e.ItemID())
	default:
		return errors.New("no handler found for the event")
	}

	return nil
}
