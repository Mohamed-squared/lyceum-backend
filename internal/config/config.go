// Path: internal/config/config.go
package config

import (
	"errors"
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

	cfg := &Config{
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		SupabaseJWTSecret:   os.Getenv("SUPABASE_JWT_SECRET"),
		ServerPort:        os.Getenv("SERVER_PORT"),
		SupabaseServiceKey: os.Getenv("SUPABASE_SERVICE_KEY"),
	}

	if cfg.SupabaseJWTSecret == "" {
		return nil, errors.New("SUPABASE_JWT_SECRET environment variable not set")
	}

	return cfg, nil
}
