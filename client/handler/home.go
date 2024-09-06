package handler

import (
	"html/template"
	"net/http"
)

type route map[string]func(http.ResponseWriter, *http.Request)

var (
	Routes  route
	pageDir = "client/src/pages/"
	layout  = "layouts.html"
)

func init() {
	Routes = route{
		"/":         HomeHandler,
		"/about/":   AboutHandler,
		"/error/":   ErrorHandler,
		"/example/": ExampleHandler,
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := Routes[r.URL.Path]; !ok {
		http.Redirect(w, r, "/error/", http.StatusSeeOther)
		return
	}
	home, err := template.ParseFiles(pageDir+layout, pageDir+"index.html")
	if err != nil {
		w.Write([]byte("Error reading file"))
		return
	}
	home.Execute(w, nil)
}
