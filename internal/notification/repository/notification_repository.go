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
func (nr *NotificationRepositoryImpl) GetAllNotifications(userID uint) ([]entity.Notification, error) {
	notifications := []entity.Notification{}
	res := nr.db.Where("user_id = ?", userID).Order("created_at desc").Find(&notifications)
	err := res.Error
	if err != nil {
		return []entity.Notification{}, err
	}
	return notifications, nil
}

func (nr *NotificationRepositoryImpl) ReadAllNotifications(userID uint) error {
	res := nr.db.Model(&entity.Notification{}).Where("user_id = ?", userID).Update("is_readed", 1)
	err := res.Error
	if err != nil {
		return err
	}
	return nil
}

func (nr *NotificationRepositoryImpl) DeleleteAllNotifications(userID uint) error {
	res := nr.db.Where("user_id = ?", userID).Delete(&entity.Notification{})
	err := res.Error
	if err != nil {
		return err
	}
	return nil
}
