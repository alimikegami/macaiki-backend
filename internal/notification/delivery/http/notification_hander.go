package http

import (
	notification "macaiki/internal/notification"
	_middL "macaiki/pkg/middleware"
	"macaiki/pkg/response"
	"macaiki/pkg/utils"
	"strconv"

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
	e.PUT("api/v1/notifications", notifHandler.ReadAllNotifications, middleware.JWT([]byte(JWTSecret)))
	e.DELETE("api/v1/notifications", notifHandler.DeleteAllNotifications, middleware.JWT([]byte(JWTSecret)))
	e.GET("api/v1/notifications/:notificationID", notifHandler.ReadNotification, middleware.JWT([]byte(JWTSecret)))
}

func (notifHandler *NotificationHandler) GetAllNotifications(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)
	notifications, err := notifHandler.notifUsecase.GetAllNotifications(uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, notifications)
}

func (notifHandler *NotificationHandler) ReadAllNotifications(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)
	notifResp, err := notifHandler.notifUsecase.ReadAllNotifications(uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, notifResp)
}

func (notifHandler *NotificationHandler) DeleteAllNotifications(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)
	notifResp, err := notifHandler.notifUsecase.DeleteAllNotifications(uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, notifResp)
}

func (notifHandler *NotificationHandler) ReadNotification(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)
	notifID, err := strconv.Atoi(c.Param("notificationID"))
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}
	notifResp, err := notifHandler.notifUsecase.GetNotificatoinDetail(uint(userID), uint(notifID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, notifResp)
}
