package main

import (
	"github.com/monstercameron/go-in-one-month/views/pages/index"
	"github.com/monstercameron/go-in-one-month/objects"
	"net/http"
)

func main() {

	// todos is a variable that holds a new instance of the Todos struct.
	todos := objects.NewTodos()
	todos.Add(objects.Todo{Id: 1, Description: "Buy milk", Checked: false})
	todos.Add(objects.Todo{Id: 2, Description: "Buy eggs", Checked: false})
	todos.Add(objects.Todo{Id: 3, Description: "Buy bread", Checked: false})

	// server templ file to / route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := views.IndexPage("Todos", todos)
		// serve text/html
		w.Header().Set("Content-Type", "text/html")
		// render the component to the response writer
		component.Render(r.Context(), w)
	})

	http.ListenAndServe(":3000", nil)
}
