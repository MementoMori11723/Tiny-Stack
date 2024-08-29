package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MementoMori11723/Tiny-Stack/client"
	"github.com/MementoMori11723/Tiny-Stack/config"
	"github.com/MementoMori11723/Tiny-Stack/server"
)

var PORT string

func init() {
	config.LoadEnv(&PORT)
}

func main() {
	go func() {
		mux := http.NewServeMux()
    mux.Handle("/", client.Client())
		mux.Handle("/api", server.API())
		log.Println("Server started on port " + PORT)
		err := http.ListenAndServe(":"+PORT, mux)
		if err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Scanln()
}
