package pages

import (
	"html/template"
)

type Route map[string]func() *template.Template

func renderTemplate(path string) *template.Template {
	return template.Must(template.ParseFiles(path))
}

var Routes = Route{
	"/": func() *template.Template {
		return renderTemplate("pages/index.html")
	},
	"/about": func() *template.Template {
		return renderTemplate("pages/about.html")
	},
	"/error": func() *template.Template {
		return renderTemplate("pages/error.html")
	},
  "/todos": func() *template.Template {
    return renderTemplate("pages/todos.html")
  },
}
