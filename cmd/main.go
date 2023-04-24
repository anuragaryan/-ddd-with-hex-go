package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ostafen/clover/v2"

	"github.com/anuragaryan/ddd-with-hex-go/internal/adapters/framework/database/nosql"
	//"github.com/anuragaryan/ddd-with-hex-go/internal/adapters/framework/database/memory"
	intHttp "github.com/anuragaryan/ddd-with-hex-go/internal/adapters/framework/presentation/http"
	"github.com/anuragaryan/ddd-with-hex-go/internal/application/events"
	todoevents "github.com/anuragaryan/ddd-with-hex-go/internal/application/events/todo"
	"github.com/anuragaryan/ddd-with-hex-go/internal/application/services/todo"
)

// TALK: dependency inversion.

func main() {
	eventsPublisher := events.NewEventPublisher()

	todoEventHandler := todoevents.NewEventHandler()
	eventsPublisher.Subscribe(todoEventHandler, todoevents.ListCreated{}, todoevents.ListItemCreated{}, todoevents.ListItemDone{})

	//inmemory := memory.New()
	db, _ := clover.Open("clover-db")
	db.CreateCollection(nosql.TodosCollection)
	defer db.Close()
	// TALK: emphasise on dependency injection.
	todoService, err := todo.NewService(
		//todo.WithMemoryRepository(inmemory),
		todo.WithNoSQLRepository(db),
		todo.WithEventsHandlers(eventsPublisher),
	)
	if err != nil {
		panic(err)
	}

	h, err := intHttp.NewHandler(
		intHttp.WithService(todoService),
	)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Route("/todo-list", func(r chi.Router) {
		r.Post("/", h.CreateList)
		r.Get("/", h.GetLists)
		r.Get("/{listID}", h.GetList)
		r.Post("/{listID}/items", h.CreateItem)
	})

	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
