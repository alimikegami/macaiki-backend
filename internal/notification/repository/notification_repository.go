package repository

import (
	notification "macaiki/internal/notification"
	entity "macaiki/internal/notification/entity"

	"gorm.io/gorm"
)

type NotificationRepositoryImpl struct {
	db *gorm.DB
}

func NewNotificaionRepository(db *gorm.DB) notification.NotificationRepository {
	return &NotificationRepositoryImpl{db: db}
}

func (nr *NotificationRepositoryImpl) StoreNotification(notification entity.Notification) error {
	res := nr.db.Create(&notification)
	err := res.Error
	if err != nil {
		return err
	}
	return nil
}
func (nr *NotificationRepositoryImpl) GetAllNotification(userID uint) ([]entity.Notification, error) {
	notifications := []entity.Notification{}
	res := nr.db.Where("user_id = ?", userID).Find(&notifications)
	err := res.Error
	if err != nil {
		return []entity.Notification{}, err
	}
	return notifications, nil
}
