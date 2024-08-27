package client

import (
	"html/template"
	"net/http"
)

var (
	pagesDir      = "pages/"
	componentsDir = "components/"
  layoutFile    = "layouts.html"

	routes = Route{
		"/":       homePage,
		"/about/": aboutPage,
		"/error/": errorPage,
	}
)

type Route map[string]func(http.ResponseWriter, *http.Request)

func homePage(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles(pagesDir+componentsDir+layoutFile, pagesDir+"index.html"))
  tmpl.Execute(w, nil)
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles(pagesDir+componentsDir+layoutFile, pagesDir+"about.html", pagesDir+componentsDir+"contacts.html"))
  tmpl.Execute(w, nil)
}

func errorPage(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles(pagesDir+componentsDir+layoutFile, pagesDir+"error.html"))
  tmpl.Execute(w, nil)
}

func Client() *http.ServeMux {
  clientMux := http.NewServeMux()
	for path, handler := range routes {
		clientMux.HandleFunc(path, handler)
	}
	return clientMux
}
