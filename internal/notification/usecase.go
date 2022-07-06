package notification

import (
	"macaiki/internal/notification/dto"
)

type NotificationUsecase interface {
	CreateNotification(notification dto.NotificationRequest) error
	GetAllNotification(userID uint) ([]dto.NotificationResponse, error)
}
