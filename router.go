package main

import (
	"net/http"
)

func DefaultRoutes() *http.ServeMux {
	routes := http.NewServeMux()
  routeMapper := map[string]func(w http.ResponseWriter, r *http.Request){
    "/": func(w http.ResponseWriter, r *http.Request) {
      RenderTemplate("home",w)
    },
    "/about": func(w http.ResponseWriter, r *http.Request) {
      RenderTemplate("about",w)
    },
    "/contact": func(w http.ResponseWriter, r *http.Request) {
      RenderTemplate("contact",w)
    },
    "/static/": func(w http.ResponseWriter, r *http.Request) {
      http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
    },
  }
  for path, handler := range routeMapper {
    routes.HandleFunc(path, handler)
  }
  return routes
}

func ApiRoutes() *http.ServeMux {
  routes := http.NewServeMux()
  routeMapper := map[string]func(w http.ResponseWriter, r *http.Request){
    "/": func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "application/json")
      w.Write([]byte("{'message' : 'API endpoint'}"))
    },
    "/admin": func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "application/json")
      w.Write([]byte("{'message':'Admin endpoint'}"))
    },
  }
  for path, handler := range routeMapper {
    routes.HandleFunc(path, handler)
  }
  return routes
}
