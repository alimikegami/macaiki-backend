package dto

import "time"

type CommentResponse struct {
	ID                    uint      `json:"id"`
	Body                  string    `json:"body"`
	UserID                uint      `json:"userID"`
	Username              string    `json:"username"`
	UserProfilePictureURL string    `json:"userProfilePictureURL"`
	ThreadID              uint      `json:"threadID"`
	CreatedAt             time.Time `json:"createdAt"`
	LikesCount            int       `json:"likesCount"`
}
