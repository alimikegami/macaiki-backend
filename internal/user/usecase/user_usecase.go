package usercase

import (
	"fmt"
	"log"
	"macaiki/internal/notification"
	notificationEntity "macaiki/internal/notification/entity"
	reportcategory "macaiki/internal/report_category"
	"macaiki/internal/user"
	"macaiki/internal/user/delivery/http/helper"
	"macaiki/internal/user/dto"
	"macaiki/internal/user/entity"
	cloudstorage "macaiki/pkg/cloud_storage"
	"macaiki/pkg/middleware"
	"macaiki/pkg/utils"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo           user.UserRepository
	reportCategoryRepo reportcategory.ReportCategoryRepository
	notificationRepo   notification.NotificationRepository
	validator          *validator.Validate
	awsS3              *cloudstorage.S3
}

func NewUserUsecase(userRepo user.UserRepository, reportCategoryRepo reportcategory.ReportCategoryRepository, notificationRepo notification.NotificationRepository, validator *validator.Validate, awsS3Instace *cloudstorage.S3) user.UserUsecase {
	return &userUsecase{
		userRepo:           userRepo,
		reportCategoryRepo: reportCategoryRepo,
		notificationRepo:   notificationRepo,
		validator:          validator,
		awsS3:              awsS3Instace,
	}
}

func (uu *userUsecase) Login(loginInfo dto.UserLoginRequest) (dto.LoginResponse, error) {
	if err := uu.validator.Struct(loginInfo); err != nil {
		return dto.LoginResponse{}, utils.ErrBadParamInput
	}

	userEntity, err := uu.userRepo.GetByEmail(loginInfo.Email)
	if err != nil {
		return dto.LoginResponse{}, utils.ErrInternalServerError
	}

	if userEntity.ID == 0 || !comparePasswords(userEntity.Password, []byte(loginInfo.Password)) {
		return dto.LoginResponse{}, utils.ErrLoginFailed
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
		return utils.ErrBadParamInput
	}

	userEmail, err := uu.userRepo.GetByEmail(user.Email)
	if err != nil {
		return utils.ErrInternalServerError
	}

	if userEmail.ID != 0 {
		return utils.ErrEmailAlreadyUsed
	}

	userUsername, err := uu.userRepo.GetByUsername(user.Username)
	if err != nil {
		return utils.ErrInternalServerError
	}

	if userUsername.ID != 0 {
		return utils.ErrUsernameAlreadyUsed
	}

	if user.Password != user.PasswordConfirmation {
		return utils.ErrPasswordDontMatch
	}

	userEntity := entity.User{
		Email:    user.Email,
		Username: user.Username,
		Password: hashAndSalt([]byte(user.Password)),
		Role:     "User",
		Name:     user.Username,
		IsBanned: 0,
	}

	err = uu.userRepo.Store(userEntity)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) GetAll(userID uint, search string) ([]dto.UserResponse, error) {
	users, err := uu.userRepo.GetAllWithDetail(userID, search)
	if err != nil {
		return []dto.UserResponse{}, utils.ErrInternalServerError
	}

	return helper.DomainUserToListUserResponse(users), err
}

