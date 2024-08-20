package pages

import (
	"html/template"
)

type Route map[string]func() *template.Template

func renderTemplate(path string) func() *template.Template {
	return func() *template.Template {
    return template.Must(template.ParseFiles(path))
  }
}

var Routes = Route{
  "/": renderTemplate("pages/index.html"),
  "/about": renderTemplate("pages/about.html"),
  "/error": renderTemplate("pages/error.html"),
  "/todos": renderTemplate("pages/todos.html"),
}
