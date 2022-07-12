package entity

import (
	userEntity "macaiki/internal/user/entity"
	"time"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name                        string
	CommunityImageUrl           string
	CommunityBackgroundImageUrl string
	Description                 string
	Followers                   []userEntity.User `gorm:"many2many:community_followers;"`
	Moderators                  []userEntity.User `gorm:"many2many:community_moderators;"`
	IsFollowed                  int               `gorm:"-:migration;<-:false"`
	IsModerator                 int               `gorm:"-:migration;<-:false"`
	TotalFollowers              int               `gorm:"-:migration;<-:false"`
	TotalModerators             int               `gorm:"-:migration;<-:false"`
}

type CommunityReport struct {
	gorm.Model
	UserID              uint
	CommunityReportedID uint
	ReportCategoryID    uint
}

type BriefReport struct {
	ThreadReportsID     uint
	CommunityReportsID  uint
	CommentReportsID    uint
	CreatedAt           time.Time
	UserID              uint
	ThreadID            uint
	CommunityReportedID uint
	CommentID           uint
	ReportCategory      string
	Username            string
	ProfileImageURL     string
	Type                string
}

type CommunityModerator struct {
	UserID      uint
	CommunityID uint
}
