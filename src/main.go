package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func runServer(w http.ResponseWriter, r *http.Request) {
	handlers := map[string]func(http.ResponseWriter, *http.Request){
		"/": func(w http.ResponseWriter, r *http.Request) {
			home := Home()
			w.WriteHeader(http.StatusOK)
			home.Render(context.Background(), w)
		},
		"/about": func(w http.ResponseWriter, r *http.Request) {
			about := About()
			w.WriteHeader(http.StatusOK)
			about.Render(context.Background(), w)
		},
		"/api": func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{message: "Hello, World!"}`)
		},
		"/404": func(w http.ResponseWriter, r *http.Request) {
			_404 := PageNotFound()
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

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Printf("Failed to open browser: %v", err)
	}
}

func main() {
	server, url := &http.Server{Addr: ":5000", Handler: http.HandlerFunc(runServer)}, "http://localhost:5000"
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("ListenAndServe(): ", err)
		}
	}()
	fmt.Println("Server is running on" + url + "\nPress 'q' and 'Enter' to stop the server")
	openBrowser(url)
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
