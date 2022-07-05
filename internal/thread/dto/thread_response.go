package dto

import (
	"time"
)

type ThreadResponse struct {
	ID          uint      `json:"ID"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	CommunityID uint      `json:"communityID"`
	ImageURL    string    `json:"imageURL"`
	UserID      uint      `json:"userID"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type DetailedThreadResponse struct {
	ID                    uint      `json:"ID"`
	Title                 string    `json:"title"`
	Body                  string    `json:"body"`
	CommunityID           uint      `json:"communityID"`
	ImageURL              string    `json:"imageURL"`
	UserID                uint      `json:"userID"`
	UserName              string    `json:"userName"`
	UserProfession        string    `json:"userProfession"`
	UserProfilePictureURL string    `json:"userProfilePictureURL"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	UpvotesCount          int       `json:"upvotesCount"`
	IsUpvoted             int       `json:"isUpvoted"`
	IsDownVoted           int       `json:"isDownvoted"`
	IsFollowed            int       `json:"isFollowed"`
}
