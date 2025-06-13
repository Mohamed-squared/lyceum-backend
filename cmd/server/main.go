// Path: cmd/server/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net" // Keep this
	"net/http"
	"os"
	// "time" // This line should be removed

	"github.com/Mohamed-squared/lyceum-backend/internal/api"
	"github.com/Mohamed-squared/lyceum-backend/internal/auth"
	"github.com/Mohamed-squared/lyceum-backend/internal/config"
	"github.com/Mohamed-squared/lyceum-backend/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	// Load config (which might include JWT Secret, ServerPort, etc.)
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// Get Database URL from environment
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Parse the config from the URL
	dbConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
	}

	// This is the key change to force IPv4
	dbConfig.ConnConfig.Config.DialFunc = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, "tcp4", addr)
	}

	// Create a new database connection pool using the modified config
	dbpool, err := pgxpool.NewWithConfig(ctx, dbConfig) // Use existing ctx
	if err != nil {
		log.Fatalf("Unable to create connection pool with custom config: %v\n", err)
	}
	defer dbpool.Close()

	// Ping the database to verify connection
	if err := dbpool.Ping(ctx); err != nil { // Use existing ctx
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	log.Println("Successfully connected to the database.")

	// Setup dependencies (dbStore, apiHandler using the new dbpool)
	dbStore := store.New(dbpool)
	apiHandler := api.New(dbStore)

	// Determine port and listen address (can still use cfg.ServerPort as fallback)
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.ServerPort // cfg might still be useful
	}
	if port == "" {
		port = "8080" // Default fallback
	}
	listenAddr := fmt.Sprintf("0.0.0.0:%s", port)

	// Setup router
	r := chi.NewRouter()

	// Middleware (Logger, CORS - keep existing, ensure cfg is available if CORS needs it)
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*.vercel.app", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Auth Middleware (ensure cfg.SupabaseJWTSecret is available)
	authMiddleware := auth.AuthMiddleware(cfg.SupabaseJWTSecret)

	// API Routes (keep existing)
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(authMiddleware)
		r.Get("/dashboard", apiHandler.HandleGetDashboard)
		r.Post("/onboarding", apiHandler.OnboardingHandler)
	})

	// Start server
	log.Printf("Starting server on %s", listenAddr)
	if err := http.ListenAndServe(listenAddr, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
