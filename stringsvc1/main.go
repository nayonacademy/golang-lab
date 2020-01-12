// main.go
package main

import (
	"net/http"

	"github.com/sinnott74/TodoService/internal/todo"
)

func main() {

	service := todo.NewInmemTodoService()

	endpoints := todo.MakeTodoEndpoints(service)

	err := http.ListenAndServe(":8000", todo.MakeHTTPHandler(endpoints))
	if err != nil {
		panic(err)
	}
}
