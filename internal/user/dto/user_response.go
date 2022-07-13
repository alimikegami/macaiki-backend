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
	ThreadReportsID     uint      `json:"threadReportsID"`
	UserReportsID       uint      `json:"userReportsID"`
	CommentReportsID    uint      `json:"commentReportsID"`
	CommunityReportsID  uint      `json:"communityReportsID"`
	CreatedAt           time.Time `json:"createdAt"`
	ThreadID            uint      `json:"threadID"`
	UserID              uint      `json:"userID"`
	CommentID           uint      `json:"commentID"`
	CommunityReportedIT uint      `json:"communityReportedID"`
	ReportCategory      string    `json:"reportCategory"`
	Username            string    `json:"username"`
	ProfileImageURL     string    `json:"profileImageURL"`
	Type                string    `json:"type"`
}

type AdminDashboardAnalytics struct {
	UsersCount      int `json:"usersCount"`
	ModeratorsCount int `json:"moderatorsCount"`
	ReportsCount    int `json:"reportsCount"`
}

type ReportedThreadResponse struct {
	ID                      uint      `json:"ID"`
	ThreadTitle             string    `json:"threadTitle"`
	ThreadBody              string    `json:"threadBody"`
	ThreadImageURL          string    `json:"threadImageURL"`
	ThreadCreatedAt         time.Time `json:"threadCreatedAt"`
	LikesCount              int       `json:"likesCount"`
	ReportedUsername        string    `json:"reportedUsername"`
	ReportedProfileImageURL string    `json:"reportedProfileImageURL"`
	ReportedUserProfession  string    `json:"reportedUserProfession"`
	ReportCategory          string    `json:"reportCategory"`
	ReportCreatedAt         time.Time `json:"reportCreatedAt"`
	Username                string    `json:"username"`
	ProfileImageURL         string    `json:"profileImageURL"`
}

type ReportedCommentResponse struct {
	ID                      uint      `json:"ID"`
	CommentBody             string    `json:"commentBody"`
	LikesCount              int       `json:"likesCount"`
	CommentCreatedAt        time.Time `json:"commentCreatedAt"`
	ReportedUsername        string    `json:"reportedUsername"`
	ReportedProfileImageURL string    `json:"reportedProfileImageURL"`
	ReportCategory          string    `json:"reportCategory"`
	ReportCreatedAt         time.Time `json:"reportCreatedAt"`
	Username                string    `json:"username"`
	ProfileImageURL         string    `json:"profileImageURL"`
}

type ReportedCommunityResponse struct {
	ID                          uint      `json:"ID"`
	CommunityName               string    `json:"communityName"`
	CommunityImageURL           string    `json:"communityImageURL"`
	CommunityBackgroundImageURL string    `json:"communityBackgroundImageURL"`
	ReportCategory              string    `json:"reportCategory"`
	ReportCreatedAt             time.Time `json:"reportCreatedAt"`
	Username                    string    `json:"username"`
	ProfileImageURL             string    `json:"profileImageURL"`
}

type ReportedUserResponse struct {
	ID                          uint   `json:"ID"`
	ReportedUserUsername        string `json:"reportedUserUsername"`
	ReportedUserName            string `json:"reportedUserName"`
	ReportedUserProfession      string `json:"reportedUserProfession"`
	ReporteduserBio             string `json:"reportedUserBio"`
	ReportedUserProfileImageURL string `json:"reportedUserProfileImageURL"`
	ReportedUserBackgroundURL   string `json:"reportedUserBackgroundURL"`
	ReportingUserUsername       string `json:"reportingUserUsername"`
	ReportingUserName           string `json:"reportinguserName"`
	FollowersCount              int    `json:"followersCount"`
	FollowingCount              int    `json:"followingCount"`
}
