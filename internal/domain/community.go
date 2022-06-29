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
}

type CommunityUsecase interface {
	GetAllCommunity(search string) ([]dto.CommunityResponse, error)
	GetCommunity(id uint) (dto.CommunityResponse, error)
	StoreCommunity(community dto.CommunityRequest, role string) error
	UpdateCommunity(id uint, community dto.CommunityRequest, role string) error
	DeleteCommunity(id uint, role string) error
}

type CommunityRepository interface {
	GetAllCommunity(search string) ([]Community, error)
	GetCommunity(id uint) (Community, error)
	StoreCommunity(community Community) error
	UpdateCommunity(community Community, communityReq Community) error
	DeleteCommunity(community Community) error
}
