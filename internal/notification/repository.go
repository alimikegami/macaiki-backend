package notification

import entity "macaiki/internal/notification/entity"

type NotificationRepository interface {
	StoreNotification(notification entity.Notification) error
	GetAllNotification(userID uint) ([]entity.Notification, error)
}
