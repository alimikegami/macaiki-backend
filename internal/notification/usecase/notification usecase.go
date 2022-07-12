package usecase

import (
	"fmt"
	notification "macaiki/internal/notification"
	dtoNotif "macaiki/internal/notification/dto"
	thread "macaiki/internal/thread"
	dtoThread "macaiki/internal/thread/dto"
	user "macaiki/internal/user"
	dtoUser "macaiki/internal/user/dto"
	"macaiki/pkg/utils"
)

type NotificationUsecaseImpl struct {
	notifRepo  notification.NotificationRepository
	userRepo   user.UserRepository
	threadRepo thread.ThreadRepository
}

func NewNotificationUsecase(notifRepo notification.NotificationRepository, userRepo user.UserRepository, threadRepo thread.ThreadRepository) notification.NotificationUsecase {
	return &NotificationUsecaseImpl{
		notifRepo:  notifRepo,
		userRepo:   userRepo,
		threadRepo: threadRepo,
	}
}

func (nu *NotificationUsecaseImpl) GetAllNotifications(userID uint) ([]dtoNotif.NotificationResponse, error) {
	notifs, err := nu.notifRepo.GetAllNotifications(userID)
	if err != nil {
		return []dtoNotif.NotificationResponse{}, utils.ErrInternalServerError
	}
	user, _ := nu.userRepo.Get(userID)

	notifResp := []dtoNotif.NotificationResponse{}

	for _, val := range notifs {
		title := user.Username
		body := ""
		if val.NotificationType == "Follow You" {
			title += " started following you"
		} else if val.NotificationType == "Upvote Thread" {
			title += " upvote your thread"
		} else if val.NotificationType == "Comment Thread" {
			// TODO: get comment from thread
			title += " comment on your thread"
		}
		notifResp = append(notifResp, dtoNotif.NotificationResponse{
			ID:                 val.ID,
			UserID:             user.ID,
			UserImageUrl:       user.ProfileImageUrl,
			NotificationTypeID: val.NotificationRefID,
			NotificationType:   val.NotificationType,
			Title:              title,
			Body:               body,
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

func (nu *NotificationUsecaseImpl) GetNotificatoinDetail(userID, notificationID uint) (interface{}, error) {
	notif, err := nu.notifRepo.GetNotification(notificationID)
	if err != nil {
		return nil, utils.ErrInternalServerError
	}

	if userID != notif.UserID {
		return nil, utils.ErrUnauthorizedAccess
	}

	err = nu.notifRepo.ReadNotification(notificationID)
	if err != nil {
		fmt.Println("masuk")
		return nil, utils.ErrInternalServerError
	}

	if notif.NotificationType == "Follow You" {
		user, err := nu.userRepo.Get(notif.NotificationRefID)
		if err != nil {
			return nil, utils.ErrInternalServerError
		}

		if user.ID == 0 {
			return nil, utils.ErrNotFound
		}

		totalFollower, _ := nu.userRepo.GetFollowerNumber(user.ID)
		totalFollowing, _ := nu.userRepo.GetFollowingNumber(user.ID)
		totalPost, _ := nu.userRepo.GetThreadsNumber(user.ID)
		return dtoUser.UserDetailResponse{
			ID:                 user.ID,
			Username:           user.Username,
			Name:               user.Name,
			ProfileImageUrl:    user.ProfileImageUrl,
			BackgroundImageUrl: user.BackgroundImageUrl,
			Bio:                user.Bio,
			Profession:         user.Profession,
			TotalFollower:      totalFollower,
			TotalFollowing:     totalFollowing,
			TotalPost:          totalPost,
			IsFollowed:         user.IsFollowed,
			IsMine:             user.IsMine,
		}, nil
	} else if notif.NotificationType == "Upvote Thread" {
		thread, err := nu.threadRepo.GetThreadByID(notif.NotificationRefID)
		if err != nil {
			return nil, err
		}
		return dtoThread.ThreadResponse{
			ID:          thread.ID,
			Title:       thread.Title,
			Body:        thread.Body,
			CommunityID: thread.CommunityID,
			ImageURL:    thread.ImageURL,
			UserID:      thread.UserID,
			CreatedAt:   thread.CreatedAt,
			UpdatedAt:   thread.UpdatedAt,
		}, nil
	} else if notif.NotificationType == "Comment Thread" {
		comment, err := nu.threadRepo.GetCommentByID(notif.NotificationRefID)
		if err != nil {
			return nil, err
		}

		user, _ := nu.userRepo.Get(comment.UserID)
		return dtoThread.CommentResponse{
			ID:                    comment.CommentID,
			Body:                  comment.Body,
			UserID:                user.ID,
			Username:              user.Username,
			UserProfilePictureURL: user.ProfileImageUrl,
			ThreadID:              comment.ThreadID,
			CreatedAt:             comment.CreatedAt,
			LikesCount:            0,
		}, nil
	}

	return nil, utils.ErrNotFound
}
