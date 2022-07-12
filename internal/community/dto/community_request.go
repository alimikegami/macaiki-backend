package dto

type CommunityRequest struct {
	ID          uint   `json:"ID"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"  validate:"required"`
}

type CommunityModeratorRequest struct {
	UserID      uint `json:"userID"`
	CommunityID uint `json:"communityID"`
}

type CommunityReportRequest struct {
	ReportCategoryID uint `json:"reportCategoryID" validate:"required"`
	CommunityID      uint `validate:"required"`
}

type ReportRequest struct {
	ThreadReportID    uint `json:"threadReportsID"`
	CommunityReportID uint `json:"communityReportsID"`
	CommentReportID   uint `json:"commentReportsID"`
}
