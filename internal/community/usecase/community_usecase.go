package usecase

import (
	"fmt"
	community "macaiki/internal/community"
	reportCategory "macaiki/internal/report_category"
	"macaiki/internal/thread"
	user "macaiki/internal/user"
	cloudstorage "macaiki/pkg/cloud_storage"
	"macaiki/pkg/utils"
	"mime/multipart"

	dtoCommunity "macaiki/internal/community/dto"
	"macaiki/internal/community/entity"
	dtoThread "macaiki/internal/thread/dto"
	dtoUser "macaiki/internal/user/dto"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CommunityUsecaseImpl struct {
	communityRepo community.CommunityRepository
	userRepo      user.UserRepository
	rcRepo        reportCategory.ReportCategoryRepository
	threadRepo    thread.ThreadRepository
	validator     *validator.Validate
	awsS3         *cloudstorage.S3
}

func NewCommunityUsecase(communityRepo community.CommunityRepository, userRepo user.UserRepository, rcRepo reportCategory.ReportCategoryRepository, threadRepo thread.ThreadRepository, validator *validator.Validate, awsS3 *cloudstorage.S3) community.CommunityUsecase {
	return &CommunityUsecaseImpl{
		communityRepo: communityRepo,
		userRepo:      userRepo,
		threadRepo:    threadRepo,
		rcRepo:        rcRepo,
		validator:     validator,
		awsS3:         awsS3,
	}
}

func (cu *CommunityUsecaseImpl) GetAllCommunities(userID int, search string) ([]dtoCommunity.CommunityDetailResponse, error) {
	communities, err := cu.communityRepo.GetAllCommunities(uint(userID), search)
	if err != nil {
		return []dtoCommunity.CommunityDetailResponse{}, utils.ErrInternalServerError
	}

	communitiesResp := []dtoCommunity.CommunityDetailResponse{}
	for _, val := range communities {
		communitiesResp = append(communitiesResp, dtoCommunity.CommunityDetailResponse{
			ID:                          val.ID,
			Name:                        val.Name,
			CommunityImageUrl:           val.CommunityImageUrl,
			CommunityBackgroundImageUrl: val.CommunityBackgroundImageUrl,
			Description:                 val.Description,
			IsFollowed:                  val.IsFollowed,
			IsModerator:                 val.IsModerator,
		})
	}

	return communitiesResp, nil
}

func (cu *CommunityUsecaseImpl) GetCommunity(userID, communityID uint) (dtoCommunity.CommunityDetailResponse, error) {
	community, err := cu.communityRepo.GetCommunityWithDetail(userID, communityID)
	if err != nil {
		return dtoCommunity.CommunityDetailResponse{}, utils.ErrInternalServerError
	}

	if community.ID == 0 {
		return dtoCommunity.CommunityDetailResponse{}, utils.ErrNotFound
	}

	communityResp := dtoCommunity.CommunityDetailResponse{
		ID:                          community.ID,
		Name:                        community.Name,
		CommunityImageUrl:           community.CommunityImageUrl,
		CommunityBackgroundImageUrl: community.CommunityBackgroundImageUrl,
		Description:                 community.Description,
		IsFollowed:                  community.IsFollowed,
		IsModerator:                 community.IsModerator,
	}

	return communityResp, err
}

func (cu *CommunityUsecaseImpl) GetCommunityAbout(userID, communityID uint) (dtoCommunity.CommunityAboutResponse, error) {
	community, err := cu.communityRepo.GetCommunityAbout(userID, communityID)
	if err != nil {
		return dtoCommunity.CommunityAboutResponse{}, utils.ErrInternalServerError
	}

	moderators, err := cu.communityRepo.GetModeratorByCommunityID(userID, communityID)
	if err != nil {
		return dtoCommunity.CommunityAboutResponse{}, utils.ErrInternalServerError
	}

	dtoUserDetail := []dtoUser.UserResponse{}

	for _, val := range moderators {
		dtoUserDetail = append(dtoUserDetail, dtoUser.UserResponse{
			ID:              val.ID,
			Username:        val.Username,
			Name:            val.Name,
			ProfileImageUrl: val.ProfileImageUrl,
			IsFollowed:      val.IsFollowed,
			IsMine:          val.IsMine,
		})
	}

	dtoCommunity := dtoCommunity.CommunityAboutResponse{
		ID:                community.ID,
		Name:              community.Name,
		CommunityImageUrl: community.CommunityImageUrl,
		Description:       community.Description,
		IsFollowed:        community.IsFollowed,
		IsModerator:       community.IsModerator,
		TotalModerator:    community.TotalModerators,
		TotalFollower:     community.TotalFollowers,
		Moderator:         dtoUserDetail,
	}

	return dtoCommunity, nil
}

func (cu *CommunityUsecaseImpl) StoreCommunity(community dtoCommunity.CommunityRequest, role string) error {
	if role != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	if err := cu.validator.Struct(community); err != nil {
		return utils.ErrBadParamInput
	}

	communityEntity := entity.Community{
		Name:        community.Name,
		Description: community.Description,
	}

	err := cu.communityRepo.StoreCommunity(communityEntity)
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
}
func (cu *CommunityUsecaseImpl) UpdateCommunity(id uint, community dtoCommunity.CommunityRequest, role string) (dtoCommunity.CommunityUpdateResponse, error) {
	if role != "Admin" {
		return dtoCommunity.CommunityUpdateResponse{}, utils.ErrUnauthorizedAccess
	}

	if err := cu.validator.Struct(community); err != nil {
		return dtoCommunity.CommunityUpdateResponse{}, utils.ErrBadParamInput
	}

	communityDB, err := cu.communityRepo.GetCommunity(id)
	if err != nil {
		return dtoCommunity.CommunityUpdateResponse{}, utils.ErrInternalServerError
	}

	if communityDB.ID == 0 {
		return dtoCommunity.CommunityUpdateResponse{}, utils.ErrNotFound
	}

	newCommunity := entity.Community{
		Name:        community.Name,
		Description: community.Description,
	}

	communityDB, err = cu.communityRepo.UpdateCommunity(communityDB, newCommunity)
	if err != nil {
		return dtoCommunity.CommunityUpdateResponse{}, utils.ErrInternalServerError
	}

	communityResp := dtoCommunity.CommunityUpdateResponse{
		ID:          communityDB.ID,
		Name:        communityDB.Name,
		Description: communityDB.Description,
	}
	return communityResp, nil
}
func (cu *CommunityUsecaseImpl) DeleteCommunity(id uint, role string) error {
	if role != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	communityDB, err := cu.communityRepo.GetCommunity(id)
	if err != nil {
		return utils.ErrInternalServerError
	}

	if communityDB.ID == 0 {
		return utils.ErrNotFound
	}

	err = cu.communityRepo.DeleteCommunity(communityDB.ID)
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
}

func (cu *CommunityUsecaseImpl) FollowCommunity(userID, communityID uint) error {
	user, err := cu.userRepo.Get(userID)
	if err != nil {
		return utils.ErrInternalServerError
	}
	if user.ID == 0 {
		return utils.ErrNotFound
	}

	community, err := cu.communityRepo.GetCommunity(communityID)
	if err != nil {
		return utils.ErrInternalServerError
	}
	if community.ID == 0 {
		return utils.ErrNotFound
	}

	err = cu.communityRepo.FollowCommunity(user, community)
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
}
func (cu *CommunityUsecaseImpl) UnfollowCommunity(userID, communityID uint) error {
	user, err := cu.userRepo.Get(userID)
	if err != nil {
		return utils.ErrInternalServerError
	}
	if user.ID == 0 {
		return utils.ErrNotFound
	}

	community, err := cu.communityRepo.GetCommunity(communityID)
	if err != nil {
		return utils.ErrInternalServerError
	}
	if community.ID == 0 {
		return utils.ErrNotFound
	}

	err = cu.communityRepo.UnfollowCommunity(user, community)
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
}

func (cu *CommunityUsecaseImpl) SetImage(id uint, img *multipart.FileHeader, role string) (string, error) {
	if role != "Admin" {
		return "", utils.ErrUnauthorizedAccess
	}

	community, err := cu.communityRepo.GetCommunity(id)
	if err != nil {
		return "", utils.ErrInternalServerError
	}

	if community.CommunityImageUrl != "" {
		err = cu.awsS3.DeleteImage(community.CommunityImageUrl, "community")
		if err != nil {
			return "", err
		}
	}

	uniqueFilename := uuid.New()
	result, err := cu.awsS3.UploadImage(uniqueFilename.String(), "community", img)
	if err != nil {
		return "", err
	}

	imageURL := aws.StringValue(&result.Location)

	err = cu.communityRepo.SetCommunityImage(id, imageURL, "community_image_url")
	if err != nil {
		return "", err
	}

	return imageURL, err
}

func (cu *CommunityUsecaseImpl) SetBackgroundImage(id uint, img *multipart.FileHeader, role string) (string, error) {
	if role != "Admin" {
		return "", utils.ErrUnauthorizedAccess
	}

	community, err := cu.communityRepo.GetCommunity(id)
	if err != nil {
		return "", utils.ErrInternalServerError
	}

	if community.CommunityImageUrl != "" {
		err = cu.awsS3.DeleteImage(community.CommunityImageUrl, "community_background")
		if err != nil {
			return "", err
		}
	}

	uniqueFilename := uuid.New()
	result, err := cu.awsS3.UploadImage(uniqueFilename.String(), "community_background", img)
	if err != nil {
		return "", err
	}

	imageURL := aws.StringValue(&result.Location)

	err = cu.communityRepo.SetCommunityImage(id, imageURL, "community_background_image_url")
	if err != nil {
		return "", err
	}

	return imageURL, err
}

func (cu *CommunityUsecaseImpl) GetThreadCommunity(userID, communityID uint) ([]dtoThread.DetailedThreadResponse, error) {
	threadsEntity, err := cu.communityRepo.GetCommunityThread(userID, communityID)
	if err != nil {
		return []dtoThread.DetailedThreadResponse{}, utils.ErrInternalServerError
	}

	dtoThreads := []dtoThread.DetailedThreadResponse{}
	for _, val := range threadsEntity {
		dtoThread := dtoThread.DetailedThreadResponse{
			ID:                    val.Thread.ID,
			Title:                 val.Thread.Title,
			Body:                  val.Thread.Body,
			CommunityID:           val.Thread.CommunityID,
			ImageURL:              val.Thread.ImageURL,
			UpvotesCount:          val.UpvotesCount,
			IsFollowed:            val.IsFollowed,
			IsUpvoted:             val.IsUpvoted,
			UserID:                val.Thread.UserID,
			UserName:              val.User.Name,
			UserProfession:        val.User.Profession,
			UserProfilePictureURL: val.User.ProfileImageUrl,
			CreatedAt:             val.Thread.CreatedAt,
			UpdatedAt:             val.Thread.UpdatedAt,
		}
		dtoThreads = append(dtoThreads, dtoThread)
	}

	return dtoThreads, nil
}

func (cu *CommunityUsecaseImpl) AddModerator(moderatorReq dtoCommunity.CommunityModeratorRequest, role string) error {
	if role != "Admin" {
		return utils.ErrUnauthorizedAccess
	}
	if moderatorReq.UserID == 0 || moderatorReq.CommunityID == 0 {
		return utils.ErrBadParamInput
	}
	user, err := cu.userRepo.Get(moderatorReq.UserID)
	if err != nil {
		return utils.ErrInternalServerError
	}

	if user.ID == 0 {
		return utils.ErrNotFound
	}

	community, err := cu.communityRepo.GetCommunity(moderatorReq.CommunityID)
	if err != nil {
		return utils.ErrInternalServerError
	}

	if community.ID == 0 {
		return utils.ErrNotFound
	}

	err = cu.communityRepo.AddModerator(user, community)
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
}

func (cu *CommunityUsecaseImpl) RemoveModerator(moderatorReq dtoCommunity.CommunityModeratorRequest, role string) error {
	if role != "Admin" {
		return utils.ErrUnauthorizedAccess
	}
	if moderatorReq.UserID == 0 || moderatorReq.CommunityID == 0 {
		return utils.ErrBadParamInput
	}
	user, err := cu.userRepo.Get(moderatorReq.UserID)
	if err != nil {
		return utils.ErrInternalServerError
	}

	if user.ID == 0 {
		return utils.ErrNotFound
	}

	community, err := cu.communityRepo.GetCommunity(moderatorReq.CommunityID)
	if err != nil {
		return utils.ErrInternalServerError
	}

	if community.ID == 0 {
		return utils.ErrNotFound
	}

	err = cu.communityRepo.RemoveModerator(user, community)
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
}

func (cu *CommunityUsecaseImpl) ReportCommunity(userID, communityID, reportCategoryID uint) error {
	community, err := cu.communityRepo.GetCommunity(communityID)
	if err != nil {
		return utils.ErrInternalServerError
	}

	if community.ID == 0 {
		return utils.ErrNotFound
	}

	reportCategory, err := cu.rcRepo.GetReportCategory(reportCategoryID)
	if err != nil {
		return utils.ErrInternalServerError
	}

	if reportCategory.ID == 0 {
		return utils.ErrNotFound
	}

	err = cu.communityRepo.StoreReportCommunity(entity.CommunityReport{
		UserID:              userID,
		CommunityReportedID: communityID,
		ReportCategoryID:    reportCategoryID,
	})
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
}

func (cu *CommunityUsecaseImpl) ReportByModerator(userID, communityID uint, reportReq dtoCommunity.ReportRequest) error {
	mods, err := cu.communityRepo.GetModeratorByUserID(userID, communityID)

	if err != nil {
		return err
	}
	fmt.Println(mods)
	fmt.Println(mods.CommunityID)
	if mods.CommunityID != communityID {
		return utils.ErrUnauthorizedAccess
	}

	if reportReq.CommentReportID != 0 {
		commentReport, err := cu.threadRepo.GetCommentReport(reportReq.CommentReportID)
		if err != nil {
			return utils.ErrInternalServerError
		}

		err = cu.threadRepo.UpdateCommentReport(commentReport, userID)
		if err != nil {
			return utils.ErrInternalServerError
		}
	} else if reportReq.CommunityReportID != 0 {
		communityReport, err := cu.communityRepo.GetReportCommunity(reportReq.CommunityReportID)
		if err != nil {
			return utils.ErrInternalServerError
		}

		err = cu.communityRepo.UpdateReportCommunity(communityReport, userID)
		if err != nil {
			return utils.ErrInternalServerError
		}
	} else {
		threadReport, err := cu.threadRepo.GetThreadReport(reportReq.ThreadReportID)
		if err != nil {
			return utils.ErrInternalServerError
		}

		err = cu.threadRepo.UpdateThreadReport(threadReport, userID)
		if err != nil {
			return utils.ErrInternalServerError
		}
	}

	return nil
}

func (cu *CommunityUsecaseImpl) GetReports(userID, communityID uint) ([]dtoCommunity.BriefReportResponse, error) {
	mods, err := cu.communityRepo.GetModeratorByUserID(userID, communityID)

	if err != nil {
		return []dtoCommunity.BriefReportResponse{}, err
	}

	if mods.CommunityID != communityID {
		return []dtoCommunity.BriefReportResponse{}, utils.ErrUnauthorizedAccess
	}

	reports, err := cu.communityRepo.GetReports(communityID)

	if err != nil {
		return []dtoCommunity.BriefReportResponse{}, utils.ErrInternalServerError
	}

	var reportsResp []dtoCommunity.BriefReportResponse

	for _, report := range reports {
		reportsResp = append(reportsResp, dtoCommunity.BriefReportResponse{
			ThreadReportsID:     report.ThreadReportsID,
			CommunityReportsID:  report.CommunityReportsID,
			CommentReportsID:    report.CommentReportsID,
			CreatedAt:           report.CreatedAt,
			ThreadID:            report.ThreadID,
			CommunityReportedID: report.CommunityReportedID,
			UserID:              report.UserID,
			CommentID:           report.CommentID,
			ReportCategory:      report.ReportCategory,
			Username:            report.Username,
			ProfileImageURL:     report.ProfileImageURL,
			Type:                report.Type,
		})
	}

	return reportsResp, nil
}
