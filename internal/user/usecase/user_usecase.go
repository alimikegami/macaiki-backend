package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	comRepo "macaiki/internal/community"
	"macaiki/internal/notification"
	notificationEntity "macaiki/internal/notification/entity"
	reportcategory "macaiki/internal/report_category"
	"macaiki/internal/thread"
	dtoThread "macaiki/internal/thread/dto"
	"macaiki/internal/user"
	"macaiki/internal/user/delivery/http/helper"
	"macaiki/internal/user/dto"
	"macaiki/internal/user/entity"
	cloudstorage "macaiki/pkg/cloud_storage"
	goMail "macaiki/pkg/gomail"
	"macaiki/pkg/middleware"
	"macaiki/pkg/utils"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo           user.UserRepository
	reportCategoryRepo reportcategory.ReportCategoryRepository
	communityRepo      comRepo.CommunityRepository
	notificationRepo   notification.NotificationRepository
	threadRepo         thread.ThreadRepository
	validator          *validator.Validate
	awsS3              *cloudstorage.S3
	goMail             *goMail.Gomail
	rdb                *redis.Client
}

var (
	DEFAULT_PROFILE    = "https://macaiki.s3.ap-southeast-3.amazonaws.com/profile/default-avatar.png"
	DEFAULT_BACKGROUND = "https://macaiki.s3.ap-southeast-3.amazonaws.com/background/default-background.png"
)

