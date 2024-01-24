package controller

import (
	"fmt"
	"github.com/monstercameron/go-in-one-month/objects"
	"github.com/monstercameron/go-in-one-month/views/components"
	"github.com/monstercameron/go-in-one-month/views/pages"
	"net/http"
)

func GetTodo(w http.ResponseWriter, r *http.Request) {
	component := pages.IndexPage("Todos", objects.TodoList)
	// serve text/html
	w.Header().Set("Content-Type", "text/html")
	// render the component to the response writer
	component.Render(r.Context(), w)
}

func CheckTodo(w http.ResponseWriter, r *http.Request) {
	// Log request details (consider using a logging library for better control)
	fmt.Printf("Received request - Method: %s, URL: %s, Protocol: %s\n", r.Method, r.URL, r.Proto)

	// Check if method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the todo ID from the URL
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Todo ID is required", http.StatusBadRequest)
		return
	}

	// Retrieve the todo item
	todo := objects.TodoList.Get(id)
	if todo == nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Update the todo item
	todo.Checked = true

	// print todos slice
	// fmt.Println(todos.Todos)

	// Create and render the TodoComponent
	component := components.TodoComponent(todo)
	w.Header().Set("Content-Type", "text/html")
	component.Render(r.Context(), w)
}
