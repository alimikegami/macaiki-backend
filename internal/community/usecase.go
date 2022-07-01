package community

import "macaiki/internal/community/dto"

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
