package handler

import (
	"context"
	"net/http"

	"github.com/MementoMori11723/Tiny-Stack/client/pages"
)

var Routes map[string]func(http.ResponseWriter, *http.Request)

func init() {
	Routes = map[string]func(http.ResponseWriter, *http.Request){
		"/":       HomeHandler,
		"/about/": AboutHandler,
		"/error/": ErrorHandler,
    "/example/": ExampleHandler,
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := Routes[r.URL.Path]; !ok {
		http.Redirect(w, r, "/error/", http.StatusSeeOther)
		return
	}
	home := pages.Home()
	home.Render(context.Background(), w)
}
