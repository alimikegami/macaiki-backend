package usecase

import (
	notification "macaiki/internal/notification"
	dtoNotif "macaiki/internal/notification/dto"
	user "macaiki/internal/user"
	"macaiki/pkg/utils"
)

type NotificationUsecaseImpl struct {
	notifRepo notification.NotificationRepository
	userRepo  user.UserRepository
}

func NewNotificationUsecase(notifRepo notification.NotificationRepository, userRepo user.UserRepository) notification.NotificationUsecase {
	return &NotificationUsecaseImpl{notifRepo: notifRepo, userRepo: userRepo}
}

func (nu *NotificationUsecaseImpl) GetAllNotifications(userID uint) ([]dtoNotif.NotificationResponse, error) {
	notifs, err := nu.notifRepo.GetAllNotifications(userID)
	if err != nil {
		return []dtoNotif.NotificationResponse{}, utils.ErrInternalServerError
	}
	user, _ := nu.userRepo.Get(userID)

	notifResp := []dtoNotif.NotificationResponse{}
	for _, val := range notifs {
		notifResp = append(notifResp, dtoNotif.NotificationResponse{
			ID:                 val.ID,
			UserID:             user.ID,
			UserImageUrl:       user.ProfileImageUrl,
			NotificationTypeID: val.NotificationRefID,
			NotificationType:   val.NotificationType,
			Title:              val.Title,
			Body:               val.Body,
			IsReaded:           val.IsReaded,
			CreatedAt:          val.CreatedAt,
		})
	}

	return notifResp, err
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