func (uu *userUsecase) Get(id, tokenUserID uint) (dto.UserDetailResponse, error) {
	userEntity, err := uu.userRepo.Get(id)
	if err != nil {
		return dto.UserDetailResponse{}, utils.ErrInternalServerError
	}
	if userEntity.ID == 0 {
		return dto.UserDetailResponse{}, utils.ErrNotFound
	}

	totalFollowing, err := uu.userRepo.GetFollowingNumber(id)
	if err != nil {
		return dto.UserDetailResponse{}, utils.ErrInternalServerError
	}

	totalFollower, err := uu.userRepo.GetFollowerNumber(id)
	if err != nil {
		return dto.UserDetailResponse{}, utils.ErrInternalServerError
	}

	totalPost, err := uu.userRepo.GetThreadsNumber(id)
	if err != nil {
		return dto.UserDetailResponse{}, utils.ErrInternalServerError
	}

	userResp := helper.DomainUserToUserDetailResponse(userEntity, totalFollowing, totalFollower, totalPost)
	return userResp, nil
}
func (uu *userUsecase) Update(user dto.UserUpdateRequest, id uint) (dto.UserUpdateResponse, error) {
	if err := uu.validator.Struct(user); err != nil {
		return dto.UserUpdateResponse{}, utils.ErrBadParamInput
	}

	// validation the user exist
	userDB, err := uu.userRepo.Get(id)
	if err != nil {
		return dto.UserUpdateResponse{}, utils.ErrInternalServerError
	}
	if userDB.ID == 0 {
		return dto.UserUpdateResponse{}, utils.ErrNotFound
	}

	userEntity := entity.User{
		Name:       user.Name,
		Bio:        user.Bio,
		Profession: user.Profession,
	}

	userDB, err = uu.userRepo.Update(&userDB, userEntity)
	if err != nil {
		return dto.UserUpdateResponse{}, utils.ErrInternalServerError
	}

	return helper.DomainUserToUserUpdateResponse(userDB), nil
}
func (uu *userUsecase) Delete(id uint, curentUserID uint, curentUserRole string) error {
	// validation the user exist
	userEntity, err := uu.userRepo.Get(id)
	if err != nil {
		return utils.ErrInternalServerError
	}
	if userEntity.ID == 0 {
		return utils.ErrNotFound
	}

	// validation that accesses is the user itself or Admin
	if curentUserID != id && curentUserRole != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	err = uu.userRepo.Delete(id)
	if err != nil {
		return utils.ErrInternalServerError
	}
	return nil

}

func (uu *userUsecase) ChangeEmail(id uint, info dto.UserLoginRequest) (string, error) {
	if err := uu.validator.Struct(info); err != nil {
		return "", utils.ErrBadParamInput
	}

	userDB, err := uu.userRepo.Get(id)
	if err != nil {
		return "", utils.ErrInternalServerError
	}

	if userDB.Email != info.Email {
		userEmail, err := uu.userRepo.GetByEmail(info.Email)
		if err != nil {
			return "", utils.ErrInternalServerError
		}
		if userEmail.ID != 0 {
			return "", utils.ErrEmailAlreadyUsed
		}
	}

	if !comparePasswords(userDB.Password, []byte(info.Password)) {
		return "", utils.ErrForbidden
	}

	userEntity := entity.User{
		Email: info.Email,
	}
	userDB, err = uu.userRepo.Update(&userDB, userEntity)
	if err != nil {
		return "", utils.ErrInternalServerError
	}

	return info.Email, nil
}
func (uu *userUsecase) ChangePassword(id uint, passwordInfo dto.UserChangePasswordRequest) error {
	if err := uu.validator.Struct(passwordInfo); err != nil {
		return utils.ErrBadParamInput
	}

	if passwordInfo.NewPassword != passwordInfo.PasswordConfirmation {
		return utils.ErrPasswordDontMatch
	}

	userDB, err := uu.userRepo.Get(id)
	if err != nil {
		return utils.ErrInternalServerError
	}

	userEntity := entity.User{
		Password: hashAndSalt([]byte(passwordInfo.NewPassword)),
	}

	_, err = uu.userRepo.Update(&userDB, userEntity)
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
}

func (uu *userUsecase) GetUserFollowers(id uint) ([]dto.UserResponse, error) {
	userEntity, err := uu.userRepo.Get(id)
	if err != nil {
		return []dto.UserResponse{}, utils.ErrInternalServerError
	}
	if userEntity.ID == 0 {
		return []dto.UserResponse{}, utils.ErrNotFound
	}

	followers, err := uu.userRepo.GetFollower(userEntity)
	if err != nil {
		return []dto.UserResponse{}, utils.ErrInternalServerError
	}
	return helper.DomainUserToListUserResponse(followers), nil
}

func (uu *userUsecase) GetUserFollowing(id uint) ([]dto.UserResponse, error) {
	userEntity, err := uu.userRepo.Get(id)
	if err != nil {
		return []dto.UserResponse{}, utils.ErrInternalServerError
	}
	if userEntity.ID == 0 {
		return []dto.UserResponse{}, utils.ErrNotFound
	}

	following, err := uu.userRepo.GetFollowing(userEntity)
	if err != nil {
		return []dto.UserResponse{}, utils.ErrInternalServerError
	}
	return helper.DomainUserToListUserResponse(following), nil
}

