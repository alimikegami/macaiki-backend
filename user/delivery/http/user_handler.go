package http

import (
	"macaiki/domain"
	"macaiki/user/delivery/http/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(e *echo.Echo, us domain.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}
	e.GET("/api/v1/users", handler.FetchUser)
}

func (u *UserHandler) FetchUser(c echo.Context) error {
	users, err := u.UserUsecase.GetAll()
	if err != nil {
		return c.JSON(response.ErrorResponse(err, getStatusCode(err)))
	}

	return c.JSON(response.SuccessResponse(users))
}

func getStatusCode(err error) int {
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusOK
	}
}
