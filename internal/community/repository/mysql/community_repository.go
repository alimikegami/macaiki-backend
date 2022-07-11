package mysql

import (
	"errors"
	"fmt"
	"macaiki/internal/community"
	"macaiki/internal/community/entity"
	communityEntity "macaiki/internal/community/entity"
	threadEntity "macaiki/internal/thread/entity"
	userEntity "macaiki/internal/user/entity"
	"macaiki/pkg/utils"

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

	res := cr.db.Raw("SELECT c.*, !isnull(cf.user_id) AS is_followed, !isnull(cm.user_id) AS is_moderator FROM `communities` AS c LEFT JOIN (SELECT * FROM community_followers WHERE user_id = ?) AS cf ON c.id = cf.community_id LEFT JOIN (SELECT * FROM community_moderators WHERE user_id = ?) AS cm ON c.id = cm.community_id WHERE c.deleted_at IS NULL AND c.name LIKE ? ORDER BY is_moderator DESC", userID, userID, "%"+search+"%").Scan(&communities)
	err := res.Error
	if err != nil {
		return []communityEntity.Community{}, err
	}

	return communities, nil
}

func (cr *CommunityRepositoryImpl) GetCommunityWithDetail(userID, communityID uint) (communityEntity.Community, error) {
	community := communityEntity.Community{}

	res := cr.db.Raw("SELECT c.*, !isnull(cf.user_id) AS is_followed, !isnull(cm.user_id) AS is_moderator FROM `communities` AS c LEFT JOIN (SELECT * FROM community_followers WHERE user_id = ?) AS cf ON c.id = cf.community_id LEFT JOIN (SELECT * FROM community_moderators WHERE user_id = ?) AS cm ON c.id = cm.community_id WHERE c.deleted_at IS NULL AND c.id = ?", userID, userID, communityID).Scan(&community)
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

func (cr *CommunityRepositoryImpl) GetCommunityThread(userID, communityID uint) ([]threadEntity.ThreadWithDetails, error) {
	threads := []threadEntity.ThreadWithDetails{}
	res := cr.db.Raw("SELECT t.*, tlc.count AS upvotes_count, !isnull(tl.user_id) AS is_upvoted, u.*, (u.id = ?) AS is_mine FROM threads AS t LEFT JOIN (SELECT t.thread_id, COUNT(*) AS count FROM thread_upvotes AS t GROUP BY t.thread_id) AS tlc ON t.id = tlc.thread_id LEFT JOIN (SELECT * FROM thread_upvotes WHERE user_id = ?) AS tl ON tl.thread_id = t.id LEFT JOIN (SELECT u.*, !ISNULL(uf.user_id) AS is_followed FROM `users` AS u LEFT JOIN (SELECT * FROM user_followers WHERE follower_id = ?) AS uf ON u.id = uf.user_id WHERE u.deleted_at IS NULL) AS u ON u.id = t.user_id WHERE t.community_id = ? AND t.deleted_at IS NULL", userID, userID, userID, communityID).Scan(&threads)
	err := res.Error
	if err != nil {
		return []threadEntity.ThreadWithDetails{}, err
	}

	return threads, nil
}

func (cr *CommunityRepositoryImpl) GetCommunityAbout(userID, communityID uint) (communityEntity.Community, error) {
	community := communityEntity.Community{}

	res := cr.db.Raw("SELECT c.*, cm.total_moderators, cf.total_followers, cm2.is_moderator FROM communities AS c LEFT JOIN (SELECT community_id, COUNT(*) AS total_followers FROM community_followers GROUP BY community_id) AS cf ON c.id = cf.community_id LEFT JOIN (SELECT community_id, COUNT(*) AS total_moderators FROM community_moderators GROUP BY community_id) AS cm ON c.id = cm.community_id LEFT JOIN (SELECT !ISNULL(user_id) AS is_moderator FROM community_moderators WHERE user_id = ?) AS cm2 ON  c.id = cm.community_id WHERE c.id = ?", userID, communityID).Scan(&community)
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
	err := cr.db.Model(&community).Association("Followers").Append(&user)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CommunityRepositoryImpl) UnfollowCommunity(user userEntity.User, community communityEntity.Community) error {
	err := cr.db.Model(&community).Association("Followers").Delete(&user)
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

func (cr *CommunityRepositoryImpl) AddModerator(user userEntity.User, community communityEntity.Community) error {
	err := cr.db.Model(&community).Association("Moderators").Append(&user)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CommunityRepositoryImpl) RemoveModerator(user userEntity.User, community communityEntity.Community) error {
	err := cr.db.Model(&community).Association("Moderators").Delete(&user)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CommunityRepositoryImpl) GetModeratorByCommunityID(userID, communityID uint) ([]userEntity.User, error) {
	users := []userEntity.User{}

	res := cr.db.Raw("SELECT u.*, !ISNULL(uf.user_id) AS is_followed, (u.id = ?) AS is_mine FROM `users` AS u LEFT JOIN (SELECT * FROM user_followers WHERE follower_id = ?) AS uf ON u.id = uf.user_id LEFT JOIN community_moderators AS cm ON cm.user_id = u.id WHERE u.deleted_at IS NULL AND cm.community_id = ?", userID, userID, communityID).Scan(&users)
	err := res.Error
	if err != nil {
		return []userEntity.User{}, err
	}

	return users, nil
}

func (cr *CommunityRepositoryImpl) StoreReportCommunity(communityReport communityEntity.CommunityReport) error {
	res := cr.db.Create(&communityReport)
	err := res.Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *CommunityRepositoryImpl) GetReports(communityID uint) ([]entity.BriefReport, error) {
	var reports []entity.BriefReport

	res := cr.db.Raw("SELECT tr.id AS 'thread_reports_id', NULL AS 'community_reports_id', NULL AS 'comment_reports_id', tr.created_at, tr.user_id, tr.thread_id, NULL AS community_reported_id, NULL as comment_id, rc.name AS report_category, u.username, u.profile_image_url, 'threads' AS type FROM thread_reports tr INNER JOIN report_categories rc ON tr.report_category_id = rc.id INNER JOIN users u ON u.id = tr.user_id INNER JOIN threads t ON tr.thread_id = t.id WHERE tr.deleted_at IS NULL AND tr.user_id NOT IN (SELECT cm.user_id FROM community_moderators cm) AND t.community_id = ? UNION SELECT NULL AS 'thread_reports_id', cr2.id AS 'community_reports_id', NULL AS 'comment_reports_id', cr2.created_at, cr2.user_id, NULL AS thread_id, cr2.community_reported_id, NULL AS comment_id, rc.name AS report_category, u.username, u.profile_image_url, 'users' AS type FROM community_reports cr2 INNER JOIN report_categories rc ON cr2.report_category_id = rc.id INNER JOIN users u ON u.id = cr2.user_id  WHERE cr2.deleted_at IS NULL AND cr2.user_id NOT IN (SELECT cm.user_id FROM community_moderators cm) AND cr2.community_reported_id = ? UNION SELECT NULL AS 'thread_reports_id', NULL AS 'community_reports_id', cr.id AS 'comment_reports_id', cr.created_at, cr.user_id, NULL AS thread_id, NULL AS community_reported_id, cr.comment_id, rc.name AS report_category, u.username, u.profile_image_url, 'comments' AS type FROM comment_reports cr INNER JOIN report_categories rc ON cr.report_category_id = rc.id INNER JOIN users u ON u.id = cr.user_id INNER JOIN comments c ON c.id = cr.comment_id INNER JOIN threads t ON c.thread_id = t.id WHERE cr.deleted_at IS NULL AND cr.user_id NOT IN (SELECT cm.user_id FROM community_moderators cm) AND t.community_id = ?;", communityID, communityID, communityID).Scan(&reports)

	if res.Error != nil {
		return []entity.BriefReport{}, utils.ErrInternalServerError
	}

	return reports, nil
}

func (cr *CommunityRepositoryImpl) GetModeratorByUserID(userID uint) (entity.CommunityModerator, error) {
	var communityMods entity.CommunityModerator

	res := cr.db.First(&communityMods, userID)

	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error.Error() == "record not found" {
			return entity.CommunityModerator{}, utils.ErrNotFound
		}
		return entity.CommunityModerator{}, utils.ErrInternalServerError
	}

	return communityMods, nil
}
