// Path: internal/store/database.go
package store

import (
	"context"
	"fmt"
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
	var displayName, major, majorLevel string
	var credits int // Assuming credits will be a number, default to 0 if not scanned

	// Query to fetch profile data including credits
	query := `SELECT display_name, major, major_level, credits FROM profiles WHERE id = $1`
	err := s.db.QueryRow(ctx, query, userID).Scan(&displayName, &major, &majorLevel, &credits)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("no profile found for user ID: %s", userID)
		}
		return nil, fmt.Errorf("database query failed: %w", err)
	}

	// Construct the response with a mix of dynamic and static data
	responseData := &types.DashboardResponse{
		WelcomeMessage: fmt.Sprintf("Welcome, %s!", displayName),
		Credits:        fmt.Sprintf("Scholar's Credits: %d", credits),
		TestGen: types.TestGenCardData{
			Title:        "TestGen Snapshot",
			Subject:      major, // Use the real major from the DB
			Chapters:     "0/15 Chapters Mastered", // Static for now
			LastExam:     "Last Exam: N/A", // Static for now
			PendingExams: "0 Pending PDF Exams", // Static for now
			ButtonText:   "Go to TestGen Dashboard",
		},
		Courses: types.CoursesCardData{ // Static data from original file
			Title:            "Courses Snapshot",
			EnrollmentStatus: "3 Courses Enrolled",
			TodaysFocus:      "Focus: Complete Chapter 3 of Quantum Mechanics",
			ButtonText:       "Go to My Courses",
		},
		Quote: types.QuoteCardData{ // Static data from original file
			Title:      "Quote of the Day",
			Quote:      "The only true wisdom is in knowing you know nothing.",
			Author:     "– Socrates",
			ButtonText: "Refresh",
		},
		News: types.NewsCardData{ // Static data from original file
			Title: "Lyceum News",
			Items: []types.NewsItem{
				{Text: "New Course Released: Advanced Calculus", Time: "2 hours ago"},
				{Text: "Community Event: Study Group this Friday", Time: "1 day ago"},
			},
		},
		QuickLinks: types.QuickLinksCardData{ // Static data from original file
			Title: "Quick Links",
			Links: []types.QuickLinkItem{
				{Text: "Generate Test", Icon: "/assets/icons/icon-test.svg"},
				{Text: "Browse Courses", Icon: "/assets/icons/icon-courses.svg"},
			},
		},
	}

	return responseData, nil
}
