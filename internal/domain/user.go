package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Role      string `json:"role" validate:"required"`
	IsBanned  int    `json:"isBanned" validate:"required"`
	ImageUrl  string `json:"imageUrl"`
	Followers []User `json:"followers" gorm:"many2many:user_followers"`
}

type UserUsecase interface {
	Login(email, password string) (string, error)
	Register(user User) (User, error)
	GetAll() ([]User, error)
	Get(id uint) (User, []User, error)
	Update(user User, id uint) (User, error)
	Delete(id uint) (User, error)
	Follow(user_id, user_follower_id uint) (User, error)
	Unfollow(user_id, user_follower_id uint) (User, error)
}

type UserRepository interface {
	GetAll() ([]User, error)
	Store(user User) (User, error)
	Get(id uint) (User, error)
	Update(userDB *User, user User) (User, error)
	Delete(id uint) (User, error)
	GetByEmail(email string) (User, error)

	Follow(user, user_follower User) (User, error)
	Unfollow(user, user_follower User) (User, error)
	GetFollowing(user User) ([]User, error)
}
