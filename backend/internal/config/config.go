package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func Load() (*Config, error) {
	_ = godotenv.Load() // Ignore error if .env file is not present

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Default to local mysql if not set
		// Format: user:password@tcp(host:port)/dbname?parseTime=true
		dbURL = "root:root@tcp(localhost:3306)/keepsy?parseTime=true"
	}

	return &Config{
		Port:        port,
		DatabaseURL: dbURL,
	}, nil
}
