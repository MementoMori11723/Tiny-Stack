package main

import (
	"fmt"
	"html/template"
	"net/http"
	"tiny/stack/database"
)

var (
	viewDir      = "pages/"
	componentDir = viewDir + "components/"
	layout       = componentDir + "layouts.html"

	pages = map[string]string{
		"home":    viewDir + "index.html",
		"about":   viewDir + "about.html",
		"contact": viewDir + "contact.html",
	}
	defaultComponents = []string{
		componentDir + "header.html",
		componentDir + "footer.html",
	}
	dataProcessor = map[string]func() database.ReturnType{
		"contact": database.Fetch,
	}
)

func processTemplate(page string) []string {
	files := []string{
		layout,
		pages[page],
	}
	return append(files, defaultComponents...)
}

func CreateTemplate(page string) *template.Template {
	templates := processTemplate(page)
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		fmt.Printf("Error parsing template %s: %v\n", page, err)
		return nil
	}
	return tmpl
}

func RenderTemplate(page string, w http.ResponseWriter) {
	tmpl := CreateTemplate(page)
	if tmpl == nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println("Template creation failed for page:", page)
		return
	}
	var err error
	if dataFunc, ok := dataProcessor[page]; ok {
		data := dataFunc()
		err = tmpl.Execute(w, data)
	} else {
		err = tmpl.Execute(w, nil)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("Template execution error for page %s: %v\n", page, err)
	}
}
