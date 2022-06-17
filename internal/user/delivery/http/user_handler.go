package http

import (
	"macaiki/internal/domain"
	"macaiki/internal/user/dto"
	"macaiki/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_middL "macaiki/pkg/middleware"
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
	e.GET("/api/v1/users/:userID", handler.GetUser)
	e.PUT("api/v1/users/:userID", handler.Update)
	e.DELETE("api/v1/users/:userID", handler.Delete)

	e.GET("/api/v1/my-profile", handler.GetUserByToken, middleware.JWT([]byte(JWTSecret)))

	e.PUT("api/v1/users/:userID/profile-images", handler.SetProfileImage)
	e.PUT("api/v1/users/:userID/background-images", handler.SetBackgroundImage)
	e.GET("api/v1/users/:userID/followers", handler.GetUserFollowers)
	e.GET("api/v1/users/:userID/following", handler.GetUserFollowing)
	e.GET("api/v1/users/:userID/follow", handler.Follow, middleware.JWT([]byte(JWTSecret)))
	e.GET("api/v1/users/:userID/unfollow", handler.Unfollow, middleware.JWT([]byte(JWTSecret)))
}

func (u *UserHandler) Login(c echo.Context) error {
	loginInfo := dto.LoginUserRequest{}

	c.Bind(&loginInfo)

	token, err := u.UserUsecase.Login(loginInfo.Email, loginInfo.Password)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, token)
}

func (u *UserHandler) Register(c echo.Context) error {
	user := dto.UserRequest{}
	c.Bind(&user)

	err := u.UserUsecase.Register(user)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (u *UserHandler) GetAllUsers(c echo.Context) error {
	res, err := u.UserUsecase.GetAll()
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, res)
}

func (u *UserHandler) GetUser(c echo.Context) error {
	num := c.Param("userID")
	userID, err := strconv.Atoi(num)
	if err != nil {
		response.ErrorResponse(c, err)
	}

	res, err := u.UserUsecase.Get(uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, res)
}

func (u *UserHandler) GetUserByToken(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

	res, err := u.UserUsecase.Get(uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, res)
}

func (u *UserHandler) Update(c echo.Context) error {
	num := c.Param("userID")
	userID, err := strconv.Atoi(num)
	if err != nil {
		response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	user := dto.UpdateUserRequest{}
	c.Bind(&user)

	res, err := u.UserUsecase.Update(user, uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, res)
}

func (u *UserHandler) Delete(c echo.Context) error {
	num := c.Param("userID")
	userID, err := strconv.Atoi(num)
	if err != nil {
		response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	err = u.UserUsecase.Delete(uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (u *UserHandler) SetProfileImage(c echo.Context) error {
	num := c.Param("userID")
	userID, err := strconv.Atoi(num)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	img, err := c.FormFile("profileImage")
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	url, err := u.UserUsecase.SetProfileImage(uint(userID), img)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, map[string]string{
		"profileImageUrl": url,
	})
}

func (u *UserHandler) SetBackgroundImage(c echo.Context) error {
	num := c.Param("userID")
	userID, err := strconv.Atoi(num)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	img, err := c.FormFile("backgroundImage")
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	url, err := u.UserUsecase.SetBackgroundImage(uint(userID), img)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, map[string]string{
		"backgroundImageUrl": url,
	})
}

func (u *UserHandler) GetUserFollowers(c echo.Context) error {
	num := c.Param("userID")
	userID, err := strconv.Atoi(num)
	if err != nil {
		return response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	followers, err := u.UserUsecase.GetUserFollowers(uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, followers)
}

func (u *UserHandler) GetUserFollowing(c echo.Context) error {
	num := c.Param("userID")
	userID, err := strconv.Atoi(num)
	if err != nil {
		return response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	followers, err := u.UserUsecase.GetUserFollowing(uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, followers)
}

func (u *UserHandler) Follow(c echo.Context) error {
	num := c.Param("userID")
	userID, err := strconv.Atoi(num)
	if err != nil {
		response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	follower_id, _ := _middL.ExtractTokenUser(c)
	err = u.UserUsecase.Follow(uint(userID), uint(follower_id))
	if err != nil {
		return response.ErrorResponse(c, err)
	}
	return response.SuccessResponse(c, nil)
}

func (u *UserHandler) Unfollow(c echo.Context) error {
	num := c.Param("userID")
	userID, err := strconv.Atoi(num)
	if err != nil {
		response.ErrorResponse(c, domain.ErrBadParamInput)
	}

	follower_id, _ := _middL.ExtractTokenUser(c)
	err = u.UserUsecase.Unfollow(uint(userID), uint(follower_id))
	if err != nil {
		return response.ErrorResponse(c, err)
	}
	return response.SuccessResponse(c, nil)
}
