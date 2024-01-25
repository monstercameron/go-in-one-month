package controller

import (
	"fmt"
	"github.com/monstercameron/go-in-one-month/helpers"
	"github.com/monstercameron/go-in-one-month/objects"
	"github.com/monstercameron/go-in-one-month/views/components"
	"github.com/monstercameron/go-in-one-month/views/pages"
	"net/http"
)

// GetTodo is a handler function that serves the "Todos" page.
// It takes in an http.ResponseWriter and an http.Request as parameters.
// It renders the "Todos" page component to the response writer.
func GetTodo(w http.ResponseWriter, r *http.Request) {
	// Log request details (consider using a logging library for better control)
	helpers.Slogger.Info("Received request", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto)

	if r.Method != http.MethodGet {
		helpers.Slogger.Error("Method not allowed", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get id from url
	id := r.URL.Query().Get("id")
	// if id is not empty
	if id != "" {
		// get todo from todo list
		todo := objects.TodoList.Get(id)
		// if todo is not nil
		if todo != nil {
			// create and render the TodoComponent
			component := components.TodoComponent(todo)
			// serve text/html
			w.Header().Set("Content-Type", "text/html")
			// render the component to the response writer
			component.Render(r.Context(), w)
			// return
			return
		}
	}

	component := pages.IndexPage("My Todo List", objects.TodoList)
	// serve text/html
	w.Header().Set("Content-Type", "text/html")
	// render the component to the response writer
	component.Render(r.Context(), w)
}

// CheckTodo handles the HTTP request for checking a todo item.
// It logs the request details, checks if the method is POST, retrieves the todo ID from the URL,
// retrieves the todo item, updates the todo item, and renders the TodoComponent.
// If any errors occur during the process, appropriate HTTP error responses are sent.
func CheckTodo(w http.ResponseWriter, r *http.Request) {
	// Log request details (consider using a logging library for better control)
	helpers.Slogger.Info("Received request", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto)

	// Check if method is POST
	if r.Method != http.MethodPost {
		helpers.Slogger.Error("Method not allowed", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the todo ID from the URL
	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.Slogger.Error("Todo ID is required", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusBadRequest)
		http.Error(w, "Todo ID is required", http.StatusBadRequest)
		return
	}

	// Retrieve the todo item
	todo := objects.TodoList.Get(id)
	if todo == nil {
		helpers.Slogger.Error("Todo not found", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusNotFound)
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Update the todo item
	todo.Checked = true

	// Create and render the TodoComponent
	component := components.TodoComponent(todo)
	w.Header().Set("Content-Type", "text/html")
	component.Render(r.Context(), w)
}

// UpdateTodos updates a todo item based on the request parameters.
// It logs the request details, checks if the method is POST, retrieves the todo ID from the URL,
// retrieves the todo item, updates the todo item with the provided description,
// and renders the updated todo item using the TodoComponent.
// If any required parameters are missing or the todo item is not found, it returns an appropriate error response.
func UpdateTodos(w http.ResponseWriter, r *http.Request) {
	// Log request details (consider using a logging library for better control)
	helpers.Slogger.Info("Received request", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto)

	// Check if method is POST
	if r.Method != http.MethodPost {
		helpers.Slogger.Info("Method not allowed", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the todo ID from the URL
	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.Slogger.Info("Todo ID is required", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusBadRequest)
		http.Error(w, "Todo ID is required", http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")
	if description == "" {
		helpers.Slogger.Info("Description is required", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusBadRequest)
		http.Error(w, "Description is required", http.StatusBadRequest)
		return
	}

	// Retrieve the todo item
	todo := objects.TodoList.Get(id)
	if todo == nil {
		helpers.Slogger.Info("Todo not found", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusNotFound)
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Update the todo item
	todo.Description = description

	// Create and render the TodoComponent
	component := components.TodoComponent(todo)
	w.Header().Set("Content-Type", "text/html")
	component.Render(r.Context(), w)
}

// DeleteTodo deletes a todo item based on the request parameters.
// It logs the request details, checks if the method is POST, retrieves the todo ID from the URL,
// retrieves the todo item, deletes the todo item,
// and renders the updated todo item using the TodoComponent.
// If any required parameters are missing or the todo item is not found, it returns an appropriate error response.
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Log request details (consider using a logging library for better control)
	helpers.Slogger.Info("Received request", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto)

	// Check if method is POST
	if r.Method != http.MethodDelete {
		helpers.Slogger.Info("Method not allowed", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the todo ID from the URL
	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.Slogger.Info("Todo ID is required", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusBadRequest)
		http.Error(w, "Todo ID is required", http.StatusBadRequest)
		return
	}

	// Retrieve the todo item
	todo := objects.TodoList.Get(id)
	if todo == nil {
		helpers.Slogger.Info("Todo not found", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusNotFound)
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Delete the todo item
	objects.TodoList.Remove(*todo)

	w.Header().Set("Content-Type", "text/html")
	// reurn a 200 OK response
	w.WriteHeader(http.StatusOK)
	// write a message to the response writer
	fmt.Fprintf(w, "")
}

// EditTodo handles the HTTP request for editing a todo item.
// It logs the request details, checks if the method is GET,
// retrieves the todo item based on the provided ID,
// and renders the TodoEditComponent to display the edit form.
// Parameters:
// - w: http.ResponseWriter - the response writer for sending the HTTP response
// - r: *http.Request - the HTTP request received
// Returns: None
func EditTodo(w http.ResponseWriter, r *http.Request) {
	// Log request details (consider using a logging library for better control)
	helpers.Slogger.Info("Received request", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto)

	// Check if method is GET
	if r.Method != http.MethodGet {
		helpers.Slogger.Info("Method not allowed", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	// Get the todo ID from the URL
	id := r.URL.Query().Get("id")
	if id == "" {
		helpers.Slogger.Info("Todo ID is required", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusBadRequest)
		http.Error(w, "Todo ID is required", http.StatusBadRequest)
		return
	}

	// Retrieve the todo item
	todo := objects.TodoList.Get(id)
	if todo == nil {
		helpers.Slogger.Info("Todo not found", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusNotFound)
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	// Create and render the TodoComponent
	component := components.TodoEditComponent(todo)
	w.Header().Set("Content-Type", "text/html")
	component.Render(r.Context(), w)
}

// CreateTodoForm handles the creation of a new todo form.
// It logs the request details, checks if the method is GET,
// validates the description field, creates a new todo object,
// adds it to the todo list, and renders the TodoComponent.
//
// Parameters:
// - w: http.ResponseWriter - the response writer used to write the HTTP response.
// - r: *http.Request - the HTTP request received.
//
// Returns: None
func CreateTodos(w http.ResponseWriter, r *http.Request) {
	// Log request details (consider using a logging library for better control)
	helpers.Slogger.Info("Received request", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto)

	// Check if method is POST
	if r.Method != http.MethodPost {
		helpers.Slogger.Info("Method not allowed", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	description := r.FormValue("description")
	if description == "" {
		helpers.Slogger.Info("Description is required", "method", r.Method, "url", r.URL.String(), "protocol", r.Proto, "status", http.StatusBadRequest)
		http.Error(w, "Description is required", http.StatusBadRequest)
		return
	}

	lastTodo := objects.TodoList.Todos[len(objects.TodoList.Todos)-1]
	id := lastTodo.Id + 1
	var todo = &objects.Todo{Id: id, Description: description, Checked: false}
	objects.TodoList.Add(todo)

	// Create and render the TodoComponent
	component := components.TodoComponent(todo)
	w.Header().Set("Content-Type", "text/html")
	component.Render(r.Context(), w)
}
