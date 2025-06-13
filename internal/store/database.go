// Path: internal/store/database.go
package store

import (
	"context"
	"database/sql" // Add this
	"fmt"
	"log" // Add this
	"github.com/Mohamed-squared/lyceum-backend/internal/types" // Assuming lyceum is the module name defined in go.mod
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Store struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

// UpdateUserProfile updates a user's profile after onboarding
func (s *Store) UpdateUserProfile(ctx context.Context, userID string, data types.OnboardingData) error {
	query := `
		UPDATE public.profiles
		SET
			display_name = $2, user_role = $3, preferred_website_language = $4,
			preferred_course_explanation_language = $5, preferred_course_material_language = $6,
			major = $7, major_level = $8, studied_subjects = $9, interested_majors = $10,
			hobbies = $11, subscribed_to_newsletter = $12, receive_quotes = $13, bio = $14,
			github_url = $15, has_completed_onboarding = TRUE, updated_at = $16
		WHERE id = $1;
	`
	_, err := s.db.Exec(ctx, query,
		userID, data.DisplayName, data.UserRole, data.PreferredWebsiteLanguage,
		data.PreferredCourseExplanationLanguage, data.PreferredCourseMaterialLanguage,
		data.Major, data.MajorLevel, data.StudiedSubjects, data.InterestedMajors,
		data.Hobbies, data.SubscribedToNewsletter, data.ReceiveQuotes, data.Bio,
		data.GithubURL, time.Now(),
	)
	return err
}

func (s *Store) GetDashboardData(ctx context.Context, userID string) (*types.DashboardResponse, error) {
	// 1. Use nullable types for all variables that can be NULL in the DB
	var displayName, major, majorLevel sql.NullString
	var credits sql.NullInt32

	// 2. The query remains the same
	query := `SELECT display_name, major, major_level, credits FROM profiles WHERE id = $1`

	// 3. Scan into the nullable types
	err := s.db.QueryRow(ctx, query, userID).Scan(&displayName, &major, &majorLevel, &credits)

	if err != nil {
		// This detailed log will now appear in your Railway Deploy Logs
		log.Printf("DATABASE ERROR for user '%s': %v", userID, err)
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("no profile found for user ID: %s", userID)
		}
		return nil, fmt.Errorf("database query failed") // Keep generic error for client
	}

	// 4. Safely extract values from nullable types, providing defaults
	finalDisplayName := "Scholar"
	if displayName.Valid {
		finalDisplayName = displayName.String
	}

	finalMajor := "Undeclared"
	if major.Valid {
		finalMajor = major.String
	}
	// majorLevel is read but not used in the example response construction,
	// but if it were, it would be:
	// finalMajorLevel := "" // Or some default
	// if majorLevel.Valid {
	//     finalMajorLevel = majorLevel.String
	// }

	finalCredits := 0
	if credits.Valid {
		finalCredits = int(credits.Int32)
	}

	// 5. Construct the final response with safe data
	responseData := &types.DashboardResponse{
		WelcomeMessage: fmt.Sprintf("Welcome, %s!", finalDisplayName),
		Credits:        fmt.Sprintf("Scholar's Credits: %d", finalCredits),
		TestGen: types.TestGenCardData{
			Title:        "TestGen Snapshot",
			Subject:      finalMajor, // Use the processed major
			Chapters:     "0/15 Chapters Mastered",
			LastExam:     "Last Exam: N/A",
			PendingExams: "0 Pending PDF Exams",
			ButtonText:   "Go to TestGen Dashboard",
		},
		// IMPORTANT: Fill these from the existing GetDashboardData function in the file
		Courses: types.CoursesCardData{
			Title:            "Courses Snapshot",
			EnrollmentStatus: "3 Courses Enrolled",
			TodaysFocus:      "Focus: Complete Chapter 3 of Quantum Mechanics",
			ButtonText:       "Go to My Courses",
		},
		Quote: types.QuoteCardData{
			Title:      "Quote of the Day",
			Quote:      "The only true wisdom is in knowing you know nothing.",
			Author:     "– Socrates",
			ButtonText: "Refresh",
		},
		News: types.NewsCardData{
			Title: "Lyceum News",
			Items: []types.NewsItem{
				{Text: "New Course Released: Advanced Calculus", Time: "2 hours ago"},
				{Text: "Community Event: Study Group this Friday", Time: "1 day ago"},
			},
		},
		QuickLinks: types.QuickLinksCardData{
			Title: "Quick Links",
			Links: []types.QuickLinkItem{
				{Text: "Generate Test", Icon: "/assets/icons/icon-test.svg"},
				{Text: "Browse Courses", Icon: "/assets/icons/icon-courses.svg"},
			},
		},
	}

	return responseData, nil
}
