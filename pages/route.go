package pages

import (
	"html/template"
	"net/http"
)

var (
	clientMux = http.NewServeMux()
  componentsDir = "pages/components/"

	Routes    = Route{
		"/":       returnTemplate("index"),
		"/about/": returnTemplate("about"),
		"/error/": returnTemplate("error"),
	}

	Componentsfiles = map[string][]string{
		"about": {
			componentsDir+"contacts.html",
		},
	}
) 

type Route map[string]func(http.ResponseWriter, *http.Request)

func init() {
	for path, handler := range Routes {
		clientMux.HandleFunc(path, handler)
	}
}

func returnFiles(file string) []string {
  return append(
    []string{
      "pages/components/layouts.html",
      "pages/"+file+".html",
    },
    Componentsfiles[file]...,
  )
}

func returnTemplate(page string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(returnFiles(page)...))
		tmpl.Execute(w, nil)
	}
}

func Client() *http.ServeMux {
	return clientMux
}
