package pages

import (
	"html/template"
	"net/http"
)

var (
	pagesDir      = "pages/"
	componentsDir = "components/"
  layoutFile    = "layouts.html"

	routes = Route{
		"/":       generateTemplate("index.html"),
		"/about/": generateTemplate("about.html"),
		"/error/": generateTemplate("error.html"),
	}

	componentFiles = map[string][]string{
		"about": {
			pagesDir + componentsDir + "contacts.html",
		},
	}
)

type Route map[string]func(http.ResponseWriter, *http.Request)

func returnFiles(file string) []string {
	return append(
		[]string{
			pagesDir + componentsDir + layoutFile,
			pagesDir + file,
		},
		componentFiles[file]...,
	)
}

// found a better way to handle templates!
func generateTemplate(page string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(returnFiles(page)...)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
		err = tmpl.Execute(w, nil)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
	}
}

func Client() *http.ServeMux {
  clientMux := http.NewServeMux()
	for path, handler := range routes {
		clientMux.HandleFunc(path, handler)
	}
	return clientMux
}
