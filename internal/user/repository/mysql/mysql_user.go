package mysql

import (
	"macaiki/internal/domain"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Db *gorm.DB
}

func NewMysqlUserRepository(Db *gorm.DB) domain.UserRepository {
	return &MysqlUserRepository{Db}
}

func (ur *MysqlUserRepository) GetAll() ([]domain.User, error) {
	users := []domain.User{}

	res := ur.Db.Find(&users)
	err := res.Error
	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (ur *MysqlUserRepository) Store(user domain.User) (domain.User, error) {
	res := ur.Db.Create(&user)
	err := res.Error
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *MysqlUserRepository) Get(id uint) (domain.User, error) {
	user := domain.User{}

	res := ur.Db.Preload("Followers").Find(&user, id)
	err := res.Error

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *MysqlUserRepository) Update(userDB *domain.User, user domain.User) (domain.User, error) {
	if user.Password == "" {
		user.Password = userDB.Password
	}
	user.ID = userDB.ID

	res := ur.Db.Model(&userDB).Updates(user)
	err := res.Error
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *MysqlUserRepository) Delete(id uint) (domain.User, error) {
	user, err := ur.Get(id)
	if err != nil {
		return domain.User{}, err
	}

	res := ur.Db.Delete(&user, "id = ?", id)
	err = res.Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (ur *MysqlUserRepository) GetByEmail(email string) (domain.User, error) {
	user := domain.User{}

	res := ur.Db.Find(&user, "email = ?", email)
	err := res.Error
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *MysqlUserRepository) Follow(user, user_follower domain.User) (domain.User, error) {
	err := ur.Db.Model(&user).Association("Followers").Append(&user_follower)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *MysqlUserRepository) Unfollow(user, user_follower domain.User) (domain.User, error) {
	err := ur.Db.Model(&user).Association("Followers").Delete(&user_follower)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (ur *MysqlUserRepository) GetFollowing(user domain.User) ([]domain.User, error) {
	users := []domain.User{}

	res := ur.Db.Raw("SELECT * FROM `users` LEFT JOIN `user_followers` `Followers` ON `users`.`id` = `Followers`.`user_id` WHERE `Followers`.`follower_id` = ?", user.ID).Scan(&users)
	err := res.Error

	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}
