package domain

import (
	"macaiki/internal/community/dto"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name                        string
	CommunityImageUrl           string
	CommunityBackgroundImageUrl string
	Description                 string
	Users                       []User `gorm:"many2many:community_followers;"`
}

type CommunityWithDetail struct {
	gorm.Model
	Name                        string
	CommunityImageUrl           string
	CommunityBackgroundImageUrl string
	Description                 string
	IsFollowed                  bool
}

type CommunityUsecase interface {
	GetAllCommunity(search string) ([]dto.CommunityResponse, error)
	GetAllCommunityDetail(userID int, search string) ([]dto.CommunityWithDetailResponse, error)
	GetCommunity(id uint) (dto.CommunityResponse, error)
	StoreCommunity(community dto.CommunityRequest, role string) error
	UpdateCommunity(id uint, community dto.CommunityRequest, role string) error
	DeleteCommunity(id uint, role string) error

	FollowCommunity(userID, communityID uint) error
	UnfollowCommunity(userID, communityID uint) error
}

type CommunityRepository interface {
	GetAllCommunity(search string) ([]Community, error)
	GetAllCommunityDetail(userID, search string) ([]CommunityWithDetail, error)
	GetCommunity(id uint) (Community, error)
	StoreCommunity(community Community) error
	UpdateCommunity(community Community, communityReq Community) error
	DeleteCommunity(community Community) error

	FollowCommunity(user User, community Community) error
	UnfollowCommunity(user User, community Community) error
}
