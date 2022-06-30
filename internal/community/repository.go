package community

import "macaiki/internal/community/entity"

type CommunityRepository interface {
	GetAllCommunity(search string) ([]entity.Community, error)
	GetCommunity(id uint) (entity.Community, error)
	StoreCommunity(community entity.Community) error
	UpdateCommunity(community entity.Community, communityReq entity.Community) error
	DeleteCommunity(community entity.Community) error
}
