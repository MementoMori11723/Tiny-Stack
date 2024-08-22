package pages

import (
	"html/template"
)

type Route map[string]func() *template.Template

type fileData map[string][]string

var files = fileData{
	"todos": {
		"pages/components/buttons.html",
		"pages/components/dialog.html",
	},
}

var Routes = Route{
	"/":      renderTemplate("index"),
	"/about": renderTemplate("about"),
	"/error": renderTemplate("error"),
	"/todos": renderTemplate("todos"),
}

func renderTemplate(path string) func() *template.Template {
	return func() *template.Template {
		additionalFiles := files[path]
		templates := []string{
			"pages/components/layout.html",
			"pages/" + path + ".html",
		}
		templates = append(templates, additionalFiles...)
		tmpl := template.Must(template.ParseFiles(templates...))
		return tmpl
	}
}
