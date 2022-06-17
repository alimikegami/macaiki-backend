package usercase

import (
	"fmt"
	"log"
	"macaiki/internal/domain"
	"macaiki/internal/user/delivery/http/helper"
	"macaiki/internal/user/dto"
	cloudstorage "macaiki/pkg/cloud_storage"
	"mime/multipart"
	"macaiki/pkg/middleware"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo  domain.UserRepository
	validator *validator.Validate
	awsS3     *cloudstorage.S3
}

func NewUserUsecase(repo domain.UserRepository, validator *validator.Validate, awsS3Instace *cloudstorage.S3) domain.UserUsecase {
	return &userUsecase{
		userRepo:  repo,
		validator: validator,
		awsS3:     awsS3Instace,
	}
}

func (uu *userUsecase) Login(email, password string) (dto.LoginResponse, error) {
	if email == "" {
		return dto.LoginResponse{}, domain.ErrBadParamInput
	}
	if password == "" {
		return dto.LoginResponse{}, domain.ErrBadParamInput
	}

	userEntity, err := uu.userRepo.GetByEmail(email)
	if err != nil {
		return dto.LoginResponse{}, domain.ErrInternalServerError
	}

	if userEntity.ID == 0 || !comparePasswords(userEntity.Password, []byte(password)) {
		return dto.LoginResponse{}, domain.ErrLoginFailed
	}

	token, err := middleware.JWTCreateToken(int(userEntity.ID), userEntity.Role)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return helper.ToLoginResponse(token), nil
}

func (uu *userUsecase) Register(user dto.UserRequest) error {
	// TO DO : error handling for existing username
	if err := uu.validator.Struct(user); err != nil {
		return domain.ErrBadParamInput
	}

	userEmail, err := uu.userRepo.GetByEmail(user.Email)
	if err != nil {
		return domain.ErrInternalServerError
	}

	if userEmail.ID != 0 {
		return domain.ErrEmailAlreadyUsed
	}

	if user.Password != user.PasswordConfirmation {
		return domain.ErrPasswordDontMatch
	}

	userEntity := domain.User{
		Email:    user.Email,
		Username: user.Username,
		Password: hashAndSalt([]byte(user.Password)),
		Role:     "User",
		Name:     user.Username,
		IsBanned: false,
	}

	err = uu.userRepo.Store(userEntity)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) GetAll() ([]dto.UserResponse, error) {
	users, err := uu.userRepo.GetAll()
	if err != nil {
		return []dto.UserResponse{}, domain.ErrInternalServerError
	}

	return helper.DomainUserToListUserResponse(users), err
}

func (uu *userUsecase) Get(id uint) (dto.UserDetailResponse, error) {
	userEntity, err := uu.userRepo.Get(id)
	if err != nil {
		return dto.UserDetailResponse{}, domain.ErrInternalServerError
	}
	if userEntity.ID == 0 {
		return dto.UserDetailResponse{}, domain.ErrNotFound
	}

	totalFollowing, err := uu.userRepo.GetFollowingNumber(id)
	if err != nil {
		return dto.UserDetailResponse{}, domain.ErrInternalServerError
	}

	totalFollower, err := uu.userRepo.GetFollowerNumber(id)
	if err != nil {
		return dto.UserDetailResponse{}, domain.ErrInternalServerError
	}

	if err != nil {
		return dto.UserDetailResponse{}, domain.ErrInternalServerError
	}
	return helper.DomainUserToUserDetailResponse(userEntity, totalFollowing, totalFollower), nil
}
func (uu *userUsecase) Update(user dto.UpdateUserRequest, id uint) (dto.UserResponse, error) {
	if err := uu.validator.Struct(user); err != nil {
		return dto.UserResponse{}, domain.ErrBadParamInput
	}

	userDB, err := uu.userRepo.Get(id)
	if err != nil {
		return dto.UserResponse{}, domain.ErrInternalServerError
	}
	if userDB.ID == 0 {
		return dto.UserResponse{}, domain.ErrNotFound
	}

	if userDB.Email != user.Email {
		userEmail, err := uu.userRepo.GetByEmail(user.Email)
		if err != nil {
			return dto.UserResponse{}, domain.ErrInternalServerError
		}
		if userEmail.ID != 0 {
			return dto.UserResponse{}, domain.ErrEmailAlreadyUsed
		}
	}

	userEntity := domain.User{
		Email:              user.Email,
		Username:           user.Username,
		Name:               user.Name,
		ProfileImageUrl:    user.ProfileImageUrl,
		BackgroundImageUrl: user.BackgroundImageUrl,
		Bio:                user.Bio,
		Proffesion:         user.Proffesion,
		Role:               user.Role,
		IsBanned:           user.IsBanned,
	}
	userDB, err = uu.userRepo.Update(&userDB, userEntity)
	if err != nil {
		return dto.UserResponse{}, domain.ErrInternalServerError
	}

	return helper.DomainUserToUserResponse(userDB), nil
}
func (uu *userUsecase) Delete(id uint) error {
	userEntity, err := uu.userRepo.Get(id)
	if err != nil {
		return domain.ErrInternalServerError
	}
	if userEntity.ID == 0 {
		return domain.ErrNotFound
	}

	_, err = uu.userRepo.Delete(id)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return nil
}

