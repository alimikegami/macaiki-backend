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
	Moderators                  []userEntity.User `gorm:"many2many:community_moderators;"`
	IsFollowed                  int               `gorm:"-:migration;<-:false"`
	IsModerator                 int               `gorm:"-:migration;<-:false"`
	TotalFollowers              int               `gorm:"-:migration;<-:false"`
	TotalModerators             int               `gorm:"-:migration;<-:false"`
}
