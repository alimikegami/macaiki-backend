package mysql

import (
	"macaiki/domain"

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

	res := ur.Db.Find(&user, id)
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
