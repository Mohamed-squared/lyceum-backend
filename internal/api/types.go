// in internal/api/types.go
package api

type DashboardResponse struct {
	WelcomeMessage string            `json:"welcomeMessage"`
	Credits        string            `json:"credits"`
	TestGen        TestGenCardData   `json:"testGen"`
	Courses        CoursesCardData   `json:"courses"`
	Quote          QuoteCardData     `json:"quote"`
	News           NewsCardData      `json:"news"`
	QuickLinks     QuickLinksCardData `json:"quickLinks"`
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
