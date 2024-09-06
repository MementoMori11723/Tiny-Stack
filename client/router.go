package client

import (
	"github.com/MementoMori11723/Tiny-Stack/client/handler"
	"net/http"
)

var (
	client = http.NewServeMux()
)

func init() {
	for path, handler := range handler.Routes {
		client.HandleFunc(path, handler)
	}
	client.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("client/src/assets"))))
}

func Client() *http.ServeMux {
	return client
}
