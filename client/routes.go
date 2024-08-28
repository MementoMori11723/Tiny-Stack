package client

import "net/http"

func Client() *http.ServeMux {
  clientMux := http.NewServeMux()
  clientMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
  })
  return clientMux
}
