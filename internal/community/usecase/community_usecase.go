package usecase

import (
	"macaiki/internal/community"
	"macaiki/internal/user"
	cloudstorage "macaiki/pkg/cloud_storage"
	"macaiki/pkg/utils"
	"strconv"

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

func (cu *CommunityUsecaseImpl) GetAllCommunity(search string) ([]dto.CommunityResponse, error) {
	communities, err := cu.communityRepo.GetAllCommunity(search)
	if err != nil {
		return []dto.CommunityResponse{}, utils.ErrInternalServerError
	}

	communitiesResp := []dto.CommunityResponse{}
	for _, val := range communities {
		communitiesResp = append(communitiesResp, dto.CommunityResponse{
			ID:                          val.ID,
			Name:                        val.Name,
			CommunityImageUrl:           val.CommunityImageUrl,
			CommunityBackgroundImageUrl: val.CommunityBackgroundImageUrl,
			Description:                 val.Description,
		})
	}

	return communitiesResp, nil
}

func (cu *CommunityUsecaseImpl) GetAllCommunityDetail(userID int, search string) ([]dto.CommunityWithDetailResponse, error) {
	communities, err := cu.communityRepo.GetAllCommunityDetail(strconv.Itoa(int(userID)), search)
	if err != nil {
		return []dto.CommunityWithDetailResponse{}, utils.ErrInternalServerError
	}

	communitiesResp := []dto.CommunityWithDetailResponse{}
	for _, val := range communities {
		communitiesResp = append(communitiesResp, dto.CommunityWithDetailResponse{
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

func (cu *CommunityUsecaseImpl) GetCommunity(id uint) (dto.CommunityResponse, error) {
	community, err := cu.communityRepo.GetCommunity(id)
	if err != nil {
		return dto.CommunityResponse{}, utils.ErrInternalServerError
	}

	if community.ID == 0 {
		return dto.CommunityResponse{}, utils.ErrNotFound
	}

	communityResp := dto.CommunityResponse{
		ID:                          community.ID,
		Name:                        community.Name,
		CommunityImageUrl:           community.CommunityImageUrl,
		CommunityBackgroundImageUrl: community.CommunityBackgroundImageUrl,
		Description:                 community.Description,
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
		Name:                        community.Name,
		CommunityImageUrl:           community.CommunityImageUrl,
		CommunityBackgroundImageUrl: community.CommunityBackgroundImageUrl,
		Description:                 community.Description,
	}

	err := cu.communityRepo.StoreCommunity(communityEntity)
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
}
func (cu *CommunityUsecaseImpl) UpdateCommunity(id uint, community dto.CommunityRequest, role string) error {
	if role != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	if err := cu.validator.Struct(community); err != nil {
		return utils.ErrBadParamInput
	}

	communityDB, err := cu.communityRepo.GetCommunity(id)
	if err != nil {
		return utils.ErrInternalServerError
	}

	if communityDB.ID == 0 {
		return utils.ErrNotFound
	}

	newCommunity := entity.Community{
		Name:                        community.Name,
		CommunityImageUrl:           community.CommunityImageUrl,
		CommunityBackgroundImageUrl: community.CommunityBackgroundImageUrl,
		Description:                 community.Description,
	}

	err = cu.communityRepo.UpdateCommunity(communityDB, newCommunity)
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
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
