package handler

import (
	"net/http"
)

func About() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/about/" {
			errorHandler(w, http.StatusNotFound)
			return
		}
		renderTemplate(w, "about.html", nil)
	})
	return mux
}
