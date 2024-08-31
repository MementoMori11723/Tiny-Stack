package server

import "net/http"

var routes = map[string]func(http.ResponseWriter, *http.Request){
  "/": sample,
}

func sample(w http.ResponseWriter, r *http.Request) {
  data := map[string]string{"message": "Hello, World!"}
  encoding(data, w)
}
