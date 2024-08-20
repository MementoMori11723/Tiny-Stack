package main

import (
	"fmt"
	"net/http"
	"os"
	"tiny/stack/pages"

	"github.com/joho/godotenv"
)

var PORT string

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
}

func main() {
	for path, handler := range pages.Routes {
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != path {
				http.Redirect(w, r, "/error", http.StatusSeeOther)
			}
			t := handler()
			t.Execute(w, nil)
		})
	}
	fmt.Println("Server is running on port:", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
