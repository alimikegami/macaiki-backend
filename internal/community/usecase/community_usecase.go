package usecase

import (
	"macaiki/internal/community"
	cloudstorage "macaiki/pkg/cloud_storage"
	"macaiki/pkg/utils"

	"macaiki/internal/community/dto"
	"macaiki/internal/community/entity"

	"github.com/go-playground/validator/v10"
)

type CommunityUsecaseImpl struct {
	communityRepo community.CommunityRepository
	validator     *validator.Validate
	awsS3         *cloudstorage.S3
}

func NewCommunityUsecase(communityRepo community.CommunityRepository, validator *validator.Validate, awsS3 *cloudstorage.S3) community.CommunityUsecase {
	return &CommunityUsecaseImpl{communityRepo: communityRepo, validator: validator, awsS3: awsS3}
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
