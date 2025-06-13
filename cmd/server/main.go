// Path: cmd/server/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Mohamed-squared/lyceum-backend/internal/api"
	"github.com/Mohamed-squared/lyceum-backend/internal/auth"
	"github.com/Mohamed-squared/lyceum-backend/internal/config"
	"github.com/Mohamed-squared/lyceum-backend/internal/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors" // Ensure this import is present
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Connect to database
	dbpool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v
", err)
	}
	defer dbpool.Close()

	log.Println("Successfully connected to the database.")

	// Setup dependencies
	dbStore := store.New(dbpool)
	apiHandler := api.New(dbStore) // Renamed for clarity from the issue's 'api'

	// Determine port and listen address
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.ServerPort
	}
	if port == "" {
		port = "8080" // Default port
	}
	listenAddr := fmt.Sprintf("0.0.0.0:%s", port)

	// Setup router
	router := chi.NewRouter() // Renamed from 'r' to 'router' to match example

	// Middleware
	router.Use(middleware.Logger) // Existing logger

	// Configure CORS Middleware
	// This MUST come before the AuthMiddleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*.vercel.app", "http://localhost:3000"}, // Allow your frontend domains
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any major browsers
	}))

	// Setup the API with authentication
	// Renamed apiHandler from 'api' to avoid conflict with package name if 'api.New' was used directly
	// authMiddleware instance
	authMiddleware := auth.AuthMiddleware(cfg.SupabaseJWTSecret)

	// Apply AuthMiddleware to a group of routes
	router.Route("/api/v1", func(r chi.Router) {
		r.Use(authMiddleware) // Apply AuthMiddleware to all /api/v1 routes

		// Your authenticated routes go here
		r.Post("/onboarding", apiHandler.OnboardingHandler)
		r.Get("/dashboard", apiHandler.HandleGetDashboard) // Now also protected by authMiddleware
	})

	// Start server
	log.Printf("Starting server on %s", listenAddr)
	if err := http.ListenAndServe(listenAddr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
