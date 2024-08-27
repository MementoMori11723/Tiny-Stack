package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MementoMori11723/Tiny-Stack/client"
	"github.com/MementoMori11723/Tiny-Stack/server"
	"github.com/joho/godotenv"
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
	mux.Handle("/",client.Client())
	mux.Handle("/api/", server.API())
	log.Println("Server started on port " + PORT)
	err := http.ListenAndServe(":"+PORT, mux)
	if err != nil {
		log.Fatal(err)
	}
}
