// Path: cmd/server/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"github.com/Mohamed-squared/lyceum-backend/internal/api"
	"github.com/Mohamed-squared/lyceum-backend/internal/auth"
	"github.com/Mohamed-squared/lyceum-backend/internal/config"
	"github.com/Mohamed-squared/lyceum-backend/internal/store"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()

	log.Println("Successfully connected to the database.")

	// Setup dependencies
	dbStore := store.New(dbpool)
	apiHandler := api.New(dbStore)

	// Determine port and listen address
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.ServerPort
	}
	if port == "" {
		port = "8080"
	}
	listenAddr := fmt.Sprintf("0.0.0.0:%s", port)

	// Setup router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	// Ensure CORS options match the issue's example EXACTLY
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*.vercel.app", "http://localhost:3000"}, // FROM ISSUE
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	authMiddleware := auth.AuthMiddleware(cfg.SupabaseJWTSecret) // Define it once

	// API Routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(authMiddleware) // Apply to the whole group

		// Previously public dashboard route, now authenticated
		r.Get("/dashboard", apiHandler.HandleGetDashboard)

		// Previously in a sub-group, now directly under /api/v1 with auth
		r.Post("/onboarding", apiHandler.OnboardingHandler)
	})

	// Start server
	log.Printf("Starting server on %s", listenAddr)
	if err := http.ListenAndServe(listenAddr, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
