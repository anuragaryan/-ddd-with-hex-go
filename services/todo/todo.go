package todo

import (
	"github.com/anuragaryan/ddd-with-hex-go/domain/todo"
	"github.com/anuragaryan/ddd-with-hex-go/domain/todo/database"
	"github.com/anuragaryan/ddd-with-hex-go/domain/todo/memory"
)

type Configuration func(os *Service) error

type Service struct {
	todos todo.Repository
}

func NewService(cfgs ...Configuration) (*Service, error) {
	s := &Service{}
	for _, cfg := range cfgs {
		err := cfg(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func withRepository(tr todo.Repository) Configuration {
	return func(s *Service) error {
		s.todos = tr
		return nil
	}
}

func WithMemoryRepository() Configuration {
	m := memory.New()
	return withRepository(m)
}

func WithDatabaseRepository() Configuration {
	d := database.New()
	return withRepository(d)
}

func (s *Service) CreateList() error {
	l, err := todo.NewList()
	if err != nil {
		return err
	}

	err = s.todos.Add(l)
	return err
}
