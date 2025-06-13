// Path: internal/api/handlers.go
package api

import (
	"encoding/json"
	"github.com/Mohamed-squared/lyceum-backend/internal/auth"
	"github.com/Mohamed-squared/lyceum-backend/internal/store"
	"net/http"
)

type API struct {
	store *store.Store
}

func New(s *store.Store) *API {
	return &API{store: s}
}

func (a *API) OnboardingHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(auth.UserIDKey).(string)
	if !ok {
		http.Error(w, "Could not retrieve user ID from context", http.StatusInternalServerError)
		return
	}

	var data store.OnboardingData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if data.DisplayName == "" {
		http.Error(w, "Display name is required", http.StatusBadRequest)
		return
	}

	if err := a.store.UpdateUserProfile(r.Context(), userID, data); err != nil {
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Profile updated successfully"}`))
}

func (a *API) HandleGetDashboard(w http.ResponseWriter, r *http.Request) {
	// In a future step, we will get the userID from the auth middleware.
	// For now, we can use a placeholder ID.
	userID := "placeholder_user_id" // Or retrieve from context if auth.UserIDKey is already set by a middleware for this path

	dashboardData, err := a.store.GetDashboardData(userID) // Corrected: a.store instead of s.Store
	if err != nil {
		http.Error(w, "Failed to retrieve dashboard data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dashboardData)
}
