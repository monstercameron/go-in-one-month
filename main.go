package main

import (
	"fmt"
	"github.com/monstercameron/go-in-one-month/controller"
	"github.com/monstercameron/go-in-one-month/helpers"
	"github.com/monstercameron/go-in-one-month/objects"
	"net/http"
)

func main() {
	// Populate the todos object with some todos
	objects.PopulateTodos()

	// Create a new HTTP server
	server := &http.Server{Addr: ":3000"}

	// Set up your HTTP handlers
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.HandleFunc("/", controller.GetTodo)
	http.HandleFunc("/check", controller.CheckTodo)
	http.HandleFunc("/update", controller.UpdateTodos)
	http.HandleFunc("/edit", controller.EditTodo)
	http.HandleFunc("/delete", controller.DeleteTodo)
	http.HandleFunc("/create", controller.CreateTodos)

	// Setup signal handling and receive shutdown signal
	done := helpers.SetupSignalHandling(server)

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
