// Path: internal/api/handlers.go
package api

import (
	"encoding/json"
	"log" // Added
	"github.com/Mohamed-squared/lyceum-backend/internal/auth"
	"github.com/Mohamed-squared/lyceum-backend/internal/store"
	"github.com/Mohamed-squared/lyceum-backend/internal/types"
	"net/http"
)

type API struct {
	store *store.Store
}

func New(s *store.Store) *API {
	return &API{store: s}
}

func (a *API) OnboardingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("--- DEBUG: OnboardingHandler started ---")
	userID, ok := r.Context().Value(auth.UserIDKey).(string)
	if !ok {
		log.Println("!!! ONBOARDING ERROR: Could not retrieve user ID from context (ok is false)")
		http.Error(w, "Could not retrieve user ID from context", http.StatusInternalServerError)
		return
	}
	log.Printf("--- DEBUG: Onboarding for User ID: %s", userID)

	var data types.OnboardingData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Printf("!!! ONBOARDING ERROR: Failed to decode JSON body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("--- DEBUG: Decoded onboarding data: %+v", data)

	// Basic validation
	if data.DisplayName == "" {
		log.Println("!!! ONBOARDING ERROR: Validation failed - Display name is empty.")
		http.Error(w, "Display name is required", http.StatusBadRequest)
		return
	}

	log.Println("--- DEBUG: Calling UpdateUserProfile in store...")
	if err := a.store.UpdateUserProfile(r.Context(), userID, data); err != nil {
		log.Printf("!!! ONBOARDING ERROR: Failed to update profile for user %s: %v", userID, err)
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	log.Println("--- DEBUG: Onboarding successful, sending 200 OK")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Profile updated successfully"}`))
}

func (a *API) HandleGetDashboard(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserIDFromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	log.Printf("--- DEBUG: Fetching dashboard data for User ID: %s", userID)
	dashboardData, err := a.store.GetDashboardData(r.Context(), userID)
	if err != nil {
		// Optional: Log the internal error for debugging
		log.Printf("!!! DATABASE ERROR: Failed to get dashboard data for user %s: %v", userID, err)
		http.Error(w, "Failed to retrieve dashboard data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dashboardData)
}

// HealthCheckHandler checks the status of the service and its database connection.
func (a *API) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Check database connection by calling the new Ping method
	dbStatus := "ok"
	if err := a.store.Ping(r.Context()); err != nil {
		log.Printf("!!! HEALTH CHECK ERROR: Database ping failed: %v", err)
		dbStatus = "error"
	}

	// Prepare the response data
	response := map[string]string{
		"status":   "ok",
		"database": dbStatus,
	}

	// As per /docs/api_contract.md, we wrap our response.
	wrappedResponse := map[string]interface{}{
		"success": true,
		"data":    response,
	}

	w.Header().Set("Content-Type", "application/json")
	if dbStatus == "error" {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	json.NewEncoder(w).Encode(wrappedResponse)
}

// CreateCourseHandler handles the creation of a new course.
func (a *API) CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	// Get creator's user ID from the authenticated context
	creatorID, err := auth.GetUserIDFromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Decode the request body
	var req types.CreateCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the input
	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	if req.Visibility != "public" && req.Visibility != "private" {
		http.Error(w, "Visibility must be 'public' or 'private'", http.StatusBadRequest)
		return
	}

	// Call the store to create the course
	newCourse, err := a.store.CreateCourse(r.Context(), creatorID, req)
	if err != nil {
		log.Printf("!!! COURSE CREATION ERROR: %v", err)
		http.Error(w, "Failed to create course", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": newCourse})
}

// GetUserCoursesHandler handles retrieving courses for the authenticated user.
func (a *API) GetUserCoursesHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.GetUserIDFromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	courses, err := a.store.GetCoursesByUserID(r.Context(), userID)
	if err != nil {
		log.Printf("!!! GET USER COURSES ERROR: %v", err)
		http.Error(w, "Failed to retrieve courses", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"success": true, "data": courses})
}
