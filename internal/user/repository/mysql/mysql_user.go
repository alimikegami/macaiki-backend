package mysql

import (
	"errors"
	"fmt"
	communityEntity "macaiki/internal/community/entity"
	threadEntity "macaiki/internal/thread/entity"
	"macaiki/internal/user"
	"macaiki/internal/user/entity"
	"macaiki/pkg/utils"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Db *gorm.DB
}

func NewMysqlUserRepository(Db *gorm.DB) user.UserRepository {
	return &MysqlUserRepository{Db}
}

func (ur *MysqlUserRepository) GetAllWithDetail(userID uint, search string) ([]entity.User, error) {
	users := []entity.User{}

	res := ur.Db.Raw("SELECT u.*, !ISNULL(uf.user_id) AS is_followed, (u.id = ?) AS is_mine FROM `users` AS u LEFT JOIN (SELECT * FROM user_followers WHERE follower_id = ?) AS uf ON u.id = uf.user_id WHERE u.deleted_at IS NULL AND (u.username LIKE ? OR u.name LIKE ?) ", userID, userID, "%"+search+"%", "%"+search+"%").Find(&users)
	err := res.Error
	if err != nil {
		return []entity.User{}, err
	}

	return users, nil
}

func (ur *MysqlUserRepository) Store(user entity.User) error {
	res := ur.Db.Create(&user)
	err := res.Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *MysqlUserRepository) Get(id uint) (entity.User, error) {
	user := entity.User{}

	res := ur.Db.Raw("SELECT u.*, !ISNULL(uf.user_id) AS is_followed, (u.id = ?) AS is_mine  FROM `users` AS u LEFT JOIN (SELECT * FROM user_followers WHERE follower_id = ?) AS uf ON u.id = uf.user_id WHERE u.deleted_at IS NULL AND u.id = ?", id, id, id).Find(&user)
	err := res.Error

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (ur *MysqlUserRepository) Update(userDB *entity.User, user entity.User) (entity.User, error) {
	if user.Password == "" {
		user.Password = userDB.Password
	}
	user.ID = userDB.ID

	res := ur.Db.Model(&userDB).Updates(user)
	err := res.Error
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (ur *MysqlUserRepository) Delete(id uint) error {
	user, err := ur.Get(id)
	if err != nil {
		return err
	}

	res := ur.Db.Delete(&user, "id = ?", id)
	err = res.Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *MysqlUserRepository) GetByEmail(email string) (entity.User, error) {
	user := entity.User{}

	res := ur.Db.Find(&user, "email = ?", email)
	err := res.Error
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (ur *MysqlUserRepository) GetByUsername(username string) (entity.User, error) {
	user := entity.User{}

	res := ur.Db.Find(&user, "username = ?", username)
	err := res.Error
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (ur *MysqlUserRepository) Follow(user, userFollower entity.User) (entity.User, error) {
	err := ur.Db.Model(&user).Association("Followers").Append(&userFollower)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (ur *MysqlUserRepository) Unfollow(user, userFollower entity.User) (entity.User, error) {
	err := ur.Db.Model(&user).Association("Followers").Delete(&userFollower)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (ur *MysqlUserRepository) GetFollowerNumber(id uint) (int, error) {
	var count int64
	res := ur.Db.Raw("SELECT COUNT(*) FROM `users` LEFT JOIN `user_followers` `Followers` ON `users`.`id` = `Followers`.`follower_id` WHERE `Followers`.`user_id` = ? AND `users`.`deleted_at` IS NULL", id).Scan(&count)
	err := res.Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (ur *MysqlUserRepository) GetFollowingNumber(id uint) (int, error) {
	var count int64
	res := ur.Db.Raw("SELECT COUNT(*) FROM `users` LEFT JOIN `user_followers` `Followers` ON `users`.`id` = `Followers`.`user_id` WHERE `Followers`.`follower_id` = ? AND `users`.`deleted_at` IS NULL", id).Scan(&count)
	err := res.Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (ur *MysqlUserRepository) GetThreadsNumber(id uint) (int, error) {
	var count int64
	res := ur.Db.Table("threads").Where("user_id = ?", id).Count(&count)
	err := res.Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (ur *MysqlUserRepository) GetFollower(userID, getFollowingUserID uint) ([]entity.User, error) {
	users := []entity.User{}

	res := ur.Db.Raw("SELECT u.*, !ISNULL(uf2.user_id) AS is_followed, (u.id=?) AS is_mine FROM users AS u LEFT JOIN user_followers AS uf ON u.id = uf.follower_id LEFT JOIN (SELECT * FROM user_followers WHERE user_followers.follower_id = ?) AS uf2 ON u.id = uf2.user_id WHERE uf.user_id = ? AND u.deleted_at IS NULL", userID, userID, getFollowingUserID).Scan(&users)
	err := res.Error

	if err != nil {
		return []entity.User{}, err
	}

	return users, nil
}

func (ur *MysqlUserRepository) GetFollowing(userID, getFollowingUserID uint) ([]entity.User, error) {
	users := []entity.User{}
	res := ur.Db.Raw("SELECT u.*, !ISNULL(uf2.user_id) AS is_followed, (u.id=?) AS is_mine FROM users AS u LEFT JOIN user_followers uf ON u.id = uf.user_id LEFT JOIN (SELECT * FROM user_followers WHERE user_followers.follower_id = ?) AS uf2 ON u.id = uf2.user_id WHERE uf.follower_id = ? AND u.deleted_at IS NULL", userID, userID, getFollowingUserID).Scan(&users)
	err := res.Error

	if err != nil {
		return []entity.User{}, err
	}

	return users, nil
}

func (ur *MysqlUserRepository) SetUserImage(id uint, imageURL string, tableName string) error {
	res := ur.Db.Model(&entity.User{}).Where("id = ?", id).Update(tableName, imageURL)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected < 1 {
		return errors.New("resource does not exists")
	}

	return nil
}

func (ur *MysqlUserRepository) StoreReport(userReport entity.UserReport) error {
	res := ur.Db.Create(&userReport)
	err := res.Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *MysqlUserRepository) StoreOTP(VerifyEmail entity.VerificationEmail) error {
	res := ur.Db.Create(&VerifyEmail)
	err := res.Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *MysqlUserRepository) GetOTP(email string) (entity.VerificationEmail, error) {
	VerifyEmail := entity.VerificationEmail{}
	res := ur.Db.Where("email = ?", email).Order("id desc").First(&VerifyEmail)
	err := res.Error
	if err != nil {
		return entity.VerificationEmail{}, err
	}

	return VerifyEmail, nil
}

func (ur *MysqlUserRepository) GetReports() ([]entity.BriefReport, error) {
	var reports []entity.BriefReport
	res := ur.Db.Raw("SELECT tr.id AS 'thread_reports_id', NULL AS 'user_reports_id', NULL AS 'comment_reports_id', tr.created_at, tr.user_id, tr.thread_id, NULL AS reported_user_id, NULL as comment_id, rc.name AS report_category, u.username, u.profile_image_url, 'threads' AS type FROM thread_reports tr INNER JOIN report_categories rc ON tr.report_category_id = rc.id INNER JOIN users u ON u.id = tr.user_id WHERE tr.deleted_at IS NULL AND u.`role` = 'Moderator' UNION SELECT NULL AS 'thread_reports_id', ur.id AS 'user_reports_id', NULL AS 'comment_reports_id', ur.created_at, ur.user_id, NULL AS thread_id, ur.reported_user_id, NULL AS comment_id, rc.name AS report_category, u.username, u.profile_image_url, 'users' AS type FROM user_reports ur INNER JOIN report_categories rc ON ur.report_category_id = rc.id INNER JOIN users u ON u.id = ur.user_id WHERE ur.deleted_at IS NULL AND u.`role` = 'Moderator' UNION SELECT NULL AS 'thread_reports_id', NULL AS 'user_reports_id', cr.id AS 'comment_reports_id', cr.created_at, cr.user_id, NULL AS thread_id, NULL AS reported_user_id, cr.comment_id, rc.name AS report_category, u.username, u.profile_image_url, 'comments' AS type FROM comment_reports cr INNER JOIN report_categories rc ON cr.report_category_id = rc.id INNER JOIN users u ON u.id = cr.user_id WHERE cr.deleted_at IS NULL AND u.`role` = 'Moderator';").Scan(&reports)

	if res.Error != nil {
		return []entity.BriefReport{}, utils.ErrInternalServerError
	}

	return reports, nil
}

func (ur *MysqlUserRepository) GetDashboardAnalytics() (entity.AdminDashboardAnalytics, error) {
	var adminAnalytics entity.AdminDashboardAnalytics

	res := ur.Db.Raw("SELECT (SELECT COUNT(*) FROM users WHERE deleted_at IS NULL) AS users_count, (SELECT COUNT(*) FROM users WHERE `role` = 'Moderator') AS moderators_count, (SELECT COUNT(*) FROM thread_reports) + (SELECT COUNT(*) FROM user_reports) + (SELECT COUNT(*) FROM comment_reports) AS reports_count;").Scan(&adminAnalytics)

	if res.Error != nil {
		return entity.AdminDashboardAnalytics{}, utils.ErrInternalServerError
	}

	return adminAnalytics, nil
}

func (ur *MysqlUserRepository) GetReportedThread(threadReportID uint) (entity.ReportedThread, error) {
	var reportedThread entity.ReportedThread

	res := ur.Db.Raw("SELECT tr.id, t.title AS thread_title, t.body AS thread_body, t.image_url AS thread_image_url, t.created_at AS thread_created_at, t2.likes_count, u.username AS reported_username, u.profile_image_url AS reported_profile_image_url, u.profession AS reported_user_profession, rc.name AS report_category, tr.created_at AS report_created_at, u2.username, u2.profile_image_url FROM thread_reports tr INNER JOIN threads t ON t.id = tr.thread_id LEFT JOIN (SELECT thread_id, COUNT(*) AS likes_count FROM thread_upvotes tu GROUP BY thread_id) AS t2 ON t.id = t2.thread_id INNER JOIN users u ON u.id = t.user_id INNER JOIN users u2 ON u2.id = tr.user_id INNER JOIN report_categories rc ON rc.id = tr.report_category_id WHERE tr.id = ?;", threadReportID).Scan(&reportedThread)


	if res.Error != nil {
		return entity.ReportedThread{}, utils.ErrInternalServerError
	}

	if reportedThread == (entity.ReportedThread{}) {
		return entity.ReportedThread{}, utils.ErrBadParamInput
	}

	return reportedThread, nil
}

func (ur *MysqlUserRepository) GetReportedCommunity(communityReportID uint) (entity.ReportedCommunity, error) {
	var reportedCommunity entity.ReportedCommunity

	res := ur.Db.Raw("SELECT cr.id, cr.created_at AS report_created_at, c.name AS community_name, c.community_image_url, c.community_background_image_url, rc.name AS report_category, u.username, u.profile_image_url  FROM community_reports cr INNER JOIN communities c ON c.id = cr.community_reported_id INNER JOIN users u ON u.id = cr.user_id INNER JOIN report_categories rc ON rc.id = cr.report_category_id WHERE cr.id = ?;", communityReportID).Scan(&reportedCommunity)

	if res.Error != nil {
		return entity.ReportedCommunity{}, utils.ErrInternalServerError
	}

	if reportedCommunity == (entity.ReportedCommunity{}) {
		return entity.ReportedCommunity{}, utils.ErrBadParamInput
	}

	return reportedCommunity, nil
}

func (ur *MysqlUserRepository) GetReportedComment(commentReportID uint) (entity.ReportedComment, error) {
	var reportedComment entity.ReportedComment

	res := ur.Db.Raw("SELECT cr.id, c.body AS comment_body, t2.likes_count, c.created_at AS comment_created_at, u.username, u.profile_image_url, u2.username AS reported_username, u2.profile_image_url AS reported_profile_image_url, rc.name AS report_category FROM comment_reports cr INNER JOIN comments c ON c.id = cr.comment_id INNER JOIN users u ON u.id = cr.user_id INNER JOIN users u2 ON c.user_id = u2.id INNER JOIN report_categories rc ON rc.id = cr.report_category_id LEFT JOIN (SELECT comment_id, COUNT(*) AS likes_count FROM comment_likes cl GROUP BY comment_id) AS t2 ON c.id = t2.comment_id WHERE cr.id = ?;", commentReportID).Scan(&reportedComment)

	if res.Error != nil {
		return entity.ReportedComment{}, utils.ErrInternalServerError
	}

	if reportedComment == (entity.ReportedComment{}) {
		return entity.ReportedComment{}, utils.ErrBadParamInput
	}

	return reportedComment, nil
}

func (ur *MysqlUserRepository) GetReportedUser(userReportID uint) (entity.ReportedUser, error) {
	var reportedUser entity.ReportedUser

	res := ur.Db.Raw("SELECT ur.id, u.username AS reported_user_username, u.name AS reported_user_name, u.profession AS reported_user_profession, u.bio AS reported_user_bio, u.profile_image_url AS reported_user_profile_image_url, u.background_image_url AS reported_user_background_url, u2.username AS reporting_user_username, u2.name AS reporting_user_name, reported_user_follower.followers_count, reported_user_following.following_count FROM user_reports ur INNER JOIN users u ON ur.reported_user_id = u.id INNER JOIN users u2 ON u2.id = ur.user_id LEFT JOIN (SELECT uf.follower_id, COUNT(*) AS following_count FROM user_followers uf GROUP BY uf.follower_id) reported_user_following ON reported_user_following.follower_id = ur.user_id LEFT JOIN (SELECT user_id, COUNT(*) AS followers_count FROM user_followers uf GROUP BY uf.user_id) reported_user_follower ON reported_user_follower.user_id = u.id WHERE ur.id = ?;", userReportID).Scan(&reportedUser)

	if res.Error != nil {
		return entity.ReportedUser{}, utils.ErrInternalServerError
	}

	if reportedUser == (entity.ReportedUser{}) {
		return entity.ReportedUser{}, utils.ErrBadParamInput
	}

	return reportedUser, nil
}

func (ur *MysqlUserRepository) DeleteUserReport(userReportID uint) error {
	res := ur.Db.Delete(&entity.UserReport{}, userReportID)

	if res.Error != nil {
		return utils.ErrInternalServerError
	}

	if res.RowsAffected < 1 {
		return utils.ErrNotFound
	}

	return nil
}

func (ur *MysqlUserRepository) DeleteThreadReport(threadReportID uint) error {
	res := ur.Db.Delete(&threadEntity.ThreadReport{}, threadReportID)

	if res.Error != nil {
		return utils.ErrInternalServerError
	}

	if res.RowsAffected < 1 {
		return utils.ErrNotFound
	}

	return nil
}

func (ur *MysqlUserRepository) DeleteCommunityReport(communityReportID uint) error {
	res := ur.Db.Delete(&communityEntity.CommunityReport{}, communityReportID)

	if res.Error != nil {
		return utils.ErrInternalServerError
	}

	if res.RowsAffected < 1 {
		return utils.ErrNotFound
	}

	return nil
}

func (ur *MysqlUserRepository) DeleteCommentReport(commentReportID uint) error {
	res := ur.Db.Delete(&threadEntity.CommentReport{}, commentReportID)

	if res.Error != nil {
		return utils.ErrInternalServerError
	}

	if res.RowsAffected < 1 {
		return utils.ErrNotFound
	}

	return nil
}

func (ur *MysqlUserRepository) GetUserReport(reportID uint) (entity.UserReport, error) {
	var userReport entity.UserReport

	res := ur.Db.First(&userReport, reportID)

	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error.Error() == "record not found" {
			return entity.UserReport{}, utils.ErrNotFound
		}
		return entity.UserReport{}, utils.ErrInternalServerError
	}

	return userReport, nil
}
