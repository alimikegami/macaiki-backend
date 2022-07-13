package http

import (
	"macaiki/internal/user"
	"macaiki/internal/user/dto"
	"macaiki/pkg/response"
	"macaiki/pkg/utils"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_middL "macaiki/pkg/middleware"
)

type UserHandler struct {
	UserUsecase user.UserUsecase
	JWTSecret   string
}

func NewUserHandler(e *echo.Echo, us user.UserUsecase, JWTSecret string) {
	handler := &UserHandler{
		UserUsecase: us,
		JWTSecret:   JWTSecret,
	}

	e.POST("/api/v1/login", handler.Login)
	e.POST("/api/v1/register", handler.Register)
	e.GET("/api/v1/users", handler.GetAllUsers, middleware.JWT([]byte(JWTSecret)))
	e.GET("/api/v1/users/:userID", handler.GetUser, middleware.JWT([]byte(JWTSecret)))
	e.DELETE("/api/v1/users/:userID", handler.Delete, middleware.JWT([]byte(JWTSecret)))
	e.DELETE("/api/v1/users", handler.DeleteUserByToken, middleware.JWT([]byte(JWTSecret)))

	e.GET("/api/v1/admin/reports", handler.GetReports, middleware.JWT([]byte(JWTSecret)))
	e.GET("/api/v1/admin/analytics", handler.GetDashboardAnalytics, middleware.JWT([]byte(JWTSecret)))
	e.GET("/api/v1/admin/reports/threads/:threadReportID", handler.GetReportedThread, middleware.JWT([]byte(JWTSecret)))
	e.GET("/api/v1/admin/reports/communities/:communityReportID", handler.GetReportedCommunity, middleware.JWT([]byte(JWTSecret)))
	e.GET("/api/v1/admin/reports/comments/:commentReportID", handler.GetReportedComment, middleware.JWT([]byte(JWTSecret)))
	e.GET("/api/v1/admin/reports/users/:userReportID", handler.GetReportedUser, middleware.JWT([]byte(JWTSecret)))

	e.DELETE("/api/v1/admin/ban/users/:userReportID", handler.BanUser, middleware.JWT([]byte(JWTSecret)))
	e.DELETE("/api/v1/admin/ban/comments/:commentReportID", handler.BanComment, middleware.JWT([]byte(JWTSecret)))
	e.DELETE("/api/v1/admin/ban/communities/:communityReportID", handler.BanCommunity, middleware.JWT([]byte(JWTSecret)))
	e.DELETE("/api/v1/admin/ban/threads/:threadReportID", handler.BanThread, middleware.JWT([]byte(JWTSecret)))

	e.PUT("/api/v1/curent-user/email", handler.ChangeEmail, middleware.JWT([]byte(JWTSecret)))
	e.PUT("/api/v1/curent-user/password", handler.ChangePassword, middleware.JWT([]byte(JWTSecret)))
	e.GET("/api/v1/curent-user/threads", handler.GetThreadByToken, middleware.JWT([]byte(JWTSecret)))
	e.GET("/api/v1/curent-user/profile", handler.GetUserByToken, middleware.JWT([]byte(JWTSecret)))
	e.PUT("/api/v1/curent-user/profile", handler.Update, middleware.JWT([]byte(JWTSecret)))
	e.PUT("/api/v1/curent-user/profile-images", handler.SetProfileImage, middleware.JWT([]byte(JWTSecret)))
	e.PUT("/api/v1/curent-user/background-images", handler.SetBackgroundImage, middleware.JWT([]byte(JWTSecret)))

	e.POST("/api/v1/curent-user/user-followers/:userID", handler.Follow, middleware.JWT([]byte(JWTSecret)))
	e.DELETE("/api/v1/curent-user/user-followers/:userID", handler.Unfollow, middleware.JWT([]byte(JWTSecret)))
	e.POST("/api/v1/users/:userID/report", handler.ReportUser, middleware.JWT([]byte(JWTSecret)))

	e.GET("/api/v1/users/:userID/followers", handler.GetUserFollowers, middleware.JWT([]byte(JWTSecret)))
	e.GET("/api/v1/users/:userID/following", handler.GetUserFollowing, middleware.JWT([]byte(JWTSecret)))
	e.GET("/api/v1/users/:userID/threads", handler.GetThreadByUserID, middleware.JWT([]byte(JWTSecret)))

	e.POST("api/v1/curent-user/email-verification", handler.SendOTP)
	e.GET("api/v1/curent-user/email-verification", handler.VerifyOTP)
}

func (u *UserHandler) Login(c echo.Context) error {
	loginInfo := dto.UserLoginRequest{}

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
	search := c.QueryParam("search")

	userID, _ := _middL.ExtractTokenUser(c)
	res, err := u.UserUsecase.GetAll(uint(userID), search)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, res)
}

func (u *UserHandler) GetUser(c echo.Context) error {
	num := c.Param("userID")
	userID, err := strconv.Atoi(num)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	tokenUserID, _ := _middL.ExtractTokenUser(c)

	res, err := u.UserUsecase.Get(uint(userID), uint(tokenUserID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, res)
}

func (u *UserHandler) GetUserByToken(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)

	res, err := u.UserUsecase.Get(uint(userID), uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, res)
}

func (u *UserHandler) Update(c echo.Context) error {
	user := dto.UserUpdateRequest{}
	c.Bind(&user)

	userID, _ := _middL.ExtractTokenUser(c)

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
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}

	curentUserID, curentUserRole := _middL.ExtractTokenUser(c)

	err = u.UserUsecase.Delete(uint(userID), uint(curentUserID), curentUserRole)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (u *UserHandler) DeleteUserByToken(c echo.Context) error {
	curentUserID, curentUserRole := _middL.ExtractTokenUser(c)

	err := u.UserUsecase.Delete(uint(curentUserID), uint(curentUserID), curentUserRole)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (u *UserHandler) ChangeEmail(c echo.Context) error {
	info := dto.UserLoginRequest{}
	userID, _ := _middL.ExtractTokenUser(c)

	c.Bind(&info)

	res, err := u.UserUsecase.ChangeEmail(uint(userID), info)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, res)
}

