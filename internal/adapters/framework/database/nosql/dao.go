package nosql

import (
	"time"

	"github.com/anuragaryan/ddd-with-hex-go/internal/application/domain/todo"
)

type list struct {
	ID        string    `clover:"id"`
	Name      string    `clover:"name"`
	Todos     []item    `clover:"todos"`
	CreatedAt time.Time `clover:"created_at"`
	UpdatedAt time.Time `clover:"updated_at"`
}

type item struct {
	ID        string    `clover:"id"`
	Status    string    `clover:"status"`
	Text      string    `clover:"text"`
	CreatedAt time.Time `clover:"created_at"`
	UpdatedAt time.Time `clover:"updated_at"`
}

func toDomainList(l list) todo.List {
	domainItems := make([]todo.Item, len(l.Todos))
	for idx, t := range l.Todos {
		domainItems[idx] = toDomainItem(t)
	}
	return todo.List{
		ID:        l.ID,
		Name:      l.Name,
		Todos:     domainItems,
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
	}
}

func toDomainItem(i item) todo.Item {
	return todo.Item{
		ID:        i.ID,
		Status:    todo.Status(i.Status),
		Text:      i.Text,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
}

func fromDomainList(l todo.List) list {
	items := make([]item, len(l.Todos))
	for idx, t := range l.Todos {
		items[idx] = fromDomainItem(t)
	}

	return list{
		ID:        l.ID,
		Name:      l.Name,
		Todos:     items,
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
	}
}

func fromDomainItem(i todo.Item) item {
	return item{
		ID:        i.ID,
		Status:    string(i.Status),
		Text:      i.Text,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
}