func (uu *userUsecase) GetUserFollowers(id uint) ([]dto.UserResponse, error) {
	userEntity, err := uu.userRepo.Get(id)
	if err != nil {
		return []dto.UserResponse{}, domain.ErrInternalServerError
	}
	if userEntity.ID == 0 {
		return []dto.UserResponse{}, domain.ErrNotFound
	}

	following, err := uu.userRepo.GetFollower(userEntity)
	if err != nil {
		return []dto.UserResponse{}, domain.ErrInternalServerError
	}
	return helper.DomainUserToListUserResponse(following), nil
}

func (uu *userUsecase) GetUserFollowing(id uint) ([]dto.UserResponse, error) {
	userEntity, err := uu.userRepo.Get(id)
	if err != nil {
		return []dto.UserResponse{}, domain.ErrInternalServerError
	}
	if userEntity.ID == 0 {
		return []dto.UserResponse{}, domain.ErrNotFound
	}

	following, err := uu.userRepo.GetFollowing(userEntity)
	if err != nil {
		return []dto.UserResponse{}, domain.ErrInternalServerError
	}
	return helper.DomainUserToListUserResponse(following), nil
}

func (uu *userUsecase) SetProfileImage(id uint, img *multipart.FileHeader) (string, error) {
	uniqueFilename := uuid.New()
	result, err := uu.awsS3.UploadImage(uniqueFilename.String(), "profile", img)
	if err != nil {
		fmt.Printf("failed to upload file, %v", err)
		return "", err
	}

	imageURL := aws.StringValue(&result.Location)
	fmt.Printf("file uploaded to, %s\n", imageURL)

	err = uu.userRepo.SetUserImage(id, imageURL, "profile_image_url")
	if err != nil {
		fmt.Println("failed to save url on database")
		return "", err
	}

	return imageURL, err
}

func (uu *userUsecase) SetBackgroundImage(id uint, img *multipart.FileHeader) (string, error) {
	uniqueFilename := uuid.New()
	result, err := uu.awsS3.UploadImage(uniqueFilename.String(), "background", img)
	if err != nil {
		fmt.Printf("failed to upload file, %v", err)
		return "", err
	}

	imageURL := aws.StringValue(&result.Location)
	fmt.Printf("file uploaded to, %s\n", imageURL)
	// fmt.Printf("file uploaded to, %s\n", uniqueFilename.String()+filepath.Ext(img.Filename))
	err = uu.userRepo.SetUserImage(id, imageURL, "background_image_url")
	if err != nil {
		fmt.Println("failed to save url on database")
		return "", err
	}

	return imageURL, err
}

func (uu *userUsecase) Follow(user_id, user_follower_id uint) error {
	user, err := uu.userRepo.Get(user_id)
	if err != nil {
		return domain.ErrInternalServerError
	}
	if user.ID == 0 {
		return domain.ErrNotFound
	}

	user_follower, err := uu.userRepo.Get(user_follower_id)
	if err != nil {
		return domain.ErrInternalServerError
	}
	if user_follower.ID == 0 {
		return domain.ErrNotFound
	}

	// if follow self account throw error bad param input
	if user.ID == user_follower.ID {
		return domain.ErrBadParamInput
	}

	// save to database
	_, err = uu.userRepo.Follow(user, user_follower)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return nil
}

func (uu *userUsecase) Unfollow(user_id, user_follower_id uint) error {
	user, err := uu.userRepo.Get(user_id)
	if err != nil {
		return domain.ErrInternalServerError
	}
	if user.ID == 0 {
		return domain.ErrNotFound
	}

	user_follower, err := uu.userRepo.Get(user_follower_id)
	if err != nil {
		return domain.ErrInternalServerError
	}
	if user_follower.ID == 0 {
		return domain.ErrNotFound
	}

	_, err = uu.userRepo.Unfollow(user, user_follower)
	if err != nil {
		return domain.ErrInternalServerError
	}
	return nil
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
