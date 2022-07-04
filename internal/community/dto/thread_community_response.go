package dto

import "time"

type DetailedThreadResponse struct {
	ID                    uint      `json:"ID"`
	Title                 string    `json:"title"`
	Body                  string    `json:"body"`
	CommunityID           uint      `json:"communityID"`
	ImageURL              string    `json:"imageURL"`
	LikesCount            int       `json:"likesCount"`
	IsLiked               int       `json:"isLiked"`
	IsFollowed            int       `json:"isFollowed"`
	UserID                uint      `json:"userID"`
	UserName              string    `json:"userName"`
	UserProfession        string    `json:"userProfession"`
	UserProfilePictureURL string    `json:"userProfilePictureURL"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}
