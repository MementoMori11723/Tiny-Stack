package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(PORT *string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	*PORT = os.Getenv("PORT")
	if *PORT == "" {
		*PORT = "8080"
	}
}
