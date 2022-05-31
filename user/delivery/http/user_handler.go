package http

import (
	"macaiki/domain"
	"macaiki/user/delivery/http/response"
	"net/http"
	"strconv"

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
	e.POST("/api/v1/login", handler.Login)
	e.POST("/api/v1/register", handler.Register)
	e.GET("/api/v1/users", handler.GetAllUsers)
	e.GET("/api/v1/users/:user_id", handler.GetUser)
	e.PUT("api/v1/users/:user_id", handler.Update)
	e.DELETE("api/v1/users/:user_id", handler.Delete)
}

func (u *UserHandler) Login(c echo.Context) error {
	loginInfo := domain.User{}

	c.Bind(&loginInfo)

	token, err := u.UserUsecase.Login(loginInfo.Email, loginInfo.Password)
	if err != nil {
		return c.JSON(response.ErrorResponse(err, getStatusCode(err)))
	}

	return c.JSON(response.SuccessResponse(token))
}

func (u *UserHandler) Register(c echo.Context) error {
	user := domain.User{}
	c.Bind(&user)

	user, err := u.UserUsecase.Register(user)
	if err != nil {
		return c.JSON(response.ErrorResponse(err, getStatusCode(err)))
	}

	return c.JSON(response.SuccessResponse(response.ToUserResponse(user)))
}

func (u *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := u.UserUsecase.GetAll()
	if err != nil {
		return c.JSON(response.ErrorResponse(err, getStatusCode(err)))
	}

	return c.JSON(response.SuccessResponse(response.ToListUserResponse(users)))
}

func (u *UserHandler) GetUser(c echo.Context) error {
	num := c.Param("user_id")
	user_id, _ := strconv.Atoi(num)

	user, err := u.UserUsecase.Get(uint(user_id))
	if err != nil {
		return c.JSON(response.ErrorResponse(err, getStatusCode(err)))
	}

	return c.JSON(response.SuccessResponse(response.ToUserResponse(user)))
}

func (u *UserHandler) Update(c echo.Context) error {
	num := c.Param("user_id")
	user_id, _ := strconv.Atoi(num)

	user := domain.User{}
	c.Bind(&user)

	user, err := u.UserUsecase.Update(user, uint(user_id))
	if err != nil {
		return c.JSON(response.ErrorResponse(err, getStatusCode(err)))
	}

	return c.JSON(response.SuccessResponse(response.ToUserResponse(user)))
}

func (u *UserHandler) Delete(c echo.Context) error {
	num := c.Param("user_id")
	user_id, _ := strconv.Atoi(num)

	user, err := u.UserUsecase.Delete(uint(user_id))
	if err != nil {
		return c.JSON(response.ErrorResponse(err, getStatusCode(err)))
	}

	return c.JSON(response.SuccessResponse(response.ToUserResponse(user)))
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
