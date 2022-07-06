package notification

import (
	"macaiki/internal/notification/dto"
)

type NotificationUsecase interface {
	CreateNotification(userID, notificationRefID uint, notificationType, body string) error
	GetAllNotifications(userID uint) ([]dto.NotificationResponse, error)
	ReadAllNotifications(userID uint) ([]dto.NotificationResponse, error)
	DeleteAllNotifications(userID uint) ([]dto.NotificationResponse, error)
	// GetNotificatoinDetail(notificationID uint) (interface{}, error)
}
