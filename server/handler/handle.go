package handler

import (
	"encoding/json"
	"net/http"
)

func Get() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"Hi": "This is a test GET request from the server/API",
		})
	})
	return mux
}

func Post() *http.ServeMux {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
      "Hi": "This is a test POST request from the server/API",
    })
  })
  return mux
}
