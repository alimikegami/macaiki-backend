package notification

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	UserID            uint
	NotificationType  string
	NotificationRefID uint
	IsReaded          int
}
