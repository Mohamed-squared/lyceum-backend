// Path: internal/store/database.go
package store

import (
	"context"
	"github.com/Mohamed-squared/lyceum-backend/internal/api" // Assuming lyceum is the module name defined in go.mod
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Store struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

// OnboardingData mirrors the JSON payload from the frontend
type OnboardingData struct {
	DisplayName                   string   `json:"displayName"`
	UserRole                      string   `json:"userRole"`
	PreferredWebsiteLanguage      string   `json:"preferred_website_language"`
	PreferredCourseExplanationLanguage string `json:"preferred_course_explanation_language"`
	PreferredCourseMaterialLanguage  string `json:"preferred_course_material_language"`
	Major                         string   `json:"major"`
	MajorLevel                    string   `json:"major_level"`
	StudiedSubjects               []string `json:"studied_subjects"`
	InterestedMajors              []string `json:"interested_majors"`
	Hobbies                       []string `json:"hobbies"`
	SubscribedToNewsletter        bool     `json:"subscribed_to_newsletter"`
	ReceiveQuotes                 bool     `json:"receive_quotes"`
	Bio                           string   `json:"bio"`
	GithubURL                     string   `json:"github_url"`
}

// UpdateUserProfile updates a user's profile after onboarding
func (s *Store) UpdateUserProfile(ctx context.Context, userID string, data OnboardingData) error {
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

func (s *Store) GetDashboardData(userID string) (*api.DashboardResponse, error) {
	// In a future step, we will query the database using the userID.
	// For now, return hardcoded mock data.

	mockData := &api.DashboardResponse{
		WelcomeMessage: "Welcome, Scholar!",
		Credits:        "Scholar's Credits: 250",
		TestGen: api.TestGenCardData{
			Title:        "TestGen Snapshot",
			Subject:      "Artin Abstract Algebra",
			Chapters:     "12/15 Chapters Mastered",
			LastExam:     "Last Exam: 88%",
			PendingExams: "2 Pending PDF Exams",
			ButtonText:   "Go to TestGen Dashboard",
		},
		Courses: api.CoursesCardData{
			Title:            "Courses Snapshot",
			EnrollmentStatus: "3 Courses Enrolled",
			TodaysFocus:      "Focus: Complete Chapter 3 of Quantum Mechanics",
			ButtonText:       "Go to My Courses",
		},
		Quote: api.QuoteCardData{
			Title:      "Quote of the Day",
			Quote:      "The only true wisdom is in knowing you know nothing.",
			Author:     "– Socrates",
			ButtonText: "Refresh",
		},
		News: api.NewsCardData{
			Title: "Lyceum News",
			Items: []api.NewsItem{
				{Text: "New Course Released: Advanced Calculus", Time: "2 hours ago"},
				{Text: "Community Event: Study Group this Friday", Time: "1 day ago"},
			},
		},
		QuickLinks: api.QuickLinksCardData{
			Title: "Quick Links",
			Links: []api.QuickLinkItem{
				{Text: "Generate Test", Icon: "/assets/icons/icon-test.svg"},
				{Text: "Browse Courses", Icon: "/assets/icons/icon-courses.svg"},
			},
		},
	}

	return mockData, nil
}
