package community

import (
	communityEntity "macaiki/internal/community/entity"
	userEntity "macaiki/internal/user/entity"
)

type CommunityRepository interface {
	GetAllCommunity(search string) ([]communityEntity.Community, error)
	GetAllCommunityDetail(userID, search string) ([]communityEntity.CommunityWithDetail, error)
	GetCommunity(id uint) (communityEntity.Community, error)
	StoreCommunity(community communityEntity.Community) error
	UpdateCommunity(community communityEntity.Community, communityReq communityEntity.Community) error
	DeleteCommunity(community communityEntity.Community) error

	FollowCommunity(user userEntity.User, community communityEntity.Community) error
	UnfollowCommunity(user userEntity.User, community communityEntity.Community) error
}
