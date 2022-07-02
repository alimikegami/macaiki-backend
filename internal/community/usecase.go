package community

import "macaiki/internal/community/dto"

type CommunityUsecase interface {
	GetAllCommunities(userID int, search string) ([]dto.CommunityResponse, error)
	GetCommunity(userID, communityID uint) (dto.CommunityResponse, error)
	StoreCommunity(community dto.CommunityRequest, role string) error
	UpdateCommunity(id uint, community dto.CommunityRequest, role string) error
	DeleteCommunity(id uint, role string) error

	FollowCommunity(userID, communityID uint) error
	UnfollowCommunity(userID, communityID uint) error
}
