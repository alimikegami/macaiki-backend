package dto

import (
	"macaiki/internal/user/dto"
	"time"
)

type CommunityDetailResponse struct {
	ID                          uint   `json:"ID"`
	Name                        string `json:"name"`
	CommunityImageUrl           string `json:"communityImageUrl"`
	CommunityBackgroundImageUrl string `json:"communityBackgroundImageUrl"`
	Description                 string `json:"description"`
	IsFollowed                  int    `json:"isFollowed"`
	IsModerator                 int    `json:"isModerator"`
}

type CommunityUpdateResponse struct {
	ID          uint   `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CommunityAboutResponse struct {
	ID                          uint               `json:"ID"`
	Name                        string             `json:"name"`
	CommunityImageUrl           string             `json:"communityImageUrl"`
	CommunityBackgroundImageUrl string             `json:"communityBackgroundImageUrl"`
	Description                 string             `json:"description"`
	IsFollowed                  int                `json:"isFollowed"`
	IsModerator                 int                `json:"isModerator"`
	TotalModerator              int                `json:"totalModerators"`
	TotalFollower               int                `json:"totalFollowers"`
	Moderator                   []dto.UserResponse `json:"moderators"`
}

type BriefReportResponse struct {
	ThreadReportsID     uint      `json:"threadReportsID"`
	CommunityReportsID  uint      `json:"communityReportsID"`
	CommentReportsID    uint      `json:"commentReportsID"`
	CreatedAt           time.Time `json:"createdAt"`
	UserID              uint      `json:"userID"`
	ThreadID            uint      `json:"threadID"`
	CommunityReportedID uint      `json:"communityReportedID"`
	CommentID           uint      `json:"commentID"`
	ReportCategory      string    `json:"reportCategory"`
	Username            string    `json:"username"`
	ProfileImageURL     string    `json:"profileImageURL"`
	Type                string    `json:"type"`
}
