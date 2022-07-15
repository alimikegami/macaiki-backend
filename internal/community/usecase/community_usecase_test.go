package usecase

import (
	communityDTO "macaiki/internal/community/dto"
	communityEntity "macaiki/internal/community/entity"
	communityMock "macaiki/internal/community/mocks"
	userEntity "macaiki/internal/user/entity"
	"macaiki/pkg/utils"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	mockCommunityEntity = communityEntity.Community{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:                        "dummy",
		CommunityImageUrl:           "dummy",
		CommunityBackgroundImageUrl: "dummy",
		Description:                 "dummy",
	}

	mockCommunityEntityArr = []communityEntity.Community{mockCommunityEntity}

	mockUserEntity = userEntity.User{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email:              "dummy@gmail.com",
		Username:           "dummy",
		Password:           "$2a$04$UP.ZNuepVAiEedwlZrvA3.ywqqNszceSuqnZQl4mozYOzO9ILY2kK",
		Name:               "dummy",
		ProfileImageUrl:    "dummy",
		BackgroundImageUrl: "dummy",
		Bio:                "dummy",
		Profession:         "dummy",
		Role:               "User",
		EmailVerifiedAt:    time.Now(),
		IsBanned:           0,
	}

	mockUserEntityArr = []userEntity.User{mockUserEntity}

	v = validator.New()
)

