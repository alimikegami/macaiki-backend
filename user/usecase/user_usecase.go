package usercase

import (
	"errors"
	"log"
	"macaiki/domain"
	"macaiki/user/delivery/http/middleware"
	"macaiki/user/delivery/http/request"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
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
		return "", domain.ErrBadParamInput
	}
	if password == "" {
		return "", domain.ErrBadParamInput
	}

	user, err := uu.userRepo.GetByEmail(email)
	if err != nil {
		return "", domain.ErrInternalServerError
	}

	if user.ID == 0 || !comparePasswords(user.Password, []byte(password)) {
		return "", domain.ErrLoginFailed
	}

	token, err := middleware.JWTCreateToken(int(user.ID), user.Role)
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
		return domain.User{}, domain.ErrInternalServerError
	}
	if userEmail.ID != 0 {
		return domain.User{}, domain.ErrEmailAlreadyUsed
	}

	user.Password = hashAndSalt([]byte(user.Password))
	user, err = uu.userRepo.Store(user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (uu *userUsecase) GetAll() ([]domain.User, error) {
	users, err := uu.userRepo.GetAll()
	if err != nil {
		return []domain.User{}, domain.ErrInternalServerError
	}

	return users, err
}

func (uu *userUsecase) Get(id uint) (domain.User, []domain.User, error) {
	user, err := uu.userRepo.Get(id)
	if err != nil {
		return domain.User{}, []domain.User{}, domain.ErrInternalServerError
	}
	if user.ID == 0 {
		return domain.User{}, []domain.User{}, domain.ErrNotFound
	}

	followings, err := uu.userRepo.GetFollowing(user)
	if err != nil {
		return domain.User{}, []domain.User{}, domain.ErrInternalServerError
	}
	return user, followings, nil
}
func (uu *userUsecase) Update(user domain.User, id uint) (domain.User, error) {
	userUpdate := request.ToUserUpdateRequest(user)
	if err := uu.validator.Struct(userUpdate); err != nil {
		return domain.User{}, domain.ErrBadParamInput
	}
	if len(userUpdate.Password) != 0 && len(userUpdate.Password) < 6 {
		return domain.User{}, errors.New("password at least 6 characters")
	}

	userDB, err := uu.userRepo.Get(id)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}
	if userDB.ID == 0 {
		return domain.User{}, domain.ErrNotFound
	}

	if userDB.Email != user.Email {
		userEmail, err := uu.userRepo.GetByEmail(user.Email)
		if err != nil {
			return domain.User{}, domain.ErrInternalServerError
		}
		if userEmail.ID != 0 {
			return domain.User{}, domain.ErrEmailAlreadyUsed
		}
	}
	userDB, err = uu.userRepo.Update(&userDB, user)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}

	return userDB, nil
}
func (uu *userUsecase) Delete(id uint) (domain.User, error) {
	user, err := uu.userRepo.Get(id)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}
	if user.ID == 0 {
		return domain.User{}, domain.ErrNotFound
	}

	res, err := uu.userRepo.Delete(id)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}
	return res, nil
}

func (uu *userUsecase) Follow(user_id, user_follower_id uint) (domain.User, error) {
	user, err := uu.userRepo.Get(user_id)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}
	if user.ID == 0 {
		return domain.User{}, domain.ErrNotFound
	}

	user_follower, err := uu.userRepo.Get(user_follower_id)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}
	if user_follower.ID == 0 {
		return domain.User{}, domain.ErrNotFound
	}

	// if follow self account throw error bad param input
	if user.ID == user_follower.ID {
		return domain.User{}, domain.ErrBadParamInput
	}

	// save to database
	res, err := uu.userRepo.Follow(user, user_follower)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}
	return res, nil
}

func (uu *userUsecase) Unfollow(user_id, user_follower_id uint) (domain.User, error) {
	user, err := uu.userRepo.Get(user_id)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}
	if user.ID == 0 {
		return domain.User{}, domain.ErrNotFound
	}

	user_follower, err := uu.userRepo.Get(user_follower_id)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}
	if user_follower.ID == 0 {
		return domain.User{}, domain.ErrNotFound
	}

	res, err := uu.userRepo.Unfollow(user, user_follower)
	if err != nil {
		return domain.User{}, domain.ErrInternalServerError
	}
	return res, nil
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println("err", err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println("err", err)
		return false
	}

	return true
}
