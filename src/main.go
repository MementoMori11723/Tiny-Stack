package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	handlers := map[string]func(http.ResponseWriter, *http.Request){
		"/": func(w http.ResponseWriter, r *http.Request) {
			home := Home()
			w.WriteHeader(http.StatusOK)
			home.Render(context.Background(), w)
		},
		"/about": func(w http.ResponseWriter, r *http.Request) {
			about := About()
			w.WriteHeader(http.StatusOK)
			about.Render(context.Background(), w)
		},
		"/404": func(w http.ResponseWriter, r *http.Request) {
			_404 := PageNotFound()
			w.WriteHeader(http.StatusNotFound)
			_404.Render(context.Background(), w)
		},
	}
	fmt.Println("App running at http://localhost:5000")
	http.ListenAndServe(":5000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler, ok := handlers[r.URL.Path]
		if !ok {
			handlers["/404"](w, r)
			return
		}
		handler(w, r)
	}))
}
