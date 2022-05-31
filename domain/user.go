package domain

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name      string `json:"name"      validate:"required"`
	Email     string `json:"email"     validate:"required,email"`
	Password  string `json:"password"  validate:"required,min=6"`
	Is_banned int    `json:"is_banned" validate:"required"`
}

type UserUsecase interface {
	GetAll() ([]User, error)
	Store(user User) (User, error)
	Get(id uint) (User, error)
	Update(user User, id uint) (User, error)
	Delete(id uint) (User, error)
}

type UserRepository interface {
	GetAll() ([]User, error)
	Store(user User) (User, error)
	Get(id uint) (User, error)
	Update(userDB *User, user User) (User, error)
	Delete(id uint) (User, error)
	GetByEmail(email string) (User, error)
}
