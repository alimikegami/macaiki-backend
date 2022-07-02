package entity

import (
	userEntity "macaiki/internal/user/entity"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name                        string
	CommunityImageUrl           string
	CommunityBackgroundImageUrl string
	Description                 string
	Users                       []userEntity.User `gorm:"many2many:community_followers;"`
	IsFollowed                  bool              `gorm:"-:migration;<-:false"`
}
