package usecase

import (
	"macaiki/internal/user/entity"
	mockUser "macaiki/internal/user/mocks"
	"macaiki/pkg/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	mockedUserEntity = entity.User{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email:              "dummy@gmail.com",
		Username:           "dummy",
		Password:           "123456",
		Name:               "dummy",
		ProfileImageUrl:    "",
		BackgroundImageUrl: "",
		Bio:                "",
		Profession:         "",
		Role:               "User",
		// EmailVerifiedAt:    time.Now(),
		IsBanned: 0,
	}

	mockedUserArr = []entity.User{mockedUserEntity, mockedUserEntity, mockedUserEntity}

	mockedUserReportEntity = entity.UserReport{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserID:           1,
		ReportedUserID:   1,
		ReportCategoryID: 1,
	}

	// v = validator.New()
)

// func TestRegister(t *testing.T) {
// 	mockUserRepo := userMock.NewUserRepository(t)
// 	mockNotifRepo := notifMock.NewNotificationRepository(t)
// 	mockThreadRepo := threadMock.NewThreadRepository(t)
// 	mockRcRepo := rcMock.NewReportCategoryRepository(t)

// 	mockUserReq := dto.UserRequest{
// 		Email:                "dummy@gmail.com",
// 		Username:             "dummy",
// 		Password:             "123456",
// 		PasswordConfirmation: "123456",
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockUserRepo.On("GetByEmail", mockUserReq.Email).Return(entity.User{}, nil).Once()
// 		mockUserRepo.On("GetByUsername", mockUserReq.Username).Return(entity.User{}, nil).Once()
// 		mockUserRepo.On("Store", mockedUserEntity).Return(nil).Once()

// 		testUserUsecase := NewUserUsecase(mockUserRepo, mockRcRepo, mockNotifRepo, mockThreadRepo, v, nil, nil)
// 		err := testUserUsecase.Register(mockUserReq)

// 		assert.NoError(t, err)
// 	})
// }

// func TestLogin(t *testing.T) {
// 	mockUserRepo := userMock.NewUserRepository(t)
// 	mockNotifRepo := notifMock.NewNotificationRepository(t)
// 	mockThreadRepo := threadMock.NewThreadRepository(t)
// 	mockRcRepo := rcMock.NewReportCategoryRepository(t)

// 	mockUserRepo.On("GetByEmail", mockedUserEntity.Email).Return(entity.User{}, nil).Once()

// 	testUserUsecase := NewUserUsecase(mockUserRepo, mockRcRepo, mockNotifRepo, mockThreadRepo, v, nil, nil)
// }

func TestGetAll(t *testing.T) {
	mockUserRepo := mockUser.NewUserRepository(t)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetAllWithDetail", uint(1), "").Return(mockedUserArr, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetAll(uint(1), "")

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("GetAllWithDetail", uint(1), "").Return(mockedUserArr, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetAll(uint(1), "")

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestGet(t *testing.T) {
	mockUserRepo := mockUser.NewUserRepository(t)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockedUserEntity, nil).Once()
		mockUserRepo.On("GetFollowingNumber", uint(1)).Return(10, nil).Once()
		mockUserRepo.On("GetFollowerNumber", uint(1)).Return(10, nil).Once()
		mockUserRepo.On("GetThreadsNumber", uint(1)).Return(10, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.Get(uint(1), uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("user-not-found", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(entity.User{}, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.Get(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("internal-server-error-1", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockedUserEntity, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.Get(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("internal-server-error-2", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockedUserEntity, nil).Once()
		mockUserRepo.On("GetFollowingNumber", uint(1)).Return(0, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.Get(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("internal-server-error-3", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockedUserEntity, nil).Once()
		mockUserRepo.On("GetFollowingNumber", uint(1)).Return(10, nil).Once()
		mockUserRepo.On("GetFollowerNumber", uint(1)).Return(10, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.Get(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("internal-server-error-4", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockedUserEntity, nil).Once()
		mockUserRepo.On("GetFollowingNumber", uint(1)).Return(10, nil).Once()
		mockUserRepo.On("GetFollowerNumber", uint(1)).Return(10, nil).Once()
		mockUserRepo.On("GetThreadsNumber", uint(1)).Return(10, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.Get(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}
