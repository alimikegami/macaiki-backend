package dto

import (
	"time"
)

type ThreadResponse struct {
	ID        uint      `json:"ID"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	TopicID   uint      `json:"topicID"`
	ImageURL  string    `json:"imageURL"`
	UserID    uint      `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
