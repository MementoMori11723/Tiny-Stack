package client

import (
	"net/http"

	"github.com/MementoMori11723/Tiny-Stack/client/handler"
)

func Client() *http.ServeMux {
  client := http.NewServeMux()
  for path, handler := range handler.Routes {
    client.HandleFunc(path, handler)
  }
  return client
}