func NewUserUsecase(userRepo user.UserRepository, reportCategoryRepo reportcategory.ReportCategoryRepository, communityRepo comRepo.CommunityRepository, notificationRepo notification.NotificationRepository, threadRepo thread.ThreadRepository, validator *validator.Validate, awsS3Instace *cloudstorage.S3, goMail *goMail.Gomail, rdb *redis.Client) user.UserUsecase {
	return &userUsecase{
		userRepo:           userRepo,
		reportCategoryRepo: reportCategoryRepo,
		communityRepo:      communityRepo,
		notificationRepo:   notificationRepo,
		threadRepo:         threadRepo,
		validator:          validator,
		awsS3:              awsS3Instace,
		goMail:             goMail,
		rdb:                rdb,
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
		Email:              user.Email,
		Username:           user.Username,
		Password:           hashAndSalt([]byte(user.Password)),
		Role:               "User",
		ProfileImageUrl:    DEFAULT_PROFILE,
		BackgroundImageUrl: DEFAULT_BACKGROUND,
		Name:               user.Username,
		IsBanned:           0,
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
	var userEntity entity.User
	var err error
	value, err := uu.rdb.Get(context.Background(), fmt.Sprintf("user:%d", tokenUserID)).Result()
	if err == redis.Nil {
		userEntity, err = uu.userRepo.GetWithDetail(id, tokenUserID)
		if err != nil {
			return dto.UserDetailResponse{}, utils.ErrInternalServerError
		}
		stringifiedThreads, err := json.Marshal(userEntity)
		if err != nil {
			log.Println(err)
			return dto.UserDetailResponse{}, errors.New("Marshaling Errors")
		}

		if userEntity.ID == 0 {
			return dto.UserDetailResponse{}, utils.ErrNotFound
		}
		err = uu.rdb.Set(context.Background(), fmt.Sprintf("user:%d", tokenUserID), stringifiedThreads, 0).Err()
		if err != nil {
			log.Println(err)
			return dto.UserDetailResponse{}, errors.New("Errors setting redis keys")
		}
	} else if err != nil {
		log.Println(err)
		return dto.UserDetailResponse{}, err
	} else {
		err = json.Unmarshal([]byte(value), &userEntity)
		if err != nil {
			log.Println(err)
			return dto.UserDetailResponse{}, errors.New("Unmarshaling Errors")
		}
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
	// validation the user exist
	statusCmd := uu.rdb.Del(context.Background(), fmt.Sprintf("user:%d", id))
	if statusCmd.Err() != nil {
		log.Println(statusCmd.Err())
		return dto.UserUpdateResponse{}, utils.ErrInternalServerError
	}

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

func (uu *userUsecase) GetUserFollowers(tokenUserID, getFollowingUserID uint) ([]dto.UserResponse, error) {
	userEntity, err := uu.userRepo.Get(getFollowingUserID)
	if err != nil {
		return []dto.UserResponse{}, utils.ErrInternalServerError
	}
	if userEntity.ID == 0 {
		return []dto.UserResponse{}, utils.ErrNotFound
	}

	followers, err := uu.userRepo.GetFollower(tokenUserID, getFollowingUserID)
	if err != nil {
		return []dto.UserResponse{}, utils.ErrInternalServerError
	}
	return helper.DomainUserToListUserResponse(followers), nil
}

func (uu *userUsecase) GetUserFollowing(tokenUserID, getFollowingUserID uint) ([]dto.UserResponse, error) {
	userEntity, err := uu.userRepo.Get(getFollowingUserID)
	if err != nil {
		return []dto.UserResponse{}, utils.ErrInternalServerError
	}
	if userEntity.ID == 0 {
		return []dto.UserResponse{}, utils.ErrNotFound
	}

	following, err := uu.userRepo.GetFollowing(tokenUserID, getFollowingUserID)
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

	if user.ID == 0 {
		return "", utils.ErrNotFound
	}

	if user.ProfileImageUrl != DEFAULT_PROFILE {
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

	if user.ID == 0 {
		return "", utils.ErrNotFound
	}

	if user.BackgroundImageUrl != DEFAULT_BACKGROUND {
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
		UserID:            userID,
		NotificationType:  "Follow You",
		NotificationRefID: userFollowerID,
		IsReaded:          0,
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
	user, err := uu.userRepo.Get(userReportedID)
	if err != nil {
		return utils.ErrInternalServerError
	}

	if user.ID == 0 {
		return utils.ErrNotFound
	}

	reportCategory, err := uu.reportCategoryRepo.GetReportCategory(reportCategoryID)
	if err != nil {
		return utils.ErrInternalServerError
	}
	if reportCategory.ID == 0 {
		return utils.ErrNotFound
	}

	err = uu.userRepo.StoreReport(entity.UserReport{
		UserID:           userID,
		ReportedUserID:   userReportedID,
		ReportCategoryID: reportCategoryID,
	})
	if err != nil {
		return utils.ErrInternalServerError
	}
	return nil
}

func (uu *userUsecase) GetThreadByToken(userID, tokenUserID uint) ([]dtoThread.DetailedThreadResponse, error) {
	dtoThreads := []dtoThread.DetailedThreadResponse{}

	threads, err := uu.threadRepo.GetThreadsByUserID(userID, tokenUserID)
	if err != nil {
		log.Println(err)
		return []dtoThread.DetailedThreadResponse{}, utils.ErrInternalServerError
	}

	for _, val := range threads {
		dtoThreads = append(dtoThreads, dtoThread.DetailedThreadResponse{
			ID:                    val.Thread.ID,
			Title:                 val.Title,
			Body:                  val.Body,
			CommunityID:           val.CommunityID,
			ImageURL:              val.ImageURL,
			UserID:                val.UserID,
			UserName:              val.Username,
			UserProfession:        val.Profession,
			UserProfilePictureURL: val.ProfileImageUrl,
			CreatedAt:             val.Thread.CreatedAt,
			UpdatedAt:             val.Thread.UpdatedAt,
			UpvotesCount:          val.UpvotesCount,
			IsUpvoted:             val.IsUpvoted,
			IsDownVoted:           val.IsDownvoted,
			IsFollowed:            val.IsFollowed,
		})
	}

	return dtoThreads, nil
}

func (uu *userUsecase) SendOTP(otpReq dto.SendOTPRequest) error {
	user, err := uu.userRepo.GetByEmail(otpReq.Email)
	if err != nil {
		return utils.ErrInternalServerError
	}

	if user.ID == 0 {
		return utils.ErrNotFound
	}

	OTPCode := uu.goMail.GenerateSecureToken(3)
	err = uu.userRepo.StoreOTP(entity.VerificationEmail{
		Email:     user.Email,
		OTPCode:   OTPCode,
		ExpiredAt: time.Now().Add(1 * time.Minute),
	})
	if err != nil {
		return utils.ErrInternalServerError
	}

	err = uu.goMail.SendMail(user.Email, user.Username, fmt.Sprintf("Thank you for registering on the Macaiki application To verify your email, please use the following OTP : %s", OTPCode))
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func (uu *userUsecase) VerifyOTP(email, OTPCode string) error {
	EmailVerif, err := uu.userRepo.GetOTP(email)
	if err != nil {
		return utils.ErrInternalServerError
	}

	if EmailVerif.OTPCode == OTPCode {
		if time.Now().After(EmailVerif.ExpiredAt) {
			return errors.New("OTP Is Expired")
		}
		user, err := uu.userRepo.GetByEmail(email)
		if err != nil {
			return utils.ErrInternalServerError
		}
		_, _ = uu.userRepo.Update(&user, entity.User{
			EmailVerifiedAt: time.Now(),
		})
	} else {
		return errors.New("OTP Not Valid")
	}

	return nil
}

func (uu *userUsecase) GetReports(curentUserRole string) ([]dto.BriefReportResponse, error) {
	if curentUserRole != "Admin" {
		return []dto.BriefReportResponse{}, utils.ErrUnauthorizedAccess
	}

	reports, err := uu.userRepo.GetReports()

	if err != nil {
		return []dto.BriefReportResponse{}, utils.ErrInternalServerError
	}

	var reportsResp []dto.BriefReportResponse

	for _, report := range reports {
		reportsResp = append(reportsResp, dto.BriefReportResponse{
			ThreadReportsID:     report.ThreadReportsID,
			UserReportsID:       report.UserReportsID,
			CommentReportsID:    report.CommentReportsID,
			CommunityReportsID:  report.CommunityReportsID,
			CreatedAt:           report.CreatedAt,
			ThreadID:            report.ThreadID,
			UserID:              report.UserID,
			CommentID:           report.CommentID,
			CommunityReportedIT: report.CommunityReportedID,
			ReportCategory:      report.ReportCategory,
			Username:            report.Username,
			ProfileImageURL:     report.ProfileImageURL,
			Type:                report.Type,
		})
	}

	return reportsResp, nil
}

func (uu *userUsecase) GetDashboardAnalytics(userRole string) (dto.AdminDashboardAnalytics, error) {
	if userRole != "Admin" {
		return dto.AdminDashboardAnalytics{}, utils.ErrUnauthorizedAccess
	}

	analytics, err := uu.userRepo.GetDashboardAnalytics()

	if err != nil {
		return dto.AdminDashboardAnalytics{}, utils.ErrInternalServerError
	}

	return dto.AdminDashboardAnalytics{
		UsersCount:      analytics.UsersCount,
		ModeratorsCount: analytics.ModeratorsCount,
		ReportsCount:    analytics.ReportsCount,
	}, nil
}

func (uu *userUsecase) GetReportedThread(userRole string, threadReportID uint) (dto.ReportedThreadResponse, error) {
	var reportedThreadResponse dto.ReportedThreadResponse

	if userRole != "Admin" {
		return dto.ReportedThreadResponse{}, utils.ErrUnauthorizedAccess
	}

	reportedThread, err := uu.userRepo.GetReportedThread(threadReportID)

	if err != nil {
		return dto.ReportedThreadResponse{}, err
	}

	reportedThreadResponse = dto.ReportedThreadResponse{
		ID:                      reportedThread.ID,
		ThreadTitle:             reportedThread.ThreadTitle,
		ThreadBody:              reportedThread.ThreadBody,
		ThreadImageURL:          reportedThread.ThreadImageURL,
		ThreadCreatedAt:         reportedThread.ThreadCreatedAt,
		LikesCount:              reportedThread.LikesCount,
		ReportedUsername:        reportedThread.ReportedUsername,
		ReportedProfileImageURL: reportedThread.ReportedProfileImageURL,
		ReportedUserProfession:  reportedThread.ReportedUserProfession,
		ReportCategory:          reportedThread.ReportCategory,
		ReportCreatedAt:         reportedThread.ReportCreatedAt,
		Username:                reportedThread.Username,
		ProfileImageURL:         reportedThread.ProfileImageURL,
	}

	return reportedThreadResponse, nil
}

func (uu *userUsecase) GetReportedCommunity(userRole string, communityReportID uint) (dto.ReportedCommunityResponse, error) {
	var reportedCommunityResponse dto.ReportedCommunityResponse

	if userRole != "Admin" {
		return dto.ReportedCommunityResponse{}, utils.ErrUnauthorizedAccess
	}

	reportedCommunity, err := uu.userRepo.GetReportedCommunity(communityReportID)

	if err != nil {
		return dto.ReportedCommunityResponse{}, err
	}

	reportedCommunityResponse = dto.ReportedCommunityResponse{
		ID:                          reportedCommunity.ID,
		CommunityName:               reportedCommunity.CommunityName,
		CommunityImageURL:           reportedCommunity.CommunityImageURL,
		CommunityBackgroundImageURL: reportedCommunity.CommunityBackgroundImageURL,
		ReportCategory:              reportedCommunity.ReportCategory,
		ReportCreatedAt:             reportedCommunity.ReportCreatedAt,
		Username:                    reportedCommunity.Username,
		ProfileImageURL:             reportedCommunity.ProfileImageURL,
	}

	return reportedCommunityResponse, nil
}

func (uu *userUsecase) GetReportedComment(userRole string, commentReportID uint) (dto.ReportedCommentResponse, error) {
	var reportedCommentResponse dto.ReportedCommentResponse

	if userRole != "Admin" {
		return dto.ReportedCommentResponse{}, utils.ErrUnauthorizedAccess
	}

	reportedComment, err := uu.userRepo.GetReportedComment(commentReportID)

	if err != nil {
		return dto.ReportedCommentResponse{}, err
	}

	reportedCommentResponse = dto.ReportedCommentResponse{
		ID:                      reportedComment.ID,
		CommentBody:             reportedComment.CommentBody,
		LikesCount:              reportedComment.LikesCount,
		CommentCreatedAt:        reportedComment.CommentCreatedAt,
		ReportedUsername:        reportedComment.ReportedUsername,
		ReportedProfileImageURL: reportedComment.ReportedProfileImageURL,
		ReportCategory:          reportedComment.ReportCategory,
		ReportCreatedAt:         reportedComment.ReportCreatedAt,
		Username:                reportedComment.Username,
		ProfileImageURL:         reportedComment.ProfileImageURL,
	}

	return reportedCommentResponse, nil
}

func (uu *userUsecase) GetReportedUser(userRole string, userReportID uint) (dto.ReportedUserResponse, error) {
	var reportedUserResponse dto.ReportedUserResponse

	if userRole != "Admin" {
		return dto.ReportedUserResponse{}, utils.ErrUnauthorizedAccess
	}

	reportedUser, err := uu.userRepo.GetReportedUser(userReportID)

	if err != nil {
		return dto.ReportedUserResponse{}, err
	}

	reportedUserResponse = dto.ReportedUserResponse{
		ID:                          reportedUser.ID,
		ReportedUserUsername:        reportedUser.ReportedUserUsername,
		ReportedUserName:            reportedUser.ReportedUserName,
		ReportedUserProfession:      reportedUser.ReportedUserProfession,
		ReporteduserBio:             reportedUser.ReporteduserBio,
		ReportedUserProfileImageURL: reportedUser.ReportedUserProfileImageURL,
		ReportedUserBackgroundURL:   reportedUser.ReportedUserBackgroundURL,
		ReportingUserUsername:       reportedUser.ReportingUserUsername,
		ReportingUserName:           reportedUser.ReportedUserName,
		FollowersCount:              reportedUser.FollowersCount,
		FollowingCount:              reportedUser.FollowingCount,
	}

	return reportedUserResponse, nil
}

func (uu *userUsecase) BanUser(userRole string, userReportID uint) error {
	if userRole != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	report, err := uu.userRepo.GetUserReport(userReportID)

	if err != nil {
		return err
	}

	err = uu.userRepo.DeleteUserReport(userReportID)

	if err != nil {
		return err
	}

	err = uu.userRepo.Delete(report.ReportedUserID)

	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) BanThread(userRole string, threadReportID uint) error {
	if userRole != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	report, err := uu.threadRepo.GetThreadReport(threadReportID)

	if err != nil {
		return err
	}

	err = uu.userRepo.DeleteThreadReport(threadReportID)

	if err != nil {
		return err
	}

	err = uu.threadRepo.DeleteThread(report.ThreadID)

	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) BanComment(userRole string, commentReportID uint) error {
	if userRole != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	report, err := uu.threadRepo.GetCommentReport(commentReportID)

	if err != nil {
		return err
	}

	err = uu.userRepo.DeleteCommentReport(commentReportID)

	if err != nil {
		return err
	}

	err = uu.threadRepo.DeleteComment(report.CommentID)

	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) BanCommunity(userRole string, communityReportID uint) error {
	if userRole != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	report, err := uu.communityRepo.GetReportCommunity(communityReportID)

	if err != nil {
		return err
	}

	err = uu.userRepo.DeleteCommunityReport(communityReportID)

	if err != nil {
		return err
	}

	err = uu.communityRepo.DeleteCommunity(report.CommunityReportedID)

	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) DeleteThreadReport(userRole string, threadReportID uint) error {
	if userRole != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	err := uu.userRepo.DeleteThreadReport(threadReportID)

	return err
}

func (uu *userUsecase) DeleteUserReport(userRole string, userReportID uint) error {
	if userRole != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	err := uu.userRepo.DeleteUserReport(userReportID)

	return err
}

func (uu *userUsecase) DeleteCommentReport(userRole string, commentReportID uint) error {
	if userRole != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	err := uu.userRepo.DeleteCommentReport(commentReportID)

	return err
}

func (uu *userUsecase) DeleteCommunityReport(userRole string, communityReportID uint) error {
	if userRole != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	err := uu.userRepo.DeleteCommunityReport(communityReportID)

	return err
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
