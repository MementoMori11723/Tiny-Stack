package server

import (
	"net/http"
)

var (
	apiMux = http.NewServeMux()
)

func init() {
  for path, handler := range routes {
	  apiMux.HandleFunc(path, handler)
  }
}

func API() *http.ServeMux {
  return apiMux
}
