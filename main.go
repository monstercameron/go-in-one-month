package main

import (
	"context"
	"fmt"
	"github.com/monstercameron/go-in-one-month/controller"
	"github.com/monstercameron/go-in-one-month/objects"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Populate the todos object with some todos
	objects.PopulateTodos()

	// Create a new HTTP server
	server := &http.Server{Addr: ":3000"}

	// Set up your HTTP handlers
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/css"))))
	http.HandleFunc("/", controller.GetTodo)
	http.HandleFunc("/check", controller.CheckTodo)
	http.HandleFunc("/description", controller.UpdateTodos)
	http.HandleFunc("/remove", controller.DeleteTodo)

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

// setupSignalHandling sets up signal handling for graceful shutdown of the HTTP server.
// It takes a pointer to an http.Server as input and returns a channel that will be closed
// when a termination signal is received.
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
