package notification

import entity "macaiki/internal/notification/entity"

type NotificationRepository interface {
	StoreNotification(notification entity.Notification) error
	GetAllNotifications(userID uint) ([]entity.Notification, error)
	ReadAllNotifications(userID uint) error
	DeleleteAllNotifications(userID uint) error
}
