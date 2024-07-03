package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"prsnl/stk/pages"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server, url := &http.Server{Addr: ":" + port, Handler: http.HandlerFunc(runServer)}, "http://localhost:"+port
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("ListenAndServe(): ", err)
		}
	}()
	fmt.Println("Server is running on " + url + "\nPress 'q' and 'Enter' to stop the server")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "q" {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from input: %v", err)
	}
	fmt.Println("Shutting down the server...")
	if err := server.Close(); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Server stopped.")
}

func runServer(w http.ResponseWriter, r *http.Request) {
	handlers := map[string]func(http.ResponseWriter, *http.Request){
		"/": func(w http.ResponseWriter, r *http.Request) {
			home := pages.Home()
			w.WriteHeader(http.StatusOK)
			home.Render(context.Background(), w)
		},
		"/about": func(w http.ResponseWriter, r *http.Request) {
			about := pages.About()
			w.WriteHeader(http.StatusOK)
			about.Render(context.Background(), w)
		},
		"/api": func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{message: "Hello, World!"}`)
		},
		"/404": func(w http.ResponseWriter, r *http.Request) {
			_404 := pages.PageNotFound()
			w.WriteHeader(http.StatusNotFound)
			_404.Render(context.Background(), w)
		},
	}
	handler, ok := handlers[r.URL.Path]
	if !ok {
		handlers["/404"](w, r)
		return
	}
	handler(w, r)
}
