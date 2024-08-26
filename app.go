package main

import (
	"github.com/MementoMori11723/Tiny-Stack/database"
	"github.com/MementoMori11723/Tiny-Stack/pages"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var PORT string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", pages.Client())
	mux.Handle("/api/", database.API())
	mux.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./pages/"+r.URL.Path)
	})
	log.Println("Server started on port " + PORT)
	err := http.ListenAndServe(":"+PORT, mux)
	if err != nil {
		log.Fatal(err)
	}
}
