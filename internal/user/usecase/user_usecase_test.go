package usecase

import (
	notifEntity "macaiki/internal/notification/entity"
	notifMock "macaiki/internal/notification/mocks"
	rcEntity "macaiki/internal/report_category/entity"
	reportCategoryMock "macaiki/internal/report_category/mocks"
	threadEntity "macaiki/internal/thread/entity"
	threadMock "macaiki/internal/thread/mocks"
	userDTO "macaiki/internal/user/dto"
	userEntity "macaiki/internal/user/entity"
	userMock "macaiki/internal/user/mocks"
	"macaiki/pkg/utils"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	mockUserEntity1 = userEntity.User{
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

	mockUserEntity2 = userEntity.User{
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email:              "dummy@gmail.com",
		Username:           "dummy",
		Password:           "123456",
		Name:               "dummy",
		ProfileImageUrl:    "dummy",
		BackgroundImageUrl: "dummy",
		Bio:                "dummy",
		Profession:         "dummy",
		Role:               "Admin",
		EmailVerifiedAt:    time.Now(),
		IsBanned:           0,
	}

	mockedUserArr = []userEntity.User{mockUserEntity1, mockUserEntity1, mockUserEntity1}

	v = validator.New()
)

// func TestLogin(t *testing.T) {
// 	mockUserRepo := userMock.NewUserRepository(t)

// 	t.Run("success", func(t *testing.T) {
// 		mockUserRepo.On("GetByEmail", mockUserEntity1.Email).Return(userEntity.User{}, nil).Once()

// 		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, v, nil, nil)
// 	})
// }

// func TestRegister(t *testing.T) {
// 	mockUserRepo := userMock.NewUserRepository(t)

// 	mockUserReq := userDTO.UserRequest{
// 		Email:                "dummy@gmail.com",
// 		Username:             "dummy",
// 		Password:             "123456",
// 		PasswordConfirmation: "123456",
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockUserRepo.On("GetByEmail", mockUserReq.Email).Return(userEntity.User{}, nil).Once()
// 		mockUserRepo.On("GetByUsername", mockUserReq.Username).Return(userEntity.User{}, nil).Once()
// 		mockUserRepo.On("Store", mockUserEntity1).Return(nil).Once()

// 		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, v, nil, nil)
// 		err := testUserUsecase.Register(mockUserReq)

// 		assert.NoError(t, err)
// 	})
// }

func TestGetAll(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetAllWithDetail", uint(1), "").Return(mockedUserArr, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetAll(uint(1), "")

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("GetAllWithDetail", uint(1), "").Return(mockedUserArr, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetAll(uint(1), "")

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestGet(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetWithDetail", uint(1), uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("GetFollowingNumber", uint(1)).Return(10, nil).Once()
		mockUserRepo.On("GetFollowerNumber", uint(1)).Return(10, nil).Once()
		mockUserRepo.On("GetThreadsNumber", uint(1)).Return(10, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.Get(uint(1), uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("user-not-found", func(t *testing.T) {
		mockUserRepo.On("GetWithDetail", uint(1), uint(1)).Return(userEntity.User{}, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.Get(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("internal-server-error-1", func(t *testing.T) {
		mockUserRepo.On("GetWithDetail", uint(1), uint(1)).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.Get(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("internal-server-error-2", func(t *testing.T) {
		mockUserRepo.On("GetWithDetail", uint(1), uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("GetFollowingNumber", uint(1)).Return(0, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.Get(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("internal-server-error-3", func(t *testing.T) {
		mockUserRepo.On("GetWithDetail", uint(1), uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("GetFollowingNumber", uint(1)).Return(10, nil).Once()
		mockUserRepo.On("GetFollowerNumber", uint(1)).Return(10, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.Get(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("internal-server-error-4", func(t *testing.T) {
		mockUserRepo.On("GetWithDetail", uint(1), uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("GetFollowingNumber", uint(1)).Return(10, nil).Once()
		mockUserRepo.On("GetFollowerNumber", uint(1)).Return(10, nil).Once()
		mockUserRepo.On("GetThreadsNumber", uint(1)).Return(10, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.Get(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestUpdate(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)

	mockUserUpdateDTO := userDTO.UserUpdateRequest{
		Name:       "dummy",
		Bio:        "dummy",
		Profession: "dummy",
	}

	mockUserEntityUpdate := userEntity.User{
		Name:       "dummy",
		Bio:        "dummy",
		Profession: "dummy",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Update", &mockUserEntity1, mockUserEntityUpdate).Return(mockUserEntity1, nil).Once()

		testUseUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		res, err := testUseUsecase.Update(mockUserUpdateDTO, uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("user-not-found", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(userEntity.User{}, nil).Once()

		testUseUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		res, err := testUseUsecase.Update(mockUserUpdateDTO, uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUseUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		res, err := testUseUsecase.Update(mockUserUpdateDTO, uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Update", &mockUserEntity1, mockUserEntityUpdate).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUseUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		res, err := testUseUsecase.Update(mockUserUpdateDTO, uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestDelete(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Delete", uint(1)).Return(nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Delete(uint(1), uint(1), "Admin")

		assert.NoError(t, err)
	})

	t.Run("user-not-found", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(userEntity.User{}, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Delete(uint(1), uint(1), "Admin")

		assert.Error(t, err)
	})

	t.Run("unautorize", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Delete(uint(1), uint(2), "User")

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Delete", uint(1)).Return(utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Delete(uint(1), uint(1), "Admin")

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Delete(uint(1), uint(1), "Admin")

		assert.Error(t, err)
	})
}

func TestChangeEmail(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)

	mockInfoDTOReqSuccess := userDTO.UserLoginRequest{
		Email:    "dummyupdate@gmail.com",
		Password: "123456",
	}

	mockInfoDTOReqSuccessFail1 := userDTO.UserLoginRequest{
		Email:    "",
		Password: "123456",
	}

	mockInfoDTOReqSuccessFail2 := userDTO.UserLoginRequest{
		Email:    "dummyupdate@gmail.com",
		Password: "1234567",
	}

	mockEntityReq := userEntity.User{
		Email: "dummyupdate@gmail.com",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("GetByEmail", mockInfoDTOReqSuccess.Email).Return(userEntity.User{}, nil).Once()
		mockUserRepo.On("Update", &mockUserEntity1, mockEntityReq).Return(mockUserEntity1, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, v, nil, nil)

		res, err := testUserUsecase.ChangeEmail(uint(1), mockInfoDTOReqSuccess)

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, v, nil, nil)

		res, err := testUserUsecase.ChangeEmail(uint(1), mockInfoDTOReqSuccess)

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("GetByEmail", mockInfoDTOReqSuccess.Email).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, v, nil, nil)

		res, err := testUserUsecase.ChangeEmail(uint(1), mockInfoDTOReqSuccess)

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("GetByEmail", mockInfoDTOReqSuccess.Email).Return(userEntity.User{}, nil).Once()
		mockUserRepo.On("Update", &mockUserEntity1, mockEntityReq).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, v, nil, nil)

		res, err := testUserUsecase.ChangeEmail(uint(1), mockInfoDTOReqSuccess)

		assert.Error(t, err)
		assert.Empty(t, res)
	})
	t.Run("bad-param-input", func(t *testing.T) {
		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, v, nil, nil)

		res, err := testUserUsecase.ChangeEmail(uint(1), mockInfoDTOReqSuccessFail1)

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("email-already-used", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("GetByEmail", mockInfoDTOReqSuccess.Email).Return(mockUserEntity1, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, v, nil, nil)

		res, err := testUserUsecase.ChangeEmail(uint(1), mockInfoDTOReqSuccess)

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("forbidden", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("GetByEmail", mockInfoDTOReqSuccess.Email).Return(userEntity.User{}, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, v, nil, nil)

		res, err := testUserUsecase.ChangeEmail(uint(1), mockInfoDTOReqSuccessFail2)

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

// func TestChangePassword(t *testing.T) {
// 	mockUserRepo := userMock.NewUserRepository(t)

// 	mockUserUpdateEntity := userEntity.User{
// 		Password: "1234567",
// 	}

// 	mockPasswordInfoDTOReqSuccess := userDTO.UserChangePasswordRequest{
// 		NewPassword:          "1234567",
// 		PasswordConfirmation: "1234567",
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
// 		mockUserRepo.On("Update", &mockUserEntity1, mockUserUpdateEntity).Return(mockUserEntity1, nil).Once()

// 		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, v, nil, nil)

// 		err := testUserUsecase.ChangePassword(uint(1), mockPasswordInfoDTOReqSuccess)

// 		assert.NoError(t, err)
// 	})

// }

func TestGetUserFollowers(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("GetFollower", uint(1), uint(1)).Return(mockedUserArr, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		_, err := testUserUsecase.GetUserFollowers(uint(1), uint(1))

		assert.NoError(t, err)
	})

	t.Run("user-not-found", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(userEntity.User{}, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		_, err := testUserUsecase.GetUserFollowers(uint(1), uint(1))

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		_, err := testUserUsecase.GetUserFollowers(uint(1), uint(1))

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("GetFollower", uint(1), uint(1)).Return([]userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		_, err := testUserUsecase.GetUserFollowers(uint(1), uint(1))

		assert.Error(t, err)
	})
}

func TestGetUserFollowing(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("GetFollowing", uint(1), uint(1)).Return(mockedUserArr, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		_, err := testUserUsecase.GetUserFollowing(uint(1), uint(1))

		assert.NoError(t, err)
	})

	t.Run("user-not-found", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(userEntity.User{}, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		_, err := testUserUsecase.GetUserFollowing(uint(1), uint(1))

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		_, err := testUserUsecase.GetUserFollowing(uint(1), uint(1))

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("GetFollowing", uint(1), uint(1)).Return([]userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		_, err := testUserUsecase.GetUserFollowing(uint(1), uint(1))

		assert.Error(t, err)
	})
}

// func TestSetProfileImage(t *testing.T) {
// 	mockUserRepo := userMock.UserRepository(t)
// 	mockAwsS3 := cloudstorage.CreateNewS3Instance()
// 	t.Run("success", func(t *testing.T) {
// 		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
// 	})
// }

func TestSetBackgroundImage(t *testing.T) {}

func TestFollow(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)
	mockNotifRepo := notifMock.NewNotificationRepository(t)

	mockNotifEntity := notifEntity.Notification{
		UserID:            1,
		NotificationType:  "Follow You",
		NotificationRefID: 2,
		IsReaded:          0,
	}
	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Get", uint(2)).Return(mockUserEntity2, nil).Once()
		mockUserRepo.On("Follow", mockUserEntity1, mockUserEntity2).Return(mockUserEntity1, nil).Once()
		mockNotifRepo.On("StoreNotification", mockNotifEntity).Return(nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, mockNotifRepo, nil, nil, nil, nil)

		err := testUserUsecase.Follow(uint(1), uint(2))

		assert.NoError(t, err)
	})

	t.Run("user-not-found", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(userEntity.User{}, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, mockNotifRepo, nil, nil, nil, nil)

		err := testUserUsecase.Follow(uint(1), uint(2))

		assert.Error(t, err)
	})

	t.Run("user-not-found", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Get", uint(2)).Return(userEntity.User{}, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, mockNotifRepo, nil, nil, nil, nil)

		err := testUserUsecase.Follow(uint(1), uint(2))

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, mockNotifRepo, nil, nil, nil, nil)

		err := testUserUsecase.Follow(uint(1), uint(2))

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Get", uint(2)).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, mockNotifRepo, nil, nil, nil, nil)

		err := testUserUsecase.Follow(uint(1), uint(2))

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Get", uint(2)).Return(mockUserEntity2, nil).Once()
		mockUserRepo.On("Follow", mockUserEntity1, mockUserEntity2).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, mockNotifRepo, nil, nil, nil, nil)

		err := testUserUsecase.Follow(uint(1), uint(2))

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Get", uint(2)).Return(mockUserEntity2, nil).Once()
		mockUserRepo.On("Follow", mockUserEntity1, mockUserEntity2).Return(mockUserEntity1, nil).Once()
		mockNotifRepo.On("StoreNotification", mockNotifEntity).Return(utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, mockNotifRepo, nil, nil, nil, nil)

		err := testUserUsecase.Follow(uint(1), uint(2))

		assert.NoError(t, err)
	})

	t.Run("bad-param-input", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, mockNotifRepo, nil, nil, nil, nil)

		err := testUserUsecase.Follow(uint(1), uint(1))

		assert.Error(t, err)
	})
}

func TestUnfollow(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Get", uint(2)).Return(mockUserEntity2, nil).Once()
		mockUserRepo.On("Unfollow", mockUserEntity1, mockUserEntity2).Return(mockUserEntity1, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Unfollow(uint(1), uint(2))

		assert.NoError(t, err)
	})

	t.Run("user-not-found", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(userEntity.User{}, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Unfollow(uint(1), uint(2))

		assert.Error(t, err)
	})

	t.Run("user-not-found", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Get", uint(2)).Return(userEntity.User{}, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Unfollow(uint(1), uint(2))

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Unfollow(uint(1), uint(2))

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Get", uint(2)).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Unfollow(uint(1), uint(2))

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Get", uint(2)).Return(mockUserEntity2, nil).Once()
		mockUserRepo.On("Unfollow", mockUserEntity1, mockUserEntity2).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Unfollow(uint(1), uint(2))

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockUserRepo.On("Get", uint(2)).Return(mockUserEntity2, nil).Once()
		mockUserRepo.On("Unfollow", mockUserEntity1, mockUserEntity2).Return(mockUserEntity1, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Unfollow(uint(1), uint(2))

		assert.NoError(t, err)
	})
}

func TestReport(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)
	mockReportCategoryRepo := reportCategoryMock.NewReportCategoryRepository(t)

	mockRcEntity := rcEntity.ReportCategory{
		ID:   uint(1),
		Name: "ga sopan",
	}

	mockUserReportEntity := userEntity.UserReport{
		UserID:           uint(1),
		ReportedUserID:   uint(2),
		ReportCategoryID: uint(1),
	}

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(1)).Return(mockUserEntity1, nil).Once()
		mockReportCategoryRepo.On("GetReportCategory", uint(1)).Return(mockRcEntity, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, mockReportCategoryRepo, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Report(uint(2), uint(1), uint(1))

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(2)).Return(userEntity.User{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, mockReportCategoryRepo, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Report(uint(1), uint(2), uint(1))

		assert.Error(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("Get", uint(2)).Return(mockUserEntity2, nil).Once()
		mockReportCategoryRepo.On("GetReportCategory", uint(1)).Return(mockRcEntity, nil).Once()
		mockUserRepo.On("StoreReport", mockUserReportEntity).Return(utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, mockReportCategoryRepo, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Report(uint(1), uint(2), uint(1))

		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("Get", uint(2)).Return(mockUserEntity2, nil).Once()
		mockReportCategoryRepo.On("GetReportCategory", uint(1)).Return(mockRcEntity, nil).Once()
		mockUserRepo.On("StoreReport", mockUserReportEntity).Return(nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, mockReportCategoryRepo, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Report(uint(1), uint(2), uint(1))

		assert.NoError(t, err)
	})

	t.Run("user-not-found", func(t *testing.T) {
		mockUserRepo.On("Get", uint(2)).Return(userEntity.User{}, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, mockReportCategoryRepo, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Report(uint(1), uint(2), uint(1))

		assert.Error(t, err)
	})

	t.Run("report-category-not-found", func(t *testing.T) {
		mockUserRepo.On("Get", uint(2)).Return(mockUserEntity2, nil).Once()
		mockReportCategoryRepo.On("GetReportCategory", uint(1)).Return(rcEntity.ReportCategory{}, nil)
		testUserUsecase := NewUserUsecase(mockUserRepo, mockReportCategoryRepo, nil, nil, nil, nil, nil, nil)

		err := testUserUsecase.Report(uint(1), uint(2), uint(1))

		assert.Error(t, err)
	})
}

func TestGetThreadByToken(t *testing.T) {
	mockThreadRepo := threadMock.NewThreadRepository(t)

	mockThreadWithDetailEntity := threadEntity.ThreadWithDetails{
		Thread: threadEntity.Thread{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Title:       "dummy",
			Body:        "dummy",
			ImageURL:    "dummy",
			UserID:      uint(1),
			CommunityID: uint(1),
		},
		User:         mockUserEntity1,
		UpvotesCount: 1,
		IsUpvoted:    1,
		IsFollowed:   1,
		IsDownvoted:  0,
	}

	mockThreadWithDetailEntityArr := []threadEntity.ThreadWithDetails{
		mockThreadWithDetailEntity,
	}

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("GetThreadsByUserID", uint(1), uint(1)).Return(mockThreadWithDetailEntityArr, nil).Once()

		testUserUsecase := NewUserUsecase(nil, nil, nil, nil, mockThreadRepo, nil, nil, nil)
		res, err := testUserUsecase.GetThreadByToken(uint(1), uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockThreadRepo.On("GetThreadsByUserID", uint(1), uint(1)).Return([]threadEntity.ThreadWithDetails{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(nil, nil, nil, nil, mockThreadRepo, nil, nil, nil)
		res, err := testUserUsecase.GetThreadByToken(uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestSendOTP(t *testing.T)   {}
func TestVerifyOTP(t *testing.T) {}

func TestGetReports(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)

	mockBriefReportEntity := userEntity.BriefReport{
		ThreadReportsID:     0,
		UserReportsID:       2,
		CommentReportsID:    0,
		CommunityReportsID:  0,
		CreatedAt:           time.Now(),
		ThreadID:            0,
		UserID:              1,
		CommentID:           0,
		CommunityReportedID: 0,
		ReportCategory:      "dummy",
		Username:            "dummy",
		ProfileImageURL:     "dummy",
		Type:                "dummy",
	}

	mockBriefReportEntityArr := []userEntity.BriefReport{mockBriefReportEntity}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetReports").Return(mockBriefReportEntityArr, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		res, err := testUserUsecase.GetReports("Admin")

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("unauthorize", func(t *testing.T) {
		testUserUsecase := NewUserUsecase(nil, nil, nil, nil, nil, nil, nil, nil)

		res, err := testUserUsecase.GetReports("User")

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("GetReports").Return([]userEntity.BriefReport{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)

		res, err := testUserUsecase.GetReports("Admin")

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestGetDashboardAnalytics(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)

	mockAdminDashboardAnalyticsEntity := userEntity.AdminDashboardAnalytics{
		UsersCount:      1000,
		ModeratorsCount: 1000,
		ReportsCount:    1000,
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetDashboardAnalytics").Return(mockAdminDashboardAnalyticsEntity, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetDashboardAnalytics("Admin")

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("unauthorize", func(t *testing.T) {
		testUserUsecase := NewUserUsecase(nil, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetDashboardAnalytics("User")

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("GetDashboardAnalytics").Return(userEntity.AdminDashboardAnalytics{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetDashboardAnalytics("Admin")

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestGetReportedThread(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)

	mockReportedThreadEntity := userEntity.ReportedThread{
		ID:                      uint(1),
		ThreadTitle:             "dummy",
		ThreadBody:              "dummy",
		ThreadImageURL:          "dummy",
		ThreadCreatedAt:         time.Now(),
		LikesCount:              1,
		ReportedUsername:        "dummy",
		ReportedProfileImageURL: "dummy",
		ReportedUserProfession:  "dummy",
		ReportCategory:          "dummy",
		ReportCreatedAt:         time.Now(),
		Username:                "dummy",
		ProfileImageURL:         "dummy",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetReportedThread", uint(1)).Return(mockReportedThreadEntity, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetReportedThread("Admin", uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("unauthorize", func(t *testing.T) {
		testUserUsecase := NewUserUsecase(nil, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetReportedThread("User", uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockUserRepo.On("GetReportedThread", uint(1)).Return(userEntity.ReportedThread{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetReportedThread("Admin", uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestGetReportedCommunity(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)

	mockReportedCommunityEntity := userEntity.ReportedCommunity{
		ID:                          uint(1),
		CommunityName:               "dummy",
		CommunityImageURL:           "dummy",
		CommunityBackgroundImageURL: "dummy",
		ReportCategory:              "dummy",
		ReportCreatedAt:             time.Now(),
		Username:                    "dummy",
		ProfileImageURL:             "dummy",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetReportedCommunity", uint(1)).Return(mockReportedCommunityEntity, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetReportedCommunity("Admin", uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("unauthorize", func(t *testing.T) {
		testUserUsecase := NewUserUsecase(nil, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetReportedCommunity("User", uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("inteernal-server-error", func(t *testing.T) {
		mockUserRepo.On("GetReportedCommunity", uint(1)).Return(userEntity.ReportedCommunity{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetReportedCommunity("Admin", uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestGetReportedComment(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)

	mockReportedCommentEntity := userEntity.ReportedComment{
		ID:                      uint(1),
		CommentBody:             "dummy",
		LikesCount:              1,
		CommentCreatedAt:        time.Now(),
		ReportedUsername:        "dummy",
		ReportedProfileImageURL: "dummy",
		ReportCategory:          "dummy",
		ReportCreatedAt:         time.Now(),
		Username:                "dummy",
		ProfileImageURL:         "dummy",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetReportedComment", uint(1)).Return(mockReportedCommentEntity, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetReportedComment("Admin", uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("unauthorize", func(t *testing.T) {
		testUserUsecase := NewUserUsecase(nil, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetReportedComment("User", uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("inteernal-server-error", func(t *testing.T) {
		mockUserRepo.On("GetReportedComment", uint(1)).Return(userEntity.ReportedComment{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetReportedComment("Admin", uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestGetReportedUser(t *testing.T) {
	mockUserRepo := userMock.NewUserRepository(t)

	mockReportedUserEntity := userEntity.ReportedUser{
		ID:                          uint(1),
		ReportedUserUsername:        "dummy",
		ReportedUserName:            "dummy",
		ReportedUserProfession:      "dummy",
		ReporteduserBio:             "dummy",
		ReportedUserProfileImageURL: "dummy",
		ReportedUserBackgroundURL:   "dummy",
		ReportingUserUsername:       "dummy",
		ReportingUserName:           "dummy",
		FollowersCount:              1,
		FollowingCount:              1,
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetReportedUser", uint(1)).Return(mockReportedUserEntity, nil).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetReportedUser("Admin", uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("unauthorize", func(t *testing.T) {
		testUserUsecase := NewUserUsecase(nil, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetReportedUser("User", uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("inteernal-server-error", func(t *testing.T) {
		mockUserRepo.On("GetReportedUser", uint(1)).Return(userEntity.ReportedUser{}, utils.ErrInternalServerError).Once()

		testUserUsecase := NewUserUsecase(mockUserRepo, nil, nil, nil, nil, nil, nil, nil)
		res, err := testUserUsecase.GetReportedUser("Admin", uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}
