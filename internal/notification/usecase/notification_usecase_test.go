package usecase

import (
	entity "macaiki/internal/notification/entity"
	"macaiki/internal/notification/mocks"
	threadMocks "macaiki/internal/thread/mocks"
	userEntity "macaiki/internal/user/entity"
	userMocks "macaiki/internal/user/mocks"
	"macaiki/pkg/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestReadAllNotifications(t *testing.T) {
	notificationMockRepo := mocks.NewNotificationRepository(t)
	userMockRepo := userMocks.NewUserRepository(t)
	threadMockRepo := threadMocks.NewThreadRepository(t)
	t.Run("success", func(t *testing.T) {
		notificationMockRepo.On("ReadAllNotifications", uint(1)).Return(nil).Once()
		notificationMockRepo.On("GetAllNotifications", uint(1)).Return([]entity.Notification{
			{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				UserID:            uint(1),
				NotificationType:  "Follow You",
				NotificationRefID: uint(2),
				IsReaded:          1,
			},
		}, nil).Once()

		userMockRepo.On("Get", uint(1)).Return(userEntity.User{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Email:              "alimikegami1@gmail.com",
			Username:           "alimikegami1",
			Password:           "asafs",
			Name:               "alim ikegami",
			ProfileImageUrl:    "adsf.jpg",
			BackgroundImageUrl: "sdfas.jpg",
			Bio:                "asfsafds",
			Profession:         "sdfas",
			Role:               "Admin",
		}, nil).Once()

		testNotificationUseCase := NewNotificationUsecase(notificationMockRepo, userMockRepo, threadMockRepo)

		notifications, err := testNotificationUseCase.ReadAllNotifications(uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, notifications)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		notificationMockRepo.On("ReadAllNotifications", uint(1)).Return(utils.ErrInternalServerError).Once()

		testNotificationUseCase := NewNotificationUsecase(notificationMockRepo, userMockRepo, threadMockRepo)

		notifications, err := testNotificationUseCase.ReadAllNotifications(uint(1))

		assert.Error(t, err)
		assert.Empty(t, notifications)
	})
}

func TestDeleteAllNotifications(t *testing.T) {
	notificationMockRepo := mocks.NewNotificationRepository(t)
	userMockRepo := userMocks.NewUserRepository(t)
	threadMockRepo := threadMocks.NewThreadRepository(t)

	t.Run("success", func(t *testing.T) {
		notificationMockRepo.On("DeleleteAllNotifications", uint(1)).Return(nil).Once()

		notificationMockRepo.On("GetAllNotifications", uint(1)).Return([]entity.Notification{
			{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				UserID:            uint(1),
				NotificationType:  "Follow You",
				NotificationRefID: uint(2),
				IsReaded:          1,
			},
		}, nil).Once()

		userMockRepo.On("Get", uint(1)).Return(userEntity.User{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Email:              "alimikegami1@gmail.com",
			Username:           "alimikegami1",
			Password:           "asafs",
			Name:               "alim ikegami",
			ProfileImageUrl:    "adsf.jpg",
			BackgroundImageUrl: "sdfas.jpg",
			Bio:                "asfsafds",
			Profession:         "sdfas",
			Role:               "Admin",
		}, nil).Once()

		testNotificationUseCase := NewNotificationUsecase(notificationMockRepo, userMockRepo, threadMockRepo)

		notifications, err := testNotificationUseCase.DeleteAllNotifications(uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, notifications)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		notificationMockRepo.On("DeleleteAllNotifications", uint(1)).Return(utils.ErrInternalServerError).Once()

		testNotificationUseCase := NewNotificationUsecase(notificationMockRepo, userMockRepo, threadMockRepo)

		notifications, err := testNotificationUseCase.DeleteAllNotifications(uint(1))

		assert.Error(t, err)
		assert.Empty(t, notifications)
	})
}
