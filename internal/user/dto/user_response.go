package dto

import "time"

type UserResponse struct {
	ID              uint   `json:"ID"`
	Username        string `json:"username"`
	Name            string `json:"name"`
	ProfileImageUrl string `json:"profileImageURL"`
	IsFollowed      int    `json:"isFollowed"`
	IsMine          int    `json:"isMine"`
}

type UserDetailResponse struct {
	ID                 uint   `json:"ID"`
	Username           string `json:"username"`
	Name               string `json:"name"`
	ProfileImageUrl    string `json:"profileImageURL"`
	BackgroundImageUrl string `json:"backgroundImageURL"`
	Bio                string `json:"bio"`
	Profession         string `json:"profession"`
	TotalFollower      int    `json:"totalFollower"`
	TotalFollowing     int    `json:"totalFollowing"`
	TotalPost          int    `json:"totalPost"`
	IsFollowed         int    `json:"isFollowed"`
	IsMine             int    `json:"isMine"`
}

type UserUpdateResponse struct {
	Name       string `json:"name"`
	Bio        string `json:"bio"`
	Profession string `json:"profession"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type BriefReportResponse struct {
	ThreadReportID  uint      `json:"threadReportID"`
	UserReportID    uint      `json:"userReportID"`
	CommentReportID uint      `json:"commentReportID"`
	CreatedAt       time.Time `json:"createdAt"`
	ThreadID        uint      `json:"threadID"`
	UserID          uint      `json:"userID"`
	CommentID       uint      `json:"commentID"`
	ReportCategory  string    `json:"reportCategory"`
	Username        string    `json:"username"`
	ProfileImageURL string    `json:"profileImageURL"`
	Type            string    `json:"type"`
}
