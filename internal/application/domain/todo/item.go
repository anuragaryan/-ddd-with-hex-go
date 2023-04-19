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
	ID        string
	Status    Status
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewItem will create a new item.
// It will return error if it's an invalid item.
func NewItem(text string) (Item, error) {
	if text == "" {
		return Item{}, ErrMissingValues
	}

	return Item{
		ID:        uuid.New().String(),
		Status:    initialised,
		Text:      text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// Done marks the status of an item as done.
func (i Item) Done() {
	i.Status = done
}
