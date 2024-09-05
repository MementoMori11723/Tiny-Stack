package server

import (
	"net/http"
	"github.com/MementoMori11723/Tiny-Stack/server/middleware"
)

var routes = map[string]func(http.ResponseWriter, *http.Request){
  "/": sample,
}

func sample(w http.ResponseWriter, r *http.Request) {
  data := map[string]string{"message": "Hello, World!"}
  middleware.Encoding(data, w)
}
