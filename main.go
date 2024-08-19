package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  PORT := os.Getenv("PORT")
  if PORT == "" {
    PORT = "8080"
  }
  e := echo.New()
  e.GET("/", func(c echo.Context) error {
    return c.String(200, "Hello, World!")
  })
  e.Logger.Fatal(e.Start(":"+PORT))
}
