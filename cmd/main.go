package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	intHttp "github.com/anuragaryan/ddd-with-hex-go/internal/adapters/framework/http"
	"github.com/anuragaryan/ddd-with-hex-go/internal/application/services/todo"
)

func main() {

	todoService, err := todo.NewService(
		todo.WithMemoryRepository(),
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
