package usercase

import (
	"macaiki/domain"

	"github.com/go-playground/validator/v10"
)

type userUsecase struct {
	userRepo  domain.UserRepository
	validator *validator.Validate
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: repo,
	}
}
func (uu *userUsecase) GetAll() ([]domain.User, error) {
	users, err := uu.userRepo.GetAll()
	if err != nil {
		return []domain.User{}, err
	}

	return users, err
}
func (uu *userUsecase) Store(user domain.User) (domain.User, error) {
	if err := uu.validator.Struct(user); err != nil {
		return domain.User{}, err
	}

	user, err := uu.userRepo.GetByEmail(user.Email)
	if err != nil {
		return domain.User{}, err
	}
	if user.ID != 0 {
		return domain.User{}, domain.ErrEmailAlreadyUsed
	}

	user, err = uu.userRepo.Store(user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
func (uu *userUsecase) Get(id uint) (domain.User, error) {
	user, err := uu.userRepo.Get(id)
	if err != nil {
		return domain.User{}, err
	}
	if user.ID == 0 {
		return domain.User{}, domain.ErrNotFound
	}

	return user, nil
}
func (uu *userUsecase) Update(user domain.User, id uint) (domain.User, error) {
	userDB, err := uu.userRepo.Get(id)
	if err != nil {
		return domain.User{}, err
	}
	if userDB.ID == 0 {
		return domain.User{}, domain.ErrNotFound
	}

	res, err := uu.userRepo.Update(&userDB, user)
	if err != nil {
		return domain.User{}, err
	}

	return res, nil
}
func (uu *userUsecase) Delete(id uint) (domain.User, error) {
	user, err := uu.userRepo.Get(id)
	if err != nil {
		return domain.User{}, err
	}
	if user.ID == 0 {
		return domain.User{}, domain.ErrNotFound
	}

	res, err := uu.userRepo.Delete(id)
	if err != nil {
		return domain.User{}, err
	}
	return res, nil
}
