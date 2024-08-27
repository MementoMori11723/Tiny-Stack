package server

import (
	"encoding/json"
	"net/http"
)

var (
	apiMux = http.NewServeMux()
)

func init() {
	apiMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"})
	})
}

func API() *http.ServeMux {
  return apiMux
}
