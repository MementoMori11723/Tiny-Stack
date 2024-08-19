package main

import (
	"fmt"
	"tiny/stack/pages"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }
  PORT := os.Getenv("PORT")
  if PORT == "" {
    fmt.Println("PORT not found in .env file, using default port 8080")
    PORT = "8080"
  }
	for path, handler := range pages.Routes {
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, handler())
		})
	}
  fmt.Println("Server is running on port:", PORT)
  http.ListenAndServe(":"+PORT, nil)
}
