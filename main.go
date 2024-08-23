package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var PORT string

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
}

func main(){
  mux := http.NewServeMux()
  view := DefaultRoutes()
  api := ApiRoutes()
  mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
    host := r.Host
    subdomain := strings.Split(host, ".")[0]
    if subdomain == "api" {
      api.ServeHTTP(w,r)
    } else {
      view.ServeHTTP(w,r)
    }
  })
  fmt.Println("Server running on port:", PORT)
  http.ListenAndServe(":"+PORT, mux)
}