func TestGetAllCommunities(t *testing.T) {
	mockCommunityRepo := communityMock.NewCommunityRepository(t)

	t.Run("success", func(t *testing.T) {
		mockCommunityRepo.On("GetAllCommunities", uint(1), "").Return(mockCommunityEntityArr, nil).Once()

		testCommunityUsecase := NewCommunityUsecase(mockCommunityRepo, nil, nil, nil, nil, nil)
		res, err := testCommunityUsecase.GetAllCommunities(1, "")

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockCommunityRepo.On("GetAllCommunities", uint(1), "").Return([]communityEntity.Community{}, utils.ErrInternalServerError).Once()

		testCommunityUsecase := NewCommunityUsecase(mockCommunityRepo, nil, nil, nil, nil, nil)
		res, err := testCommunityUsecase.GetAllCommunities(1, "")

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestGetCommunity(t *testing.T) {
	mockCommunityRepo := communityMock.NewCommunityRepository(t)

	t.Run("success", func(t *testing.T) {
		mockCommunityRepo.On("GetCommunityWithDetail", uint(1), uint(1)).Return(mockCommunityEntity, nil).Once()

		testCommunityUsecase := NewCommunityUsecase(mockCommunityRepo, nil, nil, nil, nil, nil)
		res, err := testCommunityUsecase.GetCommunity(uint(1), uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockCommunityRepo.On("GetCommunityWithDetail", uint(1), uint(1)).Return(communityEntity.Community{}, utils.ErrInternalServerError).Once()

		testCommunityUsecase := NewCommunityUsecase(mockCommunityRepo, nil, nil, nil, nil, nil)
		res, err := testCommunityUsecase.GetCommunity(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("community-not-found", func(t *testing.T) {
		mockCommunityRepo.On("GetCommunityWithDetail", uint(1), uint(1)).Return(communityEntity.Community{}, nil).Once()

		testCommunityUsecase := NewCommunityUsecase(mockCommunityRepo, nil, nil, nil, nil, nil)
		res, err := testCommunityUsecase.GetCommunity(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestGetCommunityAbout(t *testing.T) {
	mockCommunityRepo := communityMock.NewCommunityRepository(t)

	t.Run("success", func(t *testing.T) {
		mockCommunityRepo.On("GetCommunityAbout", uint(1), uint(1)).Return(mockCommunityEntity, nil).Once()
		mockCommunityRepo.On("GetModeratorByCommunityID", uint(1), uint(1)).Return(mockUserEntityArr, nil).Once()

		testCommunityUsecase := NewCommunityUsecase(mockCommunityRepo, nil, nil, nil, nil, nil)
		res, err := testCommunityUsecase.GetCommunityAbout(uint(1), uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockCommunityRepo.On("GetCommunityAbout", uint(1), uint(1)).Return(communityEntity.Community{}, utils.ErrInternalServerError).Once()

		testCommunityUsecase := NewCommunityUsecase(mockCommunityRepo, nil, nil, nil, nil, nil)
		res, err := testCommunityUsecase.GetCommunityAbout(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("success", func(t *testing.T) {
		mockCommunityRepo.On("GetCommunityAbout", uint(1), uint(1)).Return(mockCommunityEntity, nil).Once()
		mockCommunityRepo.On("GetModeratorByCommunityID", uint(1), uint(1)).Return([]userEntity.User{}, utils.ErrInternalServerError).Once()

		testCommunityUsecase := NewCommunityUsecase(mockCommunityRepo, nil, nil, nil, nil, nil)
		res, err := testCommunityUsecase.GetCommunityAbout(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestStoreCommunity(t *testing.T) {
	mockCommunityRepo := communityMock.NewCommunityRepository(t)

	mockCommunityEntityReq := communityEntity.Community{
		Name:        "dummy",
		Description: "dummy",
	}

	mockCommunityDTOReq := communityDTO.CommunityRequest{
		Name:        "dummy",
		Description: "dummy",
	}

	mockCommunityDTOReqFail := communityDTO.CommunityRequest{
		Name:        "",
		Description: "",
	}

	t.Run("success", func(t *testing.T) {
		mockCommunityRepo.On("GetCommunity", uint(1)).Return(mockCommunityEntity, nil).Once()
		mockCommunityRepo.On("UpdateCommunity", mockCommunityEntity, mockCommunityEntityReq).Return(mockCommunityEntity, nil)

		testCommunityUsecase := NewCommunityUsecase(mockCommunityRepo, nil, nil, nil, v, nil)
		res, err := testCommunityUsecase.UpdateCommunity(uint(1), mockCommunityDTOReq, "Admin")

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	// t.Run("internal-server-error", func(t *testing.T) {
	// 	mockCommunityRepo.On("GetCommunity", uint(1)).Return(communityEntity.Community{}, utils.ErrInternalServerError).Once()

	// 	testCommunityUsecase := NewCommunityUsecase(mockCommunityRepo, nil, nil, nil, v, nil)
	// 	res, err := testCommunityUsecase.UpdateCommunity(uint(1), mockCommunityDTOReq, "Admin")

	// 	assert.Error(t, err)
	// 	assert.Empty(t, res)
	// })

	// t.Run("internal-server-error", func(t *testing.T) {
	// 	mockCommunityRepo.On("GetCommunity", uint(1)).Return(mockCommunityEntity, nil).Once()
	// 	mockCommunityRepo.On("UpdateCommunity", mockCommunityEntity, mockCommunityEntityReq).Return(communityEntity.Community{}, utils.ErrInternalServerError)

	// 	testCommunityUsecase := NewCommunityUsecase(mockCommunityRepo, nil, nil, nil, v, nil)
	// 	res, err := testCommunityUsecase.UpdateCommunity(uint(1), mockCommunityDTOReq, "Admin")

	// 	assert.Error(t, err)
	// 	assert.Empty(t, res)
	// })

	t.Run("bad-param-input", func(t *testing.T) {
		testCommunityUsecase := NewCommunityUsecase(mockCommunityRepo, nil, nil, nil, v, nil)
		res, err := testCommunityUsecase.UpdateCommunity(uint(1), mockCommunityDTOReq, "User")

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("community-not-found", func(t *testing.T) {
		mockCommunityRepo.On("GetCommunity", uint(1)).Return(communityEntity.Community{}, nil).Once()

		testCommunityUsecase := NewCommunityUsecase(mockCommunityRepo, nil, nil, nil, v, nil)
		res, err := testCommunityUsecase.UpdateCommunity(uint(1), mockCommunityDTOReq, "Admin")

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("unauthorize", func(t *testing.T) {
		testCommunityUsecase := NewCommunityUsecase(mockCommunityRepo, nil, nil, nil, v, nil)
		res, err := testCommunityUsecase.UpdateCommunity(uint(1), mockCommunityDTOReqFail, "Admin")

		assert.Error(t, err)
		assert.Empty(t, res)
	})

}
