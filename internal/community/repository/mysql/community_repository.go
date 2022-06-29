package mysql

import (
	"macaiki/internal/domain"

	"gorm.io/gorm"
)

type CommunityRepositoryImpl struct {
	db *gorm.DB
}

func NewCommunityRepository(db *gorm.DB) domain.CommunityRepository {
	return &CommunityRepositoryImpl{db}
}

func (cr *CommunityRepositoryImpl) GetAllCommunity(search string) ([]domain.Community, error) {
	communities := []domain.Community{}

	res := cr.db.Where("name LIKE ?", "%"+search+"%").Find(&communities)
	err := res.Error
	if err != nil {
		return []domain.Community{}, err
	}

	return communities, nil
}

func (cr *CommunityRepositoryImpl) GetCommunity(id uint) (domain.Community, error) {
	community := domain.Community{}

	res := cr.db.Find(&community, id)
	err := res.Error

	if err != nil {
		return domain.Community{}, err
	}

	return community, nil
}

func (cr *CommunityRepositoryImpl) StoreCommunity(community domain.Community) error {
	res := cr.db.Create(&community)
	err := res.Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *CommunityRepositoryImpl) UpdateCommunity(community domain.Community, communityReq domain.Community) error {

	res := cr.db.Model(&community).Updates(communityReq)
	err := res.Error
	if err != nil {
		return err
	}

	return nil
}
func (cr *CommunityRepositoryImpl) DeleteCommunity(community domain.Community) error {
	res := cr.db.Delete(&community)
	err := res.Error
	if err != nil {
		return err
	}
	return nil
}