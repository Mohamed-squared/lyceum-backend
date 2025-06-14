package store

import (
	"context"
	"fmt"
	"time"

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
		WHERE id = $1::uuid;
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

// GetDashboardData retrieves data needed for the dashboard
func (s *Store) GetDashboardData(ctx context.Context, userID string) (*types.DashboardResponse, error) {
	var displayName, major, majorLevel *string
	var credits int = 0

	query := `SELECT display_name, major, major_level FROM profiles WHERE id = $1`
	err := s.db.QueryRow(ctx, query, userID).Scan(&displayName, &major, &majorLevel)

	if err != nil && err != pgx.ErrNoRows {
		return nil, fmt.Errorf("database query failed: %w", err)
	}

	var safeDisplayName = "Scholar"
	if displayName != nil {
		safeDisplayName = *displayName
	}

	var safeMajor = "Not specified"
	if major != nil {
		safeMajor = *major
	}

	responseData := &types.DashboardResponse{
		WelcomeMessage: fmt.Sprintf("Welcome, %s!", safeDisplayName),
		Credits:        fmt.Sprintf("Scholar's Credits: %d", credits),
		TestGen: types.TestGenCardData{
			Title:        "TestGen Snapshot",
			Subject:      safeMajor,
			Chapters:     "0/15 Chapters Mastered",
			LastExam:     "N/A",
			PendingExams: "0 Pending PDF Exams",
			ButtonText:   "Go to TestGen Dashboard",
		},
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
