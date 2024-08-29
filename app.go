package main

import (
	"log"
	"net/http"

	"github.com/MementoMori11723/Tiny-Stack/config"
	"github.com/MementoMori11723/Tiny-Stack/server"
)

var PORT string

func init() {
	config.LoadEnv(&PORT)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", server.API())
	log.Println("Server started on port " + PORT)
	err := http.ListenAndServe(":"+PORT, mux)
	if err != nil {
		log.Fatal(err)
	}
}
