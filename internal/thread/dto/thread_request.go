package dto

type ThreadRequest struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	CommunityID uint   `json:"communityID"`
}

type ThreadReportRequest struct {
	ReportCategoryID uint `json:"reportCategoryID"`
	ThreadID         uint `json:"threadID"`
	UserID           uint `json:"userID"`
}
