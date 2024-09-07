package main

import (
	"fmt"
	"log"
	"net/http"
	"tiny-stack/client"
	"tiny-stack/server"
)

var (
	PORT string
)

func init() {
	PORT = "8080"
}

func main() {
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/", client.New())
		mux.Handle("/api/", server.New())
		log.Println("Listening on port " + PORT)
		http.ListenAndServe(":"+PORT, mux)
	}()
	fmt.Scanln()
}
