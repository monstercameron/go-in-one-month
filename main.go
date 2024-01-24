package main

import (
	"context"
	"fmt"
	"github.com/monstercameron/go-in-one-month/objects"
	"github.com/monstercameron/go-in-one-month/views/components"
	"github.com/monstercameron/go-in-one-month/views/pages"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// todos is a variable that holds a new instance of the Todos struct.
var todos = objects.NewTodos()

func main() {
	// Populate the todos object with some todos
	populateTodos()

	// Create a new HTTP server
	server := &http.Server{Addr: ":3000"}

	// Set up your HTTP handlers
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/css"))))
	http.HandleFunc("/", getTodo)
	http.HandleFunc("/check", checkTodo)

	// Setup signal handling and receive shutdown signal
	done := setupSignalHandling(server)

	// Start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server ListenAndServe: %v\n", err)
		}
	}()
	fmt.Println("Server started on port 3000")

	// Block until a shutdown signal is received
	<-done
	fmt.Println("Server gracefully stopped")
}

func populateTodos() {
	todos.Add(objects.Todo{Id: 1, Description: "Buy milk", Checked: false})
	todos.Add(objects.Todo{Id: 2, Description: "Buy eggs", Checked: true})
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
	// print request details
	fmt.Println(r.Method, r.URL, r.Proto)

	// check if method is post
	if r.Method != http.MethodPost {
		// if not, return a 405
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// get the todo id from the url
	id := r.URL.Query().Get("id")
	// get the todo from the todos object
	todo := todos.Get(id)
	// check the todo
	todo.Checked = true
	// return a TodoComponent html
	component := components.TodoComponent(*todo)
	// serve text/html
	w.Header().Set("Content-Type", "text/html")
	component.Render(r.Context(), w)
}

// setupSignalHandling sets up listening for SIGTERM and SIGINT signals
// and initiates a graceful server shutdown when such signals are received.
func setupSignalHandling(server *http.Server) <-chan struct{} {
	done := make(chan struct{})

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("Signal received, starting graceful shutdown")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("HTTP server Shutdown: %v\n", err)
		}

		close(done)
	}()

	return done
}
