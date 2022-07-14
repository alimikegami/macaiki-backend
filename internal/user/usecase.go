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

	GetThreadByToken(userID, tokenUserID uint) ([]dtoThread.DetailedThreadResponse, error)
	SendOTP(email dto.SendOTPRequest) error
	VerifyOTP(email, OTPCode string) error

	BanUser(userRole string, userReportID uint) error
	BanThread(userRole string, threadReportID uint) error
	BanComment(userRole string, commentReportID uint) error
	BanCommunity(userRole string, communityReportID uint) error

	GetReports(curentUserRole string) ([]dto.BriefReportResponse, error)
	GetDashboardAnalytics(userRole string) (dto.AdminDashboardAnalytics, error)
	GetReportedThread(userRole string, threadReportID uint) (dto.ReportedThreadResponse, error)
	GetReportedCommunity(userRole string, communityReportID uint) (dto.ReportedCommunityResponse, error)
	GetReportedComment(userRole string, commentReportID uint) (dto.ReportedCommentResponse, error)
	GetReportedUser(userRole string, userReportID uint) (dto.ReportedUserResponse, error)

	DeleteThreadReport(userRole string, threadReportID uint) error
	DeleteUserReport(userRole string, userReportID uint) error
	DeleteCommentReport(userRole string, commentReportID uint) error
	DeleteCommunityReport(userRole string, communityReportID uint) error
}
