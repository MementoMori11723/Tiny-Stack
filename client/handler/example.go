package handler

import (
	"html/template"
	"net/http"
)

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	example, err := template.ParseFiles(pageDir+layout, pageDir+"example.html")
	if err != nil {
		w.Write([]byte("Error reading file"))
		return
	}
	example.Execute(w, nil)
}
