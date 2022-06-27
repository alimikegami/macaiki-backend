package dto

type CommentRequest struct {
	Body      string `json:"body"`
	UserID    uint   `json:"userID"`
	ThreadID  uint   `json:"threadID"`
	CommentID uint   `json:"commentID"`
}
