// Path: internal/config/config.go
package config

import (
	"os"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	DatabaseURL       string
	SupabaseJWTSecret   string
	ServerPort        string
	SupabaseServiceKey string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		SupabaseJWTSecret:   os.Getenv("SUPABASE_JWT_SECRET"),
		ServerPort:        os.Getenv("SERVER_PORT"),
		SupabaseServiceKey: os.Getenv("SUPABASE_SERVICE_KEY"),
	}, nil
}
