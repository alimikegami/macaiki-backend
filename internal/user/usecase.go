package user

import (
	dtoThread "macaiki/internal/thread/dto"
	"macaiki/internal/user/dto"
	"mime/multipart"
)

type UserUsecase interface {
	Login(loginInfo dto.UserLoginRequest) (dto.LoginResponse, error)
	Register(user dto.UserRequest) error
	GetAll(userID uint, search string) ([]dto.UserResponse, error)
	Get(id, tokenUserID uint) (dto.UserDetailResponse, error)
	Update(userUpdate dto.UserUpdateRequest, id uint) (dto.UserUpdateResponse, error)
	Delete(id uint, curentUserID uint, curentUser string) error

	ChangeEmail(id uint, info dto.UserLoginRequest) (string, error)
	ChangePassword(id uint, passwordInfo dto.UserChangePasswordRequest) error

	SetProfileImage(id uint, img *multipart.FileHeader) (string, error)
	SetBackgroundImage(id uint, img *multipart.FileHeader) (string, error)
	GetUserFollowers(tokenUserID, getFollowingUserID uint) ([]dto.UserResponse, error)
	GetUserFollowing(tokenUserID, getFollowingUserID uint) ([]dto.UserResponse, error)
	Follow(userID, userFollowerID uint) error
	Unfollow(userID, userFollowerID uint) error

	Report(userID, userReportedID, ReportCategoryID uint) error

	GetThreadByToken(tokenUserID uint) ([]dtoThread.ThreadResponse, error)
	SendOTP(email dto.SendOTPRequest) error
	VerifyOTP(email, OTPCode string) error
	GetReports(curentUserRole string) ([]dto.BriefReportResponse, error)
	GetDashboardAnalytics(userRole string) (dto.AdminDashboardAnalytics, error)

}
