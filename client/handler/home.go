package handler

import (
	"net/http"
	"text/template"
)

var (
	pageDir = "client/src/"
	layout  = "layout.html"
)

func renderTemplate(w http.ResponseWriter, s string, d interface{}) {
	templates, err := template.ParseFiles(pageDir+layout, pageDir+s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = templates.ExecuteTemplate(w, layout, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func errorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	msg := map[int]string{
		http.StatusNotFound:            "Page Not Found",
		http.StatusInternalServerError: "Internal Server Error",
		http.StatusServiceUnavailable:  "Service Unavailable",
		http.StatusUnauthorized:        "Unauthorized",
		http.StatusForbidden:           "Forbidden",
		http.StatusBadRequest:          "Bad Request",
	}
	renderTemplate(w, "error.html", struct {
		Status int
		Msg    string
	}{
		Status: status,
		Msg:    msg[status],
	})
}

func Home() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			errorHandler(w, http.StatusNotFound)
			return
		}
		renderTemplate(w, "index.html", nil)
	})
	return mux
}
