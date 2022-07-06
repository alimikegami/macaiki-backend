package dto

type NotificationRequest struct {
	UserID             uint
	NotificationType   string
	NotificationTypeID uint
	Title              string
	Body               string
	IsReaded           int
}
