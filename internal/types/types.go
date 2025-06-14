// in internal/types/types.go
package types // Changed from package api

type DashboardResponse struct {
	WelcomeMessage string            `json:"welcomeMessage"`
	Credits        string            `json:"credits"`
	TestGen        TestGenCardData   `json:"testGen"`
	Courses        CoursesCardData   `json:"courses"`
	Quote          QuoteCardData     `json:"quote"`
	News           NewsCardData      `json:"news"`
	QuickLinks     QuickLinksCardData `json:"quickLinks"`
	ProfilePictureURL string `json:"profilePictureUrl"`
	ProfileBannerURL  string `json:"profileBannerUrl"`
}

type TestGenCardData struct {
	Title        string `json:"title"`
	Subject      string `json:"subject"`
	Chapters     string `json:"chapters"`
	LastExam     string `json:"lastExam"`
	PendingExams string `json:"pendingExams"`
	ButtonText   string `json:"buttonText"`
}

type CoursesCardData struct {
	Title             string `json:"title"`
	EnrollmentStatus  string `json:"enrollmentStatus"`
	TodaysFocus       string `json:"todaysFocus"`
	ButtonText        string `json:"buttonText"`
}

type QuoteCardData struct {
	Title      string `json:"title"`
	Quote      string `json:"quote"`
	Author     string `json:"author"`
	ButtonText string `json:"buttonText"`
}

type NewsItem struct {
	Text string `json:"text"`
	Time string `json:"time"`
}

type NewsCardData struct {
	Title string     `json:"title"`
	Items []NewsItem `json:"items"`
}

type QuickLinkItem struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
}

type QuickLinksCardData struct {
	Title string          `json:"title"`
	Links []QuickLinkItem `json:"links"`
}

type OnboardingData struct {
	DisplayName                      string   `json:"displayName"`
	UserRole                         string   `json:"userRole"`
	PreferredWebsiteLanguage         string   `json:"preferred_website_language"`
	PreferredCourseExplanationLanguage string   `json:"preferred_course_explanation_language"`
	PreferredCourseMaterialLanguage    string   `json:"preferred_course_material_language"`
	Major                            string   `json:"major"`
	MajorLevel                       string   `json:"major_level"`
	StudiedSubjects                  []string `json:"studied_subjects"`
	InterestedMajors                 []string `json:"interested_majors"`
	Hobbies                          []string `json:"hobbies"`
	SubscribedToNewsletter           bool     `json:"subscribed_to_newsletter"`
	ReceiveQuotes                    bool     `json:"receive_quotes"`
	Bio                              string   `json:"bio"`
	GithubURL                        string   `json:"github_url"`
	ProfilePictureURL string `json:"profile_picture_url,omitempty"`
	ProfileBannerURL  string `json:"profile_banner_url,omitempty"`
}
