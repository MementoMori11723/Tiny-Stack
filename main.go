package main

import (
	"fmt"
	"log"
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
			err := t.Execute(w, returnData(path))
			if err != nil {
				log.Printf("Error executing template for path %s: %v", path, err)
			}
		})
	}
	fmt.Println("Server is running on port:", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
