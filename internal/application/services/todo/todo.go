package todo

import (
	"github.com/ostafen/clover/v2"

	"github.com/anuragaryan/ddd-with-hex-go/internal/adapters/framework/database/memory"
	"github.com/anuragaryan/ddd-with-hex-go/internal/adapters/framework/database/nosql"
	"github.com/anuragaryan/ddd-with-hex-go/internal/application/domain/todo"
	todo2 "github.com/anuragaryan/ddd-with-hex-go/internal/application/events/todo"
	"github.com/anuragaryan/ddd-with-hex-go/internal/ports"
)

type Configuration func(os *Service) error

type Service struct {
	todos  ports.StoragePort
	events ports.EventHandlerPort
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

func WithMemoryRepository(m *memory.TodoListRepository) Configuration {
	return withRepository(m)
}

func WithNoSQLRepository(db *clover.DB) Configuration {
	d := nosql.New(db)
	return withRepository(d)
}

func WithEventsHandlers(e ports.EventHandlerPort) Configuration {
	return func(s *Service) error {
		s.events = e
		return nil
	}
}

func (s *Service) CreateList(name string) error {
	l, err := todo.NewList(name)
	if err != nil {
		return err
	}

	err = s.todos.Add(*l)
	if err != nil {
		return err
	}

	err = s.events.Notify(todo2.NewListCreatedEvent(l.ID))

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