func (uu *userUsecase) SetProfileImage(id uint, img *multipart.FileHeader) (string, error) {
	user, err := uu.userRepo.Get(id)
	if err != nil {
		return "", utils.ErrInternalServerError
	}

	if user.ProfileImageUrl != "" {
		err = uu.awsS3.DeleteImage(user.ProfileImageUrl, "profile")
		if err != nil {
			return "", err
		}
	}

	uniqueFilename := uuid.New()
	result, err := uu.awsS3.UploadImage(uniqueFilename.String(), "profile", img)
	if err != nil {
		return "", err
	}

	imageURL := aws.StringValue(&result.Location)

	err = uu.userRepo.SetUserImage(id, imageURL, "profile_image_url")
	if err != nil {
		return "", err
	}

	return imageURL, err
}

func (uu *userUsecase) SetBackgroundImage(id uint, img *multipart.FileHeader) (string, error) {
	user, err := uu.userRepo.Get(id)
	if err != nil {
		return "", utils.ErrInternalServerError
	}

	if user.BackgroundImageUrl != "" {
		err = uu.awsS3.DeleteImage(user.BackgroundImageUrl, "background")
		if err != nil {
			return "", err
		}
	}

	uniqueFilename := uuid.New()
	result, err := uu.awsS3.UploadImage(uniqueFilename.String(), "background", img)
	if err != nil {
		return "", err
	}

	imageURL := aws.StringValue(&result.Location)
	err = uu.userRepo.SetUserImage(id, imageURL, "background_image_url")
	if err != nil {
		return "", err
	}

	return imageURL, err
}

func (uu *userUsecase) Follow(userID, userFollowerID uint) error {
	user, err := uu.userRepo.Get(userID)
	if err != nil {
		return utils.ErrInternalServerError
	}
	if user.ID == 0 {
		return utils.ErrNotFound
	}

	user_follower, err := uu.userRepo.Get(userFollowerID)
	if err != nil {
		return utils.ErrInternalServerError
	}
	if user_follower.ID == 0 {
		return utils.ErrNotFound
	}

	// if follow self account throw error bad param input
	if user.ID == user_follower.ID {
		return utils.ErrBadParamInput
	}

	// save to database
	_, err = uu.userRepo.Follow(user, user_follower)
	if err != nil {
		return utils.ErrInternalServerError
	}
	err = uu.notificationRepo.StoreNotification(notificationEntity.Notification{
		UserID:             userID,
		NotificationType:   "Follow You",
		NotificationTypeID: userFollowerID,
		IsReaded:           0,
	})
	if err != nil {
		fmt.Println("failed to send notification")
	}

	return nil
}

func (uu *userUsecase) Unfollow(userID, userFollowerID uint) error {
	user, err := uu.userRepo.Get(userID)
	if err != nil {
		return utils.ErrInternalServerError
	}
	if user.ID == 0 {
		return utils.ErrNotFound
	}

	user_follower, err := uu.userRepo.Get(userFollowerID)
	if err != nil {
		return utils.ErrInternalServerError
	}
	if user_follower.ID == 0 {
		return utils.ErrNotFound
	}

	_, err = uu.userRepo.Unfollow(user, user_follower)
	if err != nil {
		return utils.ErrInternalServerError
	}
	return nil
}

func (uu *userUsecase) Report(userID, userReportedID, reportCategoryID uint) error {
	var err error

	if userID == userReportedID {
		return utils.ErrBadParamInput
	}

	_, err = uu.userRepo.Get(userID)
	if err != nil {
		return utils.ErrNotFound
	}

	_, err = uu.userRepo.Get(userReportedID)
	if err != nil {
		return utils.ErrNotFound
	}

	_, err = uu.reportCategoryRepo.GetReportCategory(reportCategoryID)
	if err != nil {
		return utils.ErrNotFound
	}

	userReport := entity.UserReport{
		UserID:           userID,
		ReportedUserID:   userReportedID,
		ReportCategoryID: reportCategoryID,
	}

	err = uu.userRepo.StoreReport(userReport)
	if err != nil {
		return utils.ErrInternalServerError
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
