package usecase

import (
	notification "macaiki/internal/notification"
	dtoNotif "macaiki/internal/notification/dto"
	entityNotif "macaiki/internal/notification/entity"
	user "macaiki/internal/user"
	entityUser "macaiki/internal/user/entity"
	"macaiki/pkg/utils"
)

type NotificationUsecaseImpl struct {
	notifRepo notification.NotificationRepository
	userRepo  user.UserRepository
}

func NewNotificationUsecase(notifRepo notification.NotificationRepository, userRepo user.UserRepository) notification.NotificationUsecase {
	return &NotificationUsecaseImpl{notifRepo: notifRepo, userRepo: userRepo}
}

func (nu *NotificationUsecaseImpl) CreateNotification(notification dtoNotif.NotificationRequest) error {
	notifEntity := entityNotif.Notification{
		UserID:             notification.UserID,
		NotificationType:   notification.NotificationType,
		NotificationTypeID: notification.NotificationTypeID,
		IsReaded:           notification.IsReaded,
	}
	err := nu.notifRepo.StoreNotification(notifEntity)
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
}

func (nu *NotificationUsecaseImpl) GetAllNotifications(userID uint) ([]dtoNotif.NotificationResponse, error) {
	notifs, err := nu.notifRepo.GetAllNotifications(userID)
	if err != nil {
		return []dtoNotif.NotificationResponse{}, utils.ErrInternalServerError
	}
	user, _ := nu.userRepo.Get(userID)

	return ToNotificationResponse(user, notifs), err
}

func (nu *NotificationUsecaseImpl) ReadAllNotifications(userID uint) ([]dtoNotif.NotificationResponse, error) {
	err := nu.notifRepo.ReadAllNotifications(userID)
	if err != nil {
		return []dtoNotif.NotificationResponse{}, utils.ErrInternalServerError
	}

	return nu.GetAllNotifications(userID)
}

func (nu *NotificationUsecaseImpl) DeleteAllNotifications(userID uint) ([]dtoNotif.NotificationResponse, error) {
	err := nu.notifRepo.DeleleteAllNotifications(userID)
	if err != nil {
		return []dtoNotif.NotificationResponse{}, utils.ErrInternalServerError
	}

	return nu.GetAllNotifications(userID)
}

func ToNotificationResponse(user entityUser.User, notifs []entityNotif.Notification) []dtoNotif.NotificationResponse {
	notifResp := []dtoNotif.NotificationResponse{}
	title := user.Username
	for _, val := range notifs {
		if val.NotificationType == "Follow You" {
			title += " started following you"
		} else if val.NotificationType == "Like Thread" {
			title += " like your thread"
		} else if val.NotificationType == "Comment Thread" {
			title += " comment on your thread"
		}
		notifResp = append(notifResp, dtoNotif.NotificationResponse{
			ID:                 val.ID,
			UserID:             user.ID,
			UserImageUrl:       user.ProfileImageUrl,
			NotificationTypeID: val.NotificationTypeID,
			NotificationType:   val.NotificationType,
			Title:              title,
			IsReaded:           val.IsReaded,
			CreatedAt:          val.CreatedAt,
		})
	}

	return notifResp
}
