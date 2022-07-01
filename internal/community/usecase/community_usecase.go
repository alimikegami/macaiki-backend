package usecase

import (
	"macaiki/internal/community"
	"macaiki/internal/user"
	cloudstorage "macaiki/pkg/cloud_storage"
	"macaiki/pkg/utils"

	"macaiki/internal/community/dto"
	"macaiki/internal/community/entity"

	"github.com/go-playground/validator/v10"
)

type CommunityUsecaseImpl struct {
	communityRepo community.CommunityRepository
	userRepo      user.UserRepository
	validator     *validator.Validate
	awsS3         *cloudstorage.S3
}

func NewCommunityUsecase(communityRepo community.CommunityRepository, userRepo user.UserRepository, validator *validator.Validate, awsS3 *cloudstorage.S3) community.CommunityUsecase {
	return &CommunityUsecaseImpl{
		communityRepo: communityRepo,
		userRepo:      userRepo,
		validator:     validator,
		awsS3:         awsS3,
	}
}

func (cu *CommunityUsecaseImpl) GetAllCommunities(userID int, search string) ([]dto.CommunityDetailResponse, error) {
	communities, err := cu.communityRepo.GetAllCommunities(uint(userID), search)
	if err != nil {
		return []dto.CommunityDetailResponse{}, utils.ErrInternalServerError
	}

	communitiesResp := []dto.CommunityDetailResponse{}
	for _, val := range communities {
		communitiesResp = append(communitiesResp, dto.CommunityDetailResponse{
			ID:                          val.ID,
			Name:                        val.Name,
			CommunityImageUrl:           val.CommunityImageUrl,
			CommunityBackgroundImageUrl: val.CommunityBackgroundImageUrl,
			Description:                 val.Description,
			IsFollowed:                  val.IsFollowed,
		})
	}

	return communitiesResp, nil
}

func (cu *CommunityUsecaseImpl) GetCommunity(userID, communityID uint) (dto.CommunityDetailResponse, error) {
	community, err := cu.communityRepo.GetCommunityWithDetail(userID, communityID)
	if err != nil {
		return dto.CommunityDetailResponse{}, utils.ErrInternalServerError
	}

	if community.ID == 0 {
		return dto.CommunityDetailResponse{}, utils.ErrNotFound
	}

	communityResp := dto.CommunityDetailResponse{
		ID:                          community.ID,
		Name:                        community.Name,
		CommunityImageUrl:           community.CommunityImageUrl,
		CommunityBackgroundImageUrl: community.CommunityBackgroundImageUrl,
		Description:                 community.Description,
		IsFollowed:                  community.IsFollowed,
	}

	return communityResp, err
}
func (cu *CommunityUsecaseImpl) StoreCommunity(community dto.CommunityRequest, role string) error {
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
func (cu *CommunityUsecaseImpl) UpdateCommunity(id uint, community dto.CommunityRequest, role string) (dto.CommunityResponse, error) {
	if role != "Admin" {
		return dto.CommunityResponse{}, utils.ErrUnauthorizedAccess
	}

	if err := cu.validator.Struct(community); err != nil {
		return dto.CommunityResponse{}, utils.ErrBadParamInput
	}

	communityDB, err := cu.communityRepo.GetCommunity(id)
	if err != nil {
		return dto.CommunityResponse{}, utils.ErrInternalServerError
	}

	if communityDB.ID == 0 {
		return dto.CommunityResponse{}, utils.ErrNotFound
	}

	newCommunity := entity.Community{
		Name:        community.Name,
		Description: community.Description,
	}

	communityDB, err = cu.communityRepo.UpdateCommunity(communityDB, newCommunity)
	if err != nil {
		return dto.CommunityResponse{}, utils.ErrInternalServerError
	}

	communityResp := dto.CommunityResponse{
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

	err = cu.communityRepo.DeleteCommunity(communityDB)
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
