package handler

import (
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
  about, err := template.ParseFiles(pageDir+layout, pageDir+"about.html")
  if err != nil {
    w.Write([]byte("Error reading file"))
    return
  }
  about.Execute(w, nil)
}
