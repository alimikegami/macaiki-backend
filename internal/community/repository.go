package community

import (
	"macaiki/internal/community/entity"
	communityEntity "macaiki/internal/community/entity"
	threadEntity "macaiki/internal/thread/entity"
	userEntity "macaiki/internal/user/entity"
)

type CommunityRepository interface {
	GetAllCommunities(userID uint, search string) ([]communityEntity.Community, error)
	GetCommunityWithDetail(userID, communityID uint) (communityEntity.Community, error)
	GetCommunity(id uint) (communityEntity.Community, error)
	GetCommunityThread(userID, communityID uint) ([]threadEntity.ThreadWithDetails, error)
	GetCommunityAbout(userID, communityID uint) (communityEntity.Community, error)
	StoreCommunity(community communityEntity.Community) error
	UpdateCommunity(community communityEntity.Community, communityReq communityEntity.Community) (communityEntity.Community, error)
	DeleteCommunity(communityID uint) error

	FollowCommunity(user userEntity.User, community communityEntity.Community) error
	UnfollowCommunity(user userEntity.User, community communityEntity.Community) error

	SetCommunityImage(id uint, imageURL string, tableName string) error
	AddModerator(user userEntity.User, community communityEntity.Community) error
	RemoveModerator(user userEntity.User, community communityEntity.Community) error
	GetModeratorByCommunityID(userID, communityID uint) ([]userEntity.User, error)
	GetModeratorByUserID(userID, communityID uint) (entity.CommunityModerator, error)

	StoreReportCommunity(communityReport communityEntity.CommunityReport) error
	UpdateReportCommunity(communityReport communityEntity.CommunityReport, userID uint) error
	GetReportCommunity(id uint) (communityEntity.CommunityReport, error)
	GetReports(communityID uint) ([]entity.BriefReport, error)
}
