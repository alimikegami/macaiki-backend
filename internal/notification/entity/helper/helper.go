package helper

import (
	entity "macaiki/internal/notification/entity"
)

func ToNotificationEntity(userID, notificationRefID uint, notificationType, body string) entity.Notification {
	title := ""
	if title == "Follow You" {
		title += " started following you"
	} else if title == "Like Thread" {
		title += " like your thread"
	} else if title == "Comment Thread" {
		title += " comment on your thread"
	}
	notifEntity := entity.Notification{
		UserID:            userID,
		NotificationType:  notificationType,
		NotificationRefID: notificationRefID,
		Title:             title,
		Body:              body,
		IsReaded:          0,
	}

	return notifEntity
}
