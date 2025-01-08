package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DatabaseDSN string

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, defaulting to environment variables")
	}

	DatabaseDSN = os.Getenv("DATABASE_DSN")
	if DatabaseDSN == "" {
		log.Fatal("DATABASE_DSN is not set in the environment")
	}
}
