package http

import (
	"macaiki/domain"
	"macaiki/user/delivery/http/response"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
	JWTSecret   string
}

func NewUserHandler(e *echo.Echo, us domain.UserUsecase, JWTSecret string) {
	handler := &UserHandler{
		UserUsecase: us,
		JWTSecret:   JWTSecret,
	}

	e.POST("/api/v1/login", handler.Login)
	e.POST("/api/v1/register", handler.Register)
	e.GET("/api/v1/users", handler.GetAllUsers, middleware.JWT([]byte(JWTSecret)))
	e.GET("/api/v1/users/:user_id", handler.GetUser)
	e.PUT("api/v1/users/:user_id", handler.Update)
	e.DELETE("api/v1/users/:user_id", handler.Delete)
}

func (u *UserHandler) Login(c echo.Context) error {
	loginInfo := domain.User{}

	c.Bind(&loginInfo)

	token, err := u.UserUsecase.Login(loginInfo.Email, loginInfo.Password)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, response.ToTokenResponse(token))
}

func (u *UserHandler) Register(c echo.Context) error {
	user := domain.User{}
	c.Bind(&user)

	user, err := u.UserUsecase.Register(user)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, response.ToUserResponse(user))
}

func (u *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := u.UserUsecase.GetAll()
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, response.ToListUserResponse(users))
}

func (u *UserHandler) GetUser(c echo.Context) error {
	num := c.Param("user_id")
	user_id, err := strconv.Atoi(num)
	if err != nil {
		response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	user, err := u.UserUsecase.Get(uint(user_id))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, response.ToUserResponse(user))
}

func (u *UserHandler) Update(c echo.Context) error {
	num := c.Param("user_id")
	user_id, err := strconv.Atoi(num)
	if err != nil {
		response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	user := domain.User{}
	c.Bind(&user)

	user, err = u.UserUsecase.Update(user, uint(user_id))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, response.ToUserResponse(user))
}

func (u *UserHandler) Delete(c echo.Context) error {
	num := c.Param("user_id")
	user_id, err := strconv.Atoi(num)
	if err != nil {
		response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	user, err := u.UserUsecase.Delete(uint(user_id))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, response.ToUserResponse(user))
}
