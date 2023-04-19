package todo

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	// ErrMissingValues is returned when an item is created without a text.
	ErrMissingValues = errors.New("missing values")
)

// Status is a value object.
type Status string

const (
	initialised Status = "initialised"
	done        Status = "done"
)

// Item is an entity.
type Item struct {
	id        string
	status    Status
	text      string
	createdAt time.Time
	updatedAt time.Time
}

// NewItem will create a new item.
// It will return error if it's an invalid item.
func NewItem(text string) (Item, error) {
	if text == "" {
		return Item{}, ErrMissingValues
	}

	return Item{
		id:        uuid.New().String(),
		status:    initialised,
		text:      text,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func (i Item) GetID() string {
	return i.id
}

func (i Item) GetText() string {
	return i.text
}

// Done marks the status of an item as done.
func (i Item) Done() {
	i.status = done
}
