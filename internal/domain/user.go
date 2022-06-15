package domain

import (
	"macaiki/internal/user/dto"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email              string `gorm:"uniqueIndex"`
	Username           string `gorm:"uniqueIndex"`
	Password           string
	Name               string
	ProfileImageUrl    string
	BackgroundImageUrl string
	Bio                string
	Proffesion         string
	Role               string
	IsBanned           bool
	Followers          []User `gorm:"many2many:user_followers"`
}

type UserUsecase interface {
	Login(email, password string) (dto.LoginResponse, error)
	Register(user dto.UserRequest) (dto.UserResponse, error)
	GetAll() ([]dto.UserResponse, error)
	Get(id uint) (dto.UserDetailResponse, error)
	Update(userUpdate dto.UpdateUserRequest, id uint) (dto.UserResponse, error)
	Delete(id uint) error

	GetUserFollowers(id uint) ([]dto.UserResponse, error)
	GetUserFollowing(id uint) ([]dto.UserResponse, error)
	Follow(user_id, user_follower_id uint) error
	Unfollow(user_id, user_follower_id uint) error
}

type UserRepository interface {
	GetAll() ([]User, error)
	Store(user User) (User, error)
	Get(id uint) (User, error)
	Update(userDB *User, user User) (User, error)
	Delete(id uint) (User, error)
	GetByEmail(email string) (User, error)

	GetFollowerNumber(id uint) (int, error)
	GetFollowingNumber(id uint) (int, error)
	Follow(user, user_follower User) (User, error)
	Unfollow(user, user_follower User) (User, error)
	GetFollower(user User) ([]User, error)
	GetFollowing(user User) ([]User, error)
}
