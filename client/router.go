package client

import (
	"net/http"
	"github.com/MementoMori11723/Tiny-Stack/client/handler"
)

var (
	client = http.NewServeMux()
)

func init() {
	for path, handler := range handler.Routes {
		client.HandleFunc(path, handler)
	}
}

func Client() *http.ServeMux {
	return client
}
