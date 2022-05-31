package usercase

import (
	"errors"
	"fmt"
	"macaiki/domain"
	"macaiki/user/delivery/http/middleware"

	"github.com/go-playground/validator/v10"
)

type userUsecase struct {
	userRepo  domain.UserRepository
	validator *validator.Validate
}

func NewUserUsecase(repo domain.UserRepository, validator *validator.Validate) domain.UserUsecase {
	return &userUsecase{
		userRepo:  repo,
		validator: validator,
	}
}

func (uu *userUsecase) Login(email, password string) (string, error) {
	if email == "" {
		return "", errors.New("Email empty")
	}
	if password == "" {
		return "", errors.New("Password empty")
	}

	user, err := uu.userRepo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	if user.ID == 0 || user.Password != password {
		return "", errors.New("Invalid email or password")
	}

	token, err := middleware.CreateToken(int(user.ID), user.Role_ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uu *userUsecase) Register(user domain.User) (domain.User, error) {

	if err := uu.validator.Struct(user); err != nil {
		return domain.User{}, domain.ErrBadParamInput
	}

	userEmail, err := uu.userRepo.GetByEmail(user.Email)
	if err != nil {
		return domain.User{}, err
	}
	if userEmail.ID != 0 {
		return domain.User{}, domain.ErrEmailAlreadyUsed
	}

	user, err = uu.userRepo.Store(user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (uu *userUsecase) GetAll() ([]domain.User, error) {
	users, err := uu.userRepo.GetAll()
	if err != nil {
		return []domain.User{}, err
	}

	return users, err
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
	fmt.Println(user)
	userDB, err := uu.userRepo.Get(id)
	if err != nil {
		return domain.User{}, err
	}
	if userDB.ID == 0 {
		return domain.User{}, domain.ErrNotFound
	}

	if userDB.Email != user.Email {
		userEmail, err := uu.userRepo.GetByEmail(user.Email)
		if err != nil {
			return domain.User{}, err
		}
		if userEmail.ID != 0 {
			return domain.User{}, domain.ErrEmailAlreadyUsed
		}
	}

	userDB, err = uu.userRepo.Update(&userDB, user)
	if err != nil {
		return domain.User{}, err
	}

	return userDB, nil
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
