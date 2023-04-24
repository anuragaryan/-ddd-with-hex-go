package nosql

import (
	"github.com/anuragaryan/ddd-with-hex-go/internal/application/domain/todo"

	"github.com/ostafen/clover/v2"
)

const TodosCollection = "todos"

type TodoListRepository struct {
	db *clover.DB
}

// New is a factory function to generate a new repository of todo lists.
func New(db *clover.DB) *TodoListRepository {
	return &TodoListRepository{
		db: db,
	}
}

func (t TodoListRepository) GetAll() ([]todo.List, error) {
	docs, err := t.db.FindAll(clover.NewQuery(TodosCollection))
	if err != nil {
		return nil, err
	}

	var ll []todo.List

	for _, doc := range docs {
		l := &list{}
		err = doc.Unmarshal(l)
		if err != nil {
			return nil, err
		}

		ll = append(ll, toDomainList(*l))
	}

	return ll, nil
}

func (t TodoListRepository) GetByID(u string) (todo.List, error) {
	doc, err := t.db.FindFirst(clover.NewQuery(TodosCollection).Where(clover.Field("id").Eq(u)))
	if err != nil {
		return todo.List{}, todo.ErrListNotFound
	}

	l := &list{}
	err = doc.Unmarshal(l)
	if err != nil {
		return todo.List{}, todo.ErrListNotFound
	}

	return toDomainList(*l), nil
}

func (t TodoListRepository) Add(list todo.List) error {
	l := fromDomainList(list)

	doc := clover.NewDocumentOf(l)
	_, err := t.db.InsertOne(TodosCollection, doc)
	return err
}

func (t TodoListRepository) Delete(u string) {
	_ = t.db.Delete(clover.NewQuery(TodosCollection).Where(clover.Field("id").Eq(u)))
}

func (t TodoListRepository) AddItem(id string, item todo.Item) error {
	doc, err := t.db.FindFirst(clover.NewQuery(TodosCollection).Where(clover.Field("id").Eq(id)))
	if err != nil {
		return todo.ErrListNotFound
	}

	l := &list{}
	err = doc.Unmarshal(l)
	if err != nil {
		return todo.ErrListNotFound
	}

	l.Todos = append(l.Todos, fromDomainItem(item))

	err = t.db.ReplaceById(TodosCollection, doc.ObjectId(), clover.NewDocumentOf(l))
	return err
}

func (t TodoListRepository) ListItem(id string) ([]todo.Item, error) {
	doc, err := t.db.FindFirst(clover.NewQuery(TodosCollection).Where(clover.Field("id").Eq(id)))
	if err != nil {
		return nil, todo.ErrListNotFound
	}

	l := &list{}
	err = doc.Unmarshal(l)
	if err != nil {
		return nil, todo.ErrListNotFound
	}

	var ii []todo.Item

	for _, i := range l.Todos {
		ii = append(ii, toDomainItem(i))
	}

	return ii, nil
}

func (t TodoListRepository) MarkItemDone(id string, itemID string) {}
