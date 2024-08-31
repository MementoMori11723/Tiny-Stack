package server

import (
	"encoding/json"
	"net/http"
)

func encoding(data map[string]string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	return nil
}
