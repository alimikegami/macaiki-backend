package user

import (
	"macaiki/internal/user/dto"
	"mime/multipart"
)

type UserUsecase interface {
	Login(loginInfo dto.LoginUserRequest) (dto.LoginResponse, error)
	Register(user dto.UserRequest) error
	// GetAll(username string) ([]dto.UserResponse, error)
	GetAllWithDetail(userID uint, search string) ([]dto.UserResponse, error)
	Get(id, tokenUserID uint) (dto.UserDetailResponse, error)
	Update(userUpdate dto.UpdateUserRequest, id, curentUserID uint) (dto.UserResponse, error)
	Delete(id uint, curentUserID uint, curentUser string) error

	ChangeEmail(id uint, info dto.LoginUserRequest) (dto.UserResponse, error)
	ChangePassword(id uint, passwordInfo dto.ChangePasswordUserRequest) error

	SetProfileImage(id uint, img *multipart.FileHeader) (string, error)
	SetBackgroundImage(id uint, img *multipart.FileHeader) (string, error)
	GetUserFollowers(id uint) ([]dto.UserResponse, error)
	GetUserFollowing(id uint) ([]dto.UserResponse, error)
	Follow(userID, userFollowerID uint) error
	Unfollow(userID, userFollowerID uint) error

	Report(userID, userReportedID, ReportCategoryID uint) error
}
