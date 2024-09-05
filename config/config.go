package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func LoadDatabasePath(DATABASE_PATH *string) {
	*DATABASE_PATH = os.Getenv("DATABASE_PATH")
	if *DATABASE_PATH == "" {
		*DATABASE_PATH = "./server/database/database.db"
	}
}

func LoadPort(PORT *string) {
	*PORT = os.Getenv("PORT")
	if *PORT == "" {
		*PORT = "8080"
	}
}
