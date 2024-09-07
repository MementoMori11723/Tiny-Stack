package server

import (
	"net/http"
	"tiny-stack/server/handler"
)

type route map[string]*http.ServeMux

var (
	mux    = http.NewServeMux()
	routes = route{
    "GET /":handler.Get(),
    "POST /":handler.Post(),
	}
)

func init() {
	for path, handler := range routes {
		mux.Handle(path, handler)
	}
}

func New() *http.ServeMux {
	return mux
}
