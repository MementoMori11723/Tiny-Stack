package main

import (
	"fmt"
	"log"
	"net/http"
	"tiny-stack/client"
	"tiny-stack/server"
  "tiny-stack/config"
)

var (
	PORT string
)

func init() {
	config.LoadEnv(&PORT)
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
