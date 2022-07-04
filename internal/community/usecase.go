package community

import (
	dtoCommunity "macaiki/internal/community/dto"
	"mime/multipart"
)

type CommunityUsecase interface {
	GetAllCommunities(userID int, search string) ([]dtoCommunity.CommunityDetailResponse, error)
	GetCommunity(userID, communityID uint) (dtoCommunity.CommunityDetailResponse, error)
	StoreCommunity(community dtoCommunity.CommunityRequest, role string) error
	UpdateCommunity(id uint, community dtoCommunity.CommunityRequest, role string) (dtoCommunity.CommunityResponse, error)
	DeleteCommunity(id uint, role string) error

	FollowCommunity(userID, communityID uint) error
	UnfollowCommunity(userID, communityID uint) error
	SetImage(id uint, img *multipart.FileHeader, role string) (string, error)
	SetBackgroundImage(id uint, img *multipart.FileHeader, role string) (string, error)

	GetThreadCommunity(userID, communityID uint) ([]dtoCommunity.DetailedThreadResponse, error)
	AddModerator(moderatorReq dto.CommunityModeratorRequest, role string) error
	RemoveModerator(moderatorReq dto.CommunityModeratorRequest, role string) error
}
