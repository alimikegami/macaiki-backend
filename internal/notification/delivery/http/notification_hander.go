package http

import (
	notification "macaiki/internal/notification"
	_middL "macaiki/pkg/middleware"
	"macaiki/pkg/response"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type NotificationHandler struct {
	notifUsecase notification.NotificationUsecase
	JWTSecret    string
}

func NewNotificationHandler(e *echo.Echo, notifUsecase notification.NotificationUsecase, JWTSecret string) {
	notifHandler := NotificationHandler{notifUsecase, JWTSecret}
	e.GET("api/v1/notifications", notifHandler.GetAllNotifications, middleware.JWT([]byte(JWTSecret)))
}

func (notifHandler *NotificationHandler) GetAllNotifications(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)
	notifications, err := notifHandler.notifUsecase.GetAllNotification(uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, notifications)
}
