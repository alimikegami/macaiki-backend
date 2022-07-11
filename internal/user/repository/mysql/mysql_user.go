package mysql

import (
	"errors"
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

func (ur *MysqlUserRepository) GetReports() ([]entity.BriefReport, error) {
	var reports []entity.BriefReport
	res := ur.Db.Raw("SELECT tr.id AS 'thread_reports_id', NULL AS 'user_reports_id', NULL AS 'comment_reports_id', tr.created_at, tr.user_id, tr.thread_id, NULL AS reported_user_id, NULL as comment_id, rc.name AS report_category, u.username, u.profile_image_url, 'threads' AS type FROM thread_reports tr INNER JOIN report_categories rc ON tr.report_category_id = rc.id INNER JOIN users u ON u.id = tr.user_id WHERE tr.deleted_at IS NULL UNION SELECT NULL AS 'thread_reports_id', ur.id AS 'user_reports_id', NULL AS 'comment_reports_id', ur.created_at, ur.user_id, NULL AS thread_id, ur.reported_user_id, NULL AS comment_id, rc.name AS report_category, u.username, u.profile_image_url, 'users' AS type FROM user_reports ur INNER JOIN report_categories rc ON ur.report_category_id = rc.id INNER JOIN users u ON u.id = ur.user_id WHERE ur.deleted_at IS NULL UNION SELECT NULL AS 'thread_reports_id', NULL AS 'user_reports_id', cr.id AS 'comment_reports_id', cr.created_at, cr.user_id, NULL AS thread_id, NULL AS reported_user_id, cr.comment_id, rc.name AS report_category, u.username, u.profile_image_url, 'comments' AS type FROM comment_reports cr INNER JOIN report_categories rc ON cr.report_category_id = rc.id INNER JOIN users u ON u.id = cr.user_id WHERE cr.deleted_at IS NULL;").Scan(&reports)

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
