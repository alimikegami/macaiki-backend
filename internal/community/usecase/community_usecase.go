package usecase

import (
	"macaiki/internal/domain"
	cloudstorage "macaiki/pkg/cloud_storage"
	"strconv"

	"macaiki/internal/community/dto"

	"github.com/go-playground/validator/v10"
)

type CommunityUsecaseImpl struct {
	communityRepo domain.CommunityRepository
	userRepo      domain.UserRepository
	validator     *validator.Validate
	awsS3         *cloudstorage.S3
}

func NewCommunityUsecase(communityRepo domain.CommunityRepository, userRepo domain.UserRepository, validator *validator.Validate, awsS3 *cloudstorage.S3) domain.CommunityUsecase {
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
		return []dto.CommunityResponse{}, domain.ErrInternalServerError
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
		return []dto.CommunityWithDetailResponse{}, domain.ErrInternalServerError
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
		return dto.CommunityResponse{}, domain.ErrInternalServerError
	}

	if community.ID == 0 {
		return dto.CommunityResponse{}, domain.ErrNotFound
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
		return domain.ErrUnauthorizedAccess
	}

	if err := cu.validator.Struct(community); err != nil {
		return domain.ErrBadParamInput
	}

	communityEntity := domain.Community{
		Name:                        community.Name,
		CommunityImageUrl:           community.CommunityImageUrl,
		CommunityBackgroundImageUrl: community.CommunityBackgroundImageUrl,
		Description:                 community.Description,
	}

	err := cu.communityRepo.StoreCommunity(communityEntity)
	if err != nil {
		return domain.ErrInternalServerError
	}

	return nil
}
func (cu *CommunityUsecaseImpl) UpdateCommunity(id uint, community dto.CommunityRequest, role string) error {
	if role != "Admin" {
		return domain.ErrUnauthorizedAccess
	}

	if err := cu.validator.Struct(community); err != nil {
		return domain.ErrBadParamInput
	}

	communityDB, err := cu.communityRepo.GetCommunity(id)
	if err != nil {
		return domain.ErrInternalServerError
	}

	if communityDB.ID == 0 {
		return domain.ErrNotFound
	}

	newCommunity := domain.Community{
		Name:                        community.Name,
		CommunityImageUrl:           community.CommunityImageUrl,
		CommunityBackgroundImageUrl: community.CommunityBackgroundImageUrl,
		Description:                 community.Description,
	}

	err = cu.communityRepo.UpdateCommunity(communityDB, newCommunity)
	if err != nil {
		return domain.ErrInternalServerError
	}

	return nil
}
func (cu *CommunityUsecaseImpl) DeleteCommunity(id uint, role string) error {
	if role != "Admin" {
		return domain.ErrUnauthorizedAccess
	}

	communityDB, err := cu.communityRepo.GetCommunity(id)
	if err != nil {
		return domain.ErrInternalServerError
	}

	if communityDB.ID == 0 {
		return domain.ErrNotFound
	}

	err = cu.communityRepo.DeleteCommunity(communityDB)
	if err != nil {
		return domain.ErrInternalServerError
	}

	return nil
}

func (cu *CommunityUsecaseImpl) FollowCommunity(userID, communityID uint) error {
	user, err := cu.userRepo.Get(userID)
	if err != nil {
		return domain.ErrInternalServerError
	}
	if user.ID == 0 {
		return domain.ErrNotFound
	}

	community, err := cu.communityRepo.GetCommunity(communityID)
	if err != nil {
		return domain.ErrInternalServerError
	}
	if community.ID == 0 {
		return domain.ErrNotFound
	}

	err = cu.communityRepo.FollowCommunity(user, community)
	if err != nil {
		return domain.ErrInternalServerError
	}

	return nil
}
func (cu *CommunityUsecaseImpl) UnfollowCommunity(userID, communityID uint) error {
	user, err := cu.userRepo.Get(userID)
	if err != nil {
		return domain.ErrInternalServerError
	}
	if user.ID == 0 {
		return domain.ErrNotFound
	}

	community, err := cu.communityRepo.GetCommunity(communityID)
	if err != nil {
		return domain.ErrInternalServerError
	}
	if community.ID == 0 {
		return domain.ErrNotFound
	}

	err = cu.communityRepo.UnfollowCommunity(user, community)
	if err != nil {
		return domain.ErrInternalServerError
	}

	return nil
}
