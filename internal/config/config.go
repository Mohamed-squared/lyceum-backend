package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration.
type Config struct {
	Port              string
	DatabaseURL       string
	SupabaseJWTSecret string
}

// Load loads the configuration from environment variables.
// It first tries to load a .env file if present.
func Load() (*Config, error) {
	// Attempt to load .env file, but don't error if it doesn't exist.
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("PORT is not set")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, errors.New("DATABASE_URL is not set")
	}

	supabaseJWTSecret := os.Getenv("SUPABASE_JWT_SECRET")
	if supabaseJWTSecret == "" {
		return nil, errors.New("SUPABASE_JWT_SECRET is not set")
	}

	return &Config{
		Port:              port,
		DatabaseURL:       databaseURL,
		SupabaseJWTSecret: supabaseJWTSecret,
	}, nil
}
