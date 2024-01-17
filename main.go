package main

import (
	"github.com/monstercameron/go-in-one-month/objects"
	"github.com/monstercameron/go-in-one-month/views/components"
	"github.com/monstercameron/go-in-one-month/views/pages"
	"net/http"
)

// todos is a variable that holds a new instance of the Todos struct.
var todos = objects.NewTodos()

func main() {
	// populate the todos object with some todos
	populateTodos()

	// server templ file to / route
	http.HandleFunc("/", getTodo)
	http.HandleFunc("/check", checkTodo)

	// start the server
	http.ListenAndServe(":3000", nil)
}

func populateTodos() {
	todos.Add(objects.Todo{Id: 1, Description: "Buy milk", Checked: false})
	todos.Add(objects.Todo{Id: 2, Description: "Buy eggs", Checked: false})
	todos.Add(objects.Todo{Id: 3, Description: "Buy bread", Checked: false})
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	component := pages.IndexPage("Todos", todos)
	// serve text/html
	w.Header().Set("Content-Type", "text/html")
	// render the component to the response writer
	component.Render(r.Context(), w)
}

func checkTodo(w http.ResponseWriter, r *http.Request) {
	// check if method is post
	if r.Method != http.MethodPost {
		// if not, return a 405
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// get the todo id from the url
	id := r.URL.Query().Get("id")
	// get the todo from the todos object
	todo, todoRef := todos.Get(id)
	// check the todo
	todo.Checked = true
	todoRef.Checked = true
	// return a TodoComponent html
	component := components.TodoComponent(todo)
	// serve text/html
	w.Header().Set("Content-Type", "text/html")
	component.Render(r.Context(), w)
}
