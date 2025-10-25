package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"midnite.com/takehometest/handler"
)

var (

	//ServiceName to identify the service
	ServiceName = "MidniteTakeHomeTask"

	//Version of task
	Version string
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/event", handler.AlertHandler)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Channel to listen for termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Channel to wait until graceful shutdown is done
	done := make(chan struct{})

	go func() {
		<-quit
		fmt.Println("Shutting down server...")

		// Create a deadline context for shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Attempt graceful shutdown
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Printf("Error during shutdown: %v\n", err)
		}

		close(done)
	}()

	fmt.Println("Server is ready to handle requests at :8080")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("ListenAndServe error: %v\n", err)
	}

	<-done
	fmt.Println("Server shutdown complete")
}
