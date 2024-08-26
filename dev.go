package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func server(mux *http.ServeMux) {
	server := &http.Server{Addr: ":" + PORT, Handler: mux}
	go func() {
		log.Printf("Server started on port %s", PORT)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %s", err)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	inputChan := make(chan struct{})
	go func() {
		var input string
		for {
			fmt.Scanln(&input)
			if input == "q" {
				inputChan <- struct{}{}
				return
			}
		}
	}()
	select {
	case <-sigChan:
		log.Println("Received interrupt signal. Shutting down...")
	case <-inputChan:
		log.Println("Received 'q'. Shutting down...")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown failed: %s", err)
	} else {
		log.Println("Server stopped gracefully.")
	}
}
