package dto

type ThreadRequest struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	CommunityID uint   `json:"communityID"`
}

type ThreadReportRequest struct {
	ReportCategoryID uint `json:"reportCategoryID"`
	ThreadID         uint
	UserID           uint
}

type CommentReportRequest struct {
	ReportCategoryID uint `json:"reportCategoryID"`
	CommentID        uint
	UserID           uint
}

type SavedThreadRequest struct {
	UserID   uint
	ThreadID uint
}
