package entity

import (
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
	Profession         string
	Role               string
	IsBanned           int
	Followers          []User       `gorm:"many2many:user_followers"`
	Report             []UserReport `gorm:"foreignKey:UserID"`
	Reported           []UserReport `gorm:"foreignKey:ReportedUserID"`
	IsFollowed         int          `gorm:"-:migration;<-:false"`
	IsMine             int          `gorm:"-:migration;<-:false"`
}

type UserReport struct {
	UserID           uint `gorm:"primaryKey"`
	ReportedUserID   uint `gorm:"primaryKey"`
	ReportCategoryID uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
