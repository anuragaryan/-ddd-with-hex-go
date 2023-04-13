package database

import (
	"sync"

	"github.com/google/uuid"

	"github.com/anuragaryan/ddd-with-hex-go/domain/todo"
)

type TodoListRepository struct {
	lists map[uuid.UUID]todo.List
	sync.Mutex
}

// New is a factory function to generate a new repository of todo lists.
func New() *TodoListRepository {
	return &TodoListRepository{
		lists: make(map[uuid.UUID]todo.List),
	}
}

func (t TodoListRepository) GetAll() ([]todo.List, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepository) GetByID(u uuid.UUID) (todo.List, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepository) Add(list todo.List) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepository) Delete(u uuid.UUID) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepository) AddItem(id uuid.UUID, item todo.Item) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepository) ListItem(id uuid.UUID) []todo.Item {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepository) MarkItemDone(id uuid.UUID, itemID uuid.UUID) {
	//TODO implement me
	panic("implement me")
}
