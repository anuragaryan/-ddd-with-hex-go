package memory

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

func (r *TodoListRepository) GetAll() ([]todo.List, error) {
	var lists []todo.List
	for _, list := range r.lists {
		lists = append(lists, list)
	}
	return lists, nil
}

func (r *TodoListRepository) GetByID(id uuid.UUID) (todo.List, error) {
	if list, ok := r.lists[id]; ok {
		return list, nil
	}
	return todo.List{}, todo.ErrListNotFound
}

func (r *TodoListRepository) Add(list todo.List) error {
	if _, ok := r.lists[list.GetID()]; ok {
		return todo.ErrListAlreadyExist
	}

	r.Lock()
	r.lists[list.GetID()] = list
	r.Unlock()

	return nil
}

func (r *TodoListRepository) Delete(id uuid.UUID) {
	r.Lock()
	delete(r.lists, id)
	r.Unlock()
}

func (r *TodoListRepository) AddItem(id uuid.UUID, item todo.Item) error {
	l, ok := r.lists[id]
	if !ok {
		return todo.ErrListNotFound
	}

	err := l.AddItem(item)
	if err != nil {
		return err
	}

	r.Lock()
	r.lists[id] = l
	r.Unlock()

	return nil
}

func (r *TodoListRepository) ListItem(id uuid.UUID) []todo.Item {
	return nil
}

func (r *TodoListRepository) MarkItemDone(id uuid.UUID, itemID uuid.UUID) {}
