package helpers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// setupSignalHandling sets up signal handling for graceful shutdown of the HTTP server.
// It takes a pointer to an http.Server as input and returns a channel that will be closed
// when a termination signal is received.
func SetupSignalHandling(server *http.Server) <-chan struct{} {
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
