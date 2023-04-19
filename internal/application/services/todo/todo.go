package todo

import (
	"github.com/anuragaryan/ddd-with-hex-go/internal/adapters/framework/database/memory"
	"github.com/anuragaryan/ddd-with-hex-go/internal/adapters/framework/database/sql"
	"github.com/anuragaryan/ddd-with-hex-go/internal/application/domain/todo"
	"github.com/anuragaryan/ddd-with-hex-go/internal/ports"
)

type Configuration func(os *Service) error

type Service struct {
	todos ports.StoragePort
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

func withRepository(tr ports.StoragePort) Configuration {
	return func(s *Service) error {
		s.todos = tr
		return nil
	}
}

func WithMemoryRepository() Configuration {
	m := memory.New()
	return withRepository(m)
}

func WithSQLRepository() Configuration {
	d := sql.New()
	return withRepository(d)
}

func (s *Service) CreateList(name string) error {
	l, err := todo.NewList(name)
	if err != nil {
		return err
	}

	err = s.todos.Add(*l)
	// TODO: Dispatch an event and later on wire up a notifier which notifies emails when a list is created/modified.

	return err
}

func (s *Service) GetList(id string) (*todo.List, error) {
	l, err := s.todos.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &l, err
}

func (s *Service) GetLists() ([]todo.List, error) {
	l, err := s.todos.GetAll()
	if err != nil {
		return nil, err
	}

	return l, err
}

func (s *Service) AddItemToList(id string, item string) error {
	l, err := s.todos.GetByID(id)
	if err != nil {
		return err
	}

	ii, err := todo.NewItem(item)
	if err != nil {
		return err
	}

	err = l.AddItem(ii)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAllItemsFromAList(id string) ([]todo.Item, error) {
	return s.todos.ListItem(id)
}
