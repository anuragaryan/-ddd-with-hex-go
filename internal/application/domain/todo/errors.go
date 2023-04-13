// Package todo holds the repository and the implementations for a TodoRepository.
package todo

import (
	"errors"
)

var (
	//ErrListNotFound is returned when a list is not found.
	ErrListNotFound = errors.New("the list was not found")
	//ErrListAlreadyExist is returned when trying to add a list that already exists.
	ErrListAlreadyExist = errors.New("the list already exists")
)
