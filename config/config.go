package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config *AppConfig

type AppConfig struct {
	DBUrl string
}

func LoadConfig() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, using default")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		dbUrl = "postgres://postgres:postgres@localhost:5432/golang?sslmode=disable"
		log.Println("⚠️  No DB_URL found in .env, using default:", dbUrl)
	}

	Config = &AppConfig{
		DBUrl: dbUrl,
	}
}
