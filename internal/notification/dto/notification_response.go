package dto

import "time"

type NotificationResponse struct {
	ID                 uint      `json:"ID"`
	UserID             uint      `json:"userID"`
	UserImageUrl       string    `json:"userImageUrl"`
	NotificationTypeID uint      `json:"notificationTypeID"`
	NotificationType   string    `json:"notificationType"`
	IsReaded           int       `json:"isReaded"`
	Title              string    `json:"title"`
	Body               string    `json:"body"`
	CreatedAt          time.Time `json:"craetedAt"`
}
