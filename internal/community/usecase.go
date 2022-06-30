package community

import "macaiki/internal/community/dto"

type CommunityUsecase interface {
	GetAllCommunity(search string) ([]dto.CommunityResponse, error)
	GetCommunity(id uint) (dto.CommunityResponse, error)
	StoreCommunity(community dto.CommunityRequest, role string) error
	UpdateCommunity(id uint, community dto.CommunityRequest, role string) error
	DeleteCommunity(id uint, role string) error
}
