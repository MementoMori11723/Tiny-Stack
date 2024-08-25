package pages

import (
	"html/template"
	"net/http"
)

var (
	clientMux = http.NewServeMux()
	Routes    = Route{
		"/":       returnTemplate("index"),
		"/about/": returnTemplate("about"),
		"/error/": returnTemplate("error"),
	}
	Componentsfiles = map[string][]string{
		"about": {
			"contacts",
		},
	}
)

type Handler func(http.ResponseWriter, *http.Request)

type Route map[string]Handler

func init() {
	for path, handler := range Routes {
		clientMux.HandleFunc(path, handler)
	}
}

func returnFiles(file string) []string {
	var Templates []string
	tmpls, _ := Componentsfiles[file]
	if len(tmpls) <= 0 {
		return []string{
			"pages/components/layouts.html",
			"pages/" + file + ".html",
		}
	} else {
		for _, tmpl := range tmpls {
			Templates = append(Templates, "pages/components/"+tmpl+".html")
		}
		return append([]string{
			"pages/components/layouts.html",
			"pages/" + file + ".html",
		}, Templates...)
	}
}

func returnTemplate(page string) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles(returnFiles(page)...))
		tmpl.Execute(w, nil)
	}
}

func Client() *http.ServeMux {
	return clientMux
}
