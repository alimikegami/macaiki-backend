package mysql

import (
	"errors"
	"macaiki/internal/community"
	communityEntity "macaiki/internal/community/entity"
	threadEntity "macaiki/internal/thread/entity"
	userEntity "macaiki/internal/user/entity"

	"gorm.io/gorm"
)

type CommunityRepositoryImpl struct {
	db *gorm.DB
}

func NewCommunityRepository(db *gorm.DB) community.CommunityRepository {
	return &CommunityRepositoryImpl{db}
}

func (cr *CommunityRepositoryImpl) GetAllCommunities(userID uint, search string) ([]communityEntity.Community, error) {
	communities := []communityEntity.Community{}

	res := cr.db.Raw("SELECT c.*, !ISNULL(cf.user_id) AS `is_followed` FROM `communities` AS c LEFT JOIN (SELECT * FROM community_followers WHERE user_id = ?) AS cf ON c.id = cf.community_id WHERE c.deleted_at IS NULL", userID).Scan(&communities)
	err := res.Error
	if err != nil {
		return []communityEntity.Community{}, err
	}

	return communities, nil
}

func (cr *CommunityRepositoryImpl) GetCommunityWithDetail(userID, communityID uint) (communityEntity.Community, error) {
	community := communityEntity.Community{}

	res := cr.db.Raw("SELECT c.*, !ISNULL(cf.user_id) AS `is_followed` FROM `communities` AS c LEFT JOIN (SELECT * FROM community_followers WHERE user_id = ?) AS cf ON c.id = cf.community_id WHERE c.id = ? AND c.deleted_at IS NULL", userID, communityID).Scan(&community)
	err := res.Error

	if err != nil {
		return communityEntity.Community{}, err
	}

	return community, nil
}

func (cr *CommunityRepositoryImpl) GetCommunity(id uint) (communityEntity.Community, error) {
	community := communityEntity.Community{}

	res := cr.db.Find(&community, id)
	err := res.Error

	if err != nil {
		return communityEntity.Community{}, err
	}

	return community, nil
}

func (cr *CommunityRepositoryImpl) StoreCommunity(community communityEntity.Community) error {
	res := cr.db.Create(&community)
	err := res.Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *CommunityRepositoryImpl) UpdateCommunity(community communityEntity.Community, communityReq communityEntity.Community) (communityEntity.Community, error) {

	res := cr.db.Model(&community).Updates(communityReq)
	err := res.Error
	if err != nil {
		return communityEntity.Community{}, err
	}

	return community, nil
}

func (cr *CommunityRepositoryImpl) DeleteCommunity(community communityEntity.Community) error {
	res := cr.db.Delete(&community)
	err := res.Error
	if err != nil {
		return err
	}
	return nil
}
func (cr *CommunityRepositoryImpl) FollowCommunity(user userEntity.User, community communityEntity.Community) error {
	err := cr.db.Model(&community).Association("Users").Append(&user)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CommunityRepositoryImpl) UnfollowCommunity(user userEntity.User, community communityEntity.Community) error {
	err := cr.db.Model(&community).Association("Users").Delete(&user)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CommunityRepositoryImpl) SetCommunityImage(id uint, imageURL string, tableName string) error {
	res := cr.db.Model(&communityEntity.Community{}).Where("id = ?", id).Update(tableName, imageURL)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected < 1 {
		return errors.New("resource does not exists")
	}

	return nil
}

func (cr *CommunityRepositoryImpl) GetThreadCommunityByID(userID, communityID uint) ([]threadEntity.ThreadWithDetails, error) {
	threads := []threadEntity.ThreadWithDetails{}
	res := cr.db.Raw("SELECT t.*, tlc.count AS likes_count, !isnull(tl.user_id) AS is_liked, u.*, (u.id = ?) AS is_mine FROM threads AS t LEFT JOIN (SELECT t.thread_id, COUNT(*) AS count FROM thread_likes AS t GROUP BY t.thread_id) AS tlc ON t.id = tlc.thread_id LEFT JOIN (SELECT * FROM thread_likes WHERE user_id = ?) AS tl ON tl.thread_id = t.id LEFT JOIN (SELECT u.*, !ISNULL(uf.user_id) AS is_followed FROM `users` AS u LEFT JOIN (SELECT * FROM user_followers WHERE follower_id = ?) AS uf ON u.id = uf.user_id WHERE u.deleted_at IS NULL) AS u ON u.id = t.user_id WHERE t.community_id = ?", userID, userID, userID, communityID).Scan(&threads)
	err := res.Error
	if err != nil {
		return []threadEntity.ThreadWithDetails{}, err
	}

	return threads, nil
}

func (cr *CommunityRepositoryImpl) AddModerator(user userEntity.User, community communityEntity.Community) error {
	err := cr.db.Model(&community).Association("Moderators").Append(&user)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CommunityRepositoryImpl) RemoveModerator(user userEntity.User, community communityEntity.Community) error {
	err := cr.db.Model(&community).Association("Moderators").Delete(user)
	if err != nil {
		return err
	}

	return nil
}
