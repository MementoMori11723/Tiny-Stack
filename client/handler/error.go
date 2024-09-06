package handler

import (
	"html/template"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
  error, err := template.ParseFiles(pageDir+layout, pageDir+"error.html")
  if err != nil {
    w.Write([]byte("Error reading file"))
    return
  }
  error.Execute(w, nil)
}