func (u *UserHandler) ChangePassword(c echo.Context) error {
	newPasswordInfo := dto.UserChangePasswordRequest{}
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
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}

	tokenUserID, _ := _middL.ExtractTokenUser(c)
	followers, err := u.UserUsecase.GetUserFollowers(uint(tokenUserID), uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, followers)
}

func (u *UserHandler) GetUserFollowing(c echo.Context) error {
	num := c.Param("userID")
	userID, err := strconv.Atoi(num)
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}

	tokenUserID, _ := _middL.ExtractTokenUser(c)
	followers, err := u.UserUsecase.GetUserFollowing(uint(tokenUserID), uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, followers)
}

func (u *UserHandler) Follow(c echo.Context) error {
	num := c.Param("userID")
	userID, err := strconv.Atoi(num)
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
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
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}

	follower_id, _ := _middL.ExtractTokenUser(c)
	err = u.UserUsecase.Unfollow(uint(userID), uint(follower_id))
	if err != nil {
		return response.ErrorResponse(c, err)
	}
	return response.SuccessResponse(c, nil)
}

func (u *UserHandler) ReportUser(c echo.Context) error {
	reportedUserID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}
	reportInfo := dto.UserReportRequest{}

	c.Bind(&reportInfo)

	userID, _ := _middL.ExtractTokenUser(c)
	err = u.UserUsecase.Report(uint(userID), uint(reportedUserID), reportInfo.ReportCategoryID)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (u *UserHandler) GetThreadByUserID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}

	tokenUserID, _ := _middL.ExtractTokenUser(c)
	threadResp, err := u.UserUsecase.GetThreadByToken(uint(userID), uint(tokenUserID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, threadResp)
}

func (u *UserHandler) GetThreadByToken(c echo.Context) error {
	userID, _ := _middL.ExtractTokenUser(c)
	threadResp, err := u.UserUsecase.GetThreadByToken(uint(userID), uint(userID))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, threadResp)
}

func (u *UserHandler) SendOTP(c echo.Context) error {
	email := dto.SendOTPRequest{}

	c.Bind(&email)

	err := u.UserUsecase.SendOTP(email)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (u *UserHandler) GetReports(c echo.Context) error {
	_, role := _middL.ExtractTokenUser(c)
	reports, err := u.UserUsecase.GetReports(role)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, reports)
}

func (u *UserHandler) VerifyOTP(c echo.Context) error {
	email := c.QueryParam("email")
	otp := c.QueryParam("otp")

	err := u.UserUsecase.VerifyOTP(email, otp)

	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (u *UserHandler) GetReportedThread(c echo.Context) error {
	threadReportID, err := strconv.Atoi(c.Param("threadReportID"))
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}
	_, role := _middL.ExtractTokenUser(c)
	thread, err := u.UserUsecase.GetReportedThread(role, uint(threadReportID))

	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, thread)
}

func (u *UserHandler) GetReportedCommunity(c echo.Context) error {
	communityReportID, err := strconv.Atoi(c.Param("communityReportID"))
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}
	_, role := _middL.ExtractTokenUser(c)
	community, err := u.UserUsecase.GetReportedCommunity(role, uint(communityReportID))

	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, community)
}

func (u *UserHandler) GetReportedComment(c echo.Context) error {
	commentReportID, err := strconv.Atoi(c.Param("commentReportID"))
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}
	_, role := _middL.ExtractTokenUser(c)
	comment, err := u.UserUsecase.GetReportedComment(role, uint(commentReportID))

	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, comment)
}

func (u *UserHandler) GetReportedUser(c echo.Context) error {
	userReportID, err := strconv.Atoi(c.Param("userReportID"))
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}
	_, role := _middL.ExtractTokenUser(c)
	user, err := u.UserUsecase.GetReportedUser(role, uint(userReportID))

	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, user)
}

func (u *UserHandler) GetDashboardAnalytics(c echo.Context) error {
	_, role := _middL.ExtractTokenUser(c)
	analytics, err := u.UserUsecase.GetDashboardAnalytics(role)
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, analytics)
}

func (u *UserHandler) BanUser(c echo.Context) error {
	userReportID, err := strconv.Atoi(c.Param("userReportID"))
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}
	_, role := _middL.ExtractTokenUser(c)
	err = u.UserUsecase.BanUser(role, uint(userReportID))

	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (u *UserHandler) BanThread(c echo.Context) error {
	threadReportID, err := strconv.Atoi(c.Param("threadReportID"))
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}
	_, role := _middL.ExtractTokenUser(c)
	err = u.UserUsecase.BanThread(role, uint(threadReportID))

	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (u *UserHandler) BanCommunity(c echo.Context) error {
	communityReportID, err := strconv.Atoi(c.Param("communityReportID"))
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}
	_, role := _middL.ExtractTokenUser(c)
	err = u.UserUsecase.BanCommunity(role, uint(communityReportID))

	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}

func (u *UserHandler) BanComment(c echo.Context) error {
	commentReportID, err := strconv.Atoi(c.Param("commentReportID"))
	if err != nil {
		return response.ErrorResponse(c, utils.ErrBadParamInput)
	}
	_, role := _middL.ExtractTokenUser(c)
	err = u.UserUsecase.BanComment(role, uint(commentReportID))

	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, nil)
}
