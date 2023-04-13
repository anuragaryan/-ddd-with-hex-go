package main

import "github.com/anuragaryan/ddd-with-hex-go/services/todo"

func main() {

	todoService, err := todo.NewService(
		todo.WithMemoryRepository(),
	)
	if err != nil {
		panic(err)
	}

	err = todoService.CreateList()
	if err != nil {
		panic(err)
	}
}
