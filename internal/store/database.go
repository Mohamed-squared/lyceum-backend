// Path: internal/store/database.go
package store

import (
	"context"
	"fmt"
	"log"

	"github.com/Mohamed-squared/lyceum-backend/internal/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines the database operations
type Store struct {
	db *pgxpool.Pool
}

// New creates a new Store
func New(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

// User represents a user in the database
type User struct {
	ID          string
	Email       string
	DisplayName string
	// Add other fields as necessary
}

// SaveUser saves a user to the database
func (s *Store) SaveUser(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (id, email, display_name)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE SET
			email = EXCLUDED.email,
			display_name = EXCLUDED.display_name;
	`
	_, err := s.db.Exec(ctx, query, user.ID, user.Email, user.DisplayName)
	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

// OnboardingData represents the data collected during onboarding
type OnboardingData struct {
	UserID       string `json:"user_id"`
	DisplayName  string `json:"display_name"`
	Major        string `json:"major"`
	MajorLevel   string `json:"major_level"`
	Pace         string `json:"pace"`
	LearningMode string `json:"learning_mode"`
}

// SaveOnboardingData saves onboarding data to the database
func (s *Store) SaveOnboardingData(ctx context.Context, data *OnboardingData) error {
	// Start a transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			// Rollback if any error occurs
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				log.Printf("failed to rollback transaction: %v (original error: %v)", rbErr, err)
			}
		} else {
			// Commit if everything is fine
			if cErr := tx.Commit(ctx); cErr != nil {
				log.Printf("failed to commit transaction: %v", cErr)
				err = cErr // Propagate commit error
			}
		}
	}()

	// Upsert into profiles table
	profileQuery := `
        INSERT INTO profiles (id, display_name, major, major_level)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (id) DO UPDATE SET
            display_name = EXCLUDED.display_name,
            major = EXCLUDED.major,
            major_level = EXCLUDED.major_level;
    `
	_, err = tx.Exec(ctx, profileQuery, data.UserID, data.DisplayName, data.Major, data.MajorLevel)
	if err != nil {
		return fmt.Errorf("failed to save profile data: %w", err)
	}

	// Upsert into learning_preferences table
	preferencesQuery := `
        INSERT INTO learning_preferences (user_id, pace, learning_mode)
        VALUES ($1, $2, $3)
        ON CONFLICT (user_id) DO UPDATE SET
            pace = EXCLUDED.pace,
            learning_mode = EXCLUDED.learning_mode;
    `
	_, err = tx.Exec(ctx, preferencesQuery, data.UserID, data.Pace, data.LearningMode)
	if err != nil {
		return fmt.Errorf("failed to save learning preferences: %w", err)
	}

	return nil // Error will be handled by defer
}

// GetDashboardData retrieves data needed for the dashboard
func (s *Store) GetDashboardData(ctx context.Context, userID string) (*types.DashboardResponse, error) {
	// Use pointers for nullable columns
	var displayName, major, majorLevel *string // Changed to pointers
	var credits int = 0                       // Default credits, assuming this column is NOT NULL or has a DB default

	// Query to select profile information. Assuming 'major_level' is a column in 'profiles'.
	// If 'credits' comes from the database and can be NULL, it should also be a pointer type (e.g., *int).
	query := `SELECT display_name, major, major_level FROM profiles WHERE id = $1`
	err := s.db.QueryRow(ctx, query, userID).Scan(&displayName, &major, &majorLevel)

	// Handle errors from QueryRow.Scan. pgx.ErrNoRows means no profile was found.
	if err != nil {
		if err == pgx.ErrNoRows {
			// No profile found, proceed with default values for nullable fields.
			// This is a valid scenario, not necessarily an error for dashboard display.
			log.Printf("No profile found for user ID %s. Using default dashboard values.", userID)
		} else {
			// Some other database error occurred
			return nil, fmt.Errorf("database query failed for user ID %s: %w", userID, err)
		}
	}

	// --- Safely handle nil pointers after scanning ---
	var safeDisplayName = "Scholar" // Default if nil or no row
	if displayName != nil {
		safeDisplayName = *displayName
	}

	var safeMajor = "Not specified" // Default if nil or no row
	if major != nil {
		safeMajor = *major
	}

	// Assuming majorLevel might also be used. If so, handle it safely too.
	// var safeMajorLevel = "N/A"
	// if majorLevel != nil {
	// 	safeMajorLevel = *majorLevel
	// }


	// Construct the response
	// Note: The example provided in the issue description for DashboardResponse was quite extensive.
	// This implementation will use the safe values for displayName and major as requested.
	// The other fields (Courses, Quote, News, QuickLinks, and parts of TestGen)
	// were static in the example and will be kept that way here for brevity unless they need to be dynamic.
	responseData := &types.DashboardResponse{
		WelcomeMessage: fmt.Sprintf("Welcome, %s!", safeDisplayName),
		Credits:        fmt.Sprintf("Scholar's Credits: %d", credits), // Using the default/fixed credits value
		TestGen: types.TestGenCardData{
			Title:        "TestGen Snapshot",
			Subject:      safeMajor, // Use the safe, non-nil value
			Chapters:     "0/15 Chapters Mastered", // Static as per example
			LastExam:     "N/A",                    // Static as per example
			PendingExams: "0 Pending PDF Exams",    // Static as per example
			ButtonText:   "Go to TestGen Dashboard",
		},
		Courses: types.CoursesCardData{
			Title:            "Courses Snapshot",
			EnrollmentStatus: "3 Courses Enrolled",                            // Static
			TodaysFocus:      "Focus: Complete Chapter 3 of Quantum Mechanics", // Static
			ButtonText:       "Go to My Courses",
		},
		Quote: types.QuoteCardData{
			Title:      "Quote of the Day",
			Quote:      "The only true wisdom is in knowing you know nothing.", // Static
			Author:     "– Socrates",                                         // Static
			ButtonText: "Refresh",
		},
		News: types.NewsCardData{
			Title: "Lyceum News",
			Items: []types.NewsItem{ // Static
				{Text: "New Course Released: Advanced Calculus", Time: "2 hours ago"},
				{Text: "Community Event: Study Group this Friday", Time: "1 day ago"},
			},
		},
		QuickLinks: types.QuickLinksCardData{
			Title: "Quick Links",
			Links: []types.QuickLinkItem{ // Static
				{Text: "Generate Test", Icon: "/assets/icons/icon-test.svg"},
				{Text: "Browse Courses", Icon: "/assets/icons/icon-courses.svg"},
			},
		},
	}

	return responseData, nil
}
