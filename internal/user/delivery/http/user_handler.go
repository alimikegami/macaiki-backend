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
	e.GET("/api/v1/users", handler.GetAllUsers)
	e.GET("/api/v1/users/:userID", handler.GetUser)
	e.PUT("/api/v1/users/:userID", handler.Update, middleware.JWT([]byte(JWTSecret)))
	e.DELETE("/api/v1/users/:userID", handler.Delete, middleware.JWT([]byte(JWTSecret)))

	e.PUT("/api/v1/curent-user/email", handler.ChangeEmail, middleware.JWT([]byte(JWTSecret)))
	e.PUT("/api/v1/curent-user/password", handler.ChangePassword, middleware.JWT([]byte(JWTSecret)))
	e.GET("/api/v1/curent-user/profile", handler.GetUserByToken, middleware.JWT([]byte(JWTSecret)))
	e.PUT("/api/v1/curent-user/profile-images", handler.SetProfileImage, middleware.JWT([]byte(JWTSecret)))
	e.PUT("/api/v1/curent-user/background-images", handler.SetBackgroundImage, middleware.JWT([]byte(JWTSecret)))

	e.POST("/api/v1/curent-user/user-followers/:userID", handler.Follow, middleware.JWT([]byte(JWTSecret)))
	e.DELETE("/api/v1/curent-user/user-followers/:userID", handler.Unfollow, middleware.JWT([]byte(JWTSecret)))
	e.POST("/api/v1/curent-user/reports", handler.ReportUser, middleware.JWT([]byte(JWTSecret)))

	e.GET("/api/v1/users/:userID/followers", handler.GetUserFollowers)
	e.GET("/api/v1/users/:userID/following", handler.GetUserFollowing)
}

func (u *UserHandler) Login(c echo.Context) error {
	loginInfo := dto.LoginUserRequest{}

	c.Bind(&loginInfo)

	token, err := u.UserUsecase.Login(loginInfo)
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
	key := c.QueryParam("search")
	res, err := u.UserUsecase.GetAll(key)
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

	curentUserID, _ := _middL.ExtractTokenUser(c)

	res, err := u.UserUsecase.Update(user, uint(userID), uint(curentUserID))
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

	curentUserID, curentUserRole := _middL.ExtractTokenUser(c)

	err = u.UserUsecase.Delete(uint(userID), uint(curentUserID), curentUserRole)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (u *UserHandler) ChangeEmail(c echo.Context) error {
	info := dto.LoginUserRequest{}
	userID, _ := _middL.ExtractTokenUser(c)

	c.Bind(&info)

	res, err := u.UserUsecase.ChangeEmail(uint(userID), info)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, res)
}

func (u *UserHandler) ChangePassword(c echo.Context) error {
	newPasswordInfo := dto.ChangePasswordUserRequest{}
	userID, _ := _middL.ExtractTokenUser(c)

	c.Bind(&newPasswordInfo)

	err := u.UserUsecase.ChangePassword(uint(userID), newPasswordInfo)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (u *UserHandler) SetProfileImage(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

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
	userID, _ := _middL.ExtractTokenUser(c)

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

func (u *UserHandler) ReportUser(c echo.Context) error {
	reportInfo := dto.UserReportRequest{}

	c.Bind(&reportInfo)

	userID, _ := _middL.ExtractTokenUser(c)
	err := u.UserUsecase.Report(uint(userID), reportInfo.UserID, reportInfo.ReportCategoryID)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}
