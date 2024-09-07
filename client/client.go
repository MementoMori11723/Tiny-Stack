package client

import (
	"net/http"
	"tiny-stack/client/handler"
)

type route map[string]*http.ServeMux

var (
	mux    = http.NewServeMux()
	routes = route{
		"/":       handler.Home(),
		"/about/": handler.About(),
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
