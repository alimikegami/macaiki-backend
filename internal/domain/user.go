package domain

import (
	"macaiki/internal/user/dto"
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email              string `gorm:"uniqueIndex;size:75"`
	Username           string `gorm:"uniqueIndex;size:50"`
	Password           string
	Name               string
	ProfileImageUrl    string
	BackgroundImageUrl string
	Bio                string
	Proffesion         string
	Role               string
	IsBanned           bool
	Followers          []User       `gorm:"many2many:user_followers"`
	Report             []UserReport `gorm:"foreignKey:UserID"`
}

type UserReport struct {
	UserID           uint `gorm:"primaryKey"`
	ReportedUserID   uint `gorm:"primaryKey"`
	ReportCategoryID uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type UserUsecase interface {
	Login(email, password string) (dto.LoginResponse, error)
	Register(user dto.UserRequest) error
	GetAll() ([]dto.UserResponse, error)
	Get(id uint) (dto.UserDetailResponse, error)
	Update(userUpdate dto.UpdateUserRequest, id uint) (dto.UserResponse, error)
	Delete(id uint) error

	SetProfileImage(id uint, img *multipart.FileHeader) (string, error)
	SetBackgroundImage(id uint, img *multipart.FileHeader) (string, error)
	GetUserFollowers(id uint) ([]dto.UserResponse, error)
	GetUserFollowing(id uint) ([]dto.UserResponse, error)
	Follow(userID, userFollowerID uint) error
	Unfollow(userID, userFollowerID uint) error

	Report(userID, userReportedID, ReportCategoryID uint) error
}

type UserRepository interface {
	GetAll() ([]User, error)
	Store(user User) error
	Get(id uint) (User, error)
	Update(userDB *User, user User) (User, error)
	Delete(id uint) (User, error)
	GetByEmail(email string) (User, error)

	GetFollowerNumber(id uint) (int, error)
	GetFollowingNumber(id uint) (int, error)
	Follow(user, userFollower User) (User, error)
	Unfollow(user, userFollower User) (User, error)
	GetFollower(user User) ([]User, error)
	GetFollowing(user User) ([]User, error)
	SetUserImage(id uint, imageURL string, tableName string) error

	StoreReport(userReport UserReport) error
}
