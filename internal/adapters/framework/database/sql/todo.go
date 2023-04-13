package sql

import (
	"github.com/anuragaryan/ddd-with-hex-go/internal/application/domain/todo"
)

type TodoListRepository struct{}

// New is a factory function to generate a new repository of todo lists.
func New() *TodoListRepository {
	return &TodoListRepository{}
}

func (t TodoListRepository) GetAll() ([]todo.List, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepository) GetByID(u string) (todo.List, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepository) Add(list todo.List) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepository) Delete(u string) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepository) AddItem(id string, item todo.Item) error {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepository) ListItem(id string) ([]todo.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (t TodoListRepository) MarkItemDone(id string, itemID string) {
	//TODO implement me
	panic("implement me")
}
