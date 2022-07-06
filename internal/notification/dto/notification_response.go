package dto

type NotificationResponse struct {
	ID                 uint   `json:"ID"`
	UserID             uint   `json:"UserID"`
	NotificationTypeID uint   `json:"notificationTypeID"`
	NotificationType   string `json:"notificationType"`
	IsReaded           int    `json:"isReaded"`
	Title              string `json:"title"`
	Body               string `json:"body"`
}
