package mysql

import (
	"macaiki/internal/community"
	"macaiki/internal/community/entity"

	"gorm.io/gorm"
)

type CommunityRepositoryImpl struct {
	db *gorm.DB
}

func NewCommunityRepository(db *gorm.DB) community.CommunityRepository {
	return &CommunityRepositoryImpl{db}
}

func (cr *CommunityRepositoryImpl) GetAllCommunity(search string) ([]entity.Community, error) {
	communities := []entity.Community{}

	res := cr.db.Where("name LIKE ?", "%"+search+"%").Find(&communities)
	err := res.Error
	if err != nil {
		return []entity.Community{}, err
	}

	return communities, nil
}

func (cr *CommunityRepositoryImpl) GetCommunity(id uint) (entity.Community, error) {
	community := entity.Community{}

	res := cr.db.Find(&community, id)
	err := res.Error

	if err != nil {
		return entity.Community{}, err
	}

	return community, nil
}

func (cr *CommunityRepositoryImpl) StoreCommunity(community entity.Community) error {
	res := cr.db.Create(&community)
	err := res.Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *CommunityRepositoryImpl) UpdateCommunity(community entity.Community, communityReq entity.Community) error {

	res := cr.db.Model(&community).Updates(communityReq)
	err := res.Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *CommunityRepositoryImpl) DeleteCommunity(community entity.Community) error {
	res := cr.db.Delete(&community)
	err := res.Error
	if err != nil {
		return err
	}
	return nil
}
