package notification

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	UserID             uint
	NotificationType   string
	NotificationTypeID uint
	IsReaded           int
}
