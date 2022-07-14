package usecase

import (
	entityMocks "macaiki/internal/notification/mocks"
	"macaiki/internal/thread/dto"
	"macaiki/internal/thread/entity"
	"macaiki/internal/thread/mocks"
	userEntity "macaiki/internal/user/entity"
	"macaiki/pkg/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	mockedEntity = entity.Thread{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:       "Title",
		Body:        "Body",
		ImageURL:    "ImageURL",
		UserID:      uint(1),
		CommunityID: uint(1),
	}

	mockedThreadDTOReq = dto.ThreadRequest{
		Title:       "Title",
		Body:        "Body",
		CommunityID: uint(1),
	}

	mockedEntityMappedFromDTO = entity.Thread{
		Title:       "Title",
		Body:        "Body",
		CommunityID: uint(1),
	}

	// mockedDTOResponse = dto.ThreadResponse{
	// 	ID:          1,
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// 	Title:       "Title",
	// 	Body:        "Body",
	// 	ImageURL:    "ImageURL",
	// 	UserID:      uint(1),
	// 	CommunityID: uint(1),
	// }

	mockedLikeCommentEntity = entity.CommentLikes{
		UserID:    1,
		CommentID: 1,
	}

	mockedCommentReportDTO = dto.CommentReportRequest{
		CommentID:        1,
		UserID:           1,
		ReportCategoryID: 1,
	}

	mockedCommentReportEntity = entity.CommentReport{
		CommentID:        1,
		UserID:           1,
		ReportCategoryID: 1,
	}

	mockedSavedThreadDTO = dto.SavedThreadRequest{
		UserID:   1,
		ThreadID: 1,
	}

	mockedSavedThreadEntity = entity.SavedThread{
		UserID:   1,
		ThreadID: 1,
	}

	mockedCommentLikesEntity = entity.CommentLikes{
		CommentID: 1,
		UserID:    1,
	}

	mockedDetailedThread = []entity.ThreadWithDetails{{Thread: entity.Thread{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:       "Title",
		Body:        "Body",
		ImageURL:    "ImageURL",
		UserID:      uint(1),
		CommunityID: uint(1),
	},
		User: userEntity.User{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Email:              "alim1@gmail.com",
			Username:           "alim1",
			Password:           "123",
			Name:               "alim1",
			ProfileImageUrl:    "asdfas.jpg",
			BackgroundImageUrl: "asfds.jpg",
			Bio:                "asdfa",
			Profession:         "asfasf",
			Role:               "User",
		}}}
)

func TestCreateThreadReport(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)
	mockNotifRepo := entityMocks.NewNotificationRepository(t)

	mockThreadReportReq := dto.ThreadReportRequest{
		UserID:           1,
		ThreadID:         1,
		ReportCategoryID: 1,
	}

	mockThreadReportEntity := entity.ThreadReport{
		UserID:           1,
		ThreadID:         1,
		ReportCategoryID: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("CreateThreadReport", mockThreadReportEntity).Return(nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)

		err := testThreadUseCase.CreateThreadReport(mockThreadReportReq)
		assert.NoError(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockThreadRepo.On("CreateThreadReport", mockThreadReportEntity).Return(utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)

		err := testThreadUseCase.CreateThreadReport(mockThreadReportReq)
		assert.Error(t, err)
	})
}

func TestCreateThread(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)
	mockNotifRepo := entityMocks.NewNotificationRepository(t)

	mockThreadReq := dto.ThreadRequest{
		Title:       "lorem ipsum",
		Body:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi at lacinia lectus, ac commodo turpis. Suspendisse pretium tortor non purus viverra, vel viverra elit vehicula. Morbi lacus metus, euismod at porttitor at, lobortis ac metus. Ut a molestie felis, id tristique est. Aliquam consectetur dui in fermentum commodo. Nam ut euismod erat. Donec mollis vulputate arcu, et hendrerit ante. Fusce aliquam, eros sit amet blandit tristique, massa enim interdum purus, a tempor felis ipsum sollicitudin felis. Fusce pulvinar commodo massa, id tincidunt nulla laoreet ut. Nullam id nisl nec metus scelerisque ultrices et ac mi. Nunc vel diam est. In euismod molestie venenatis. Fusce elementum commodo magna, nec vulputate urna tristique accumsan.",
		CommunityID: 1,
	}

	mockThreadEntity := entity.Thread{
		Title:       "lorem ipsum",
		Body:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi at lacinia lectus, ac commodo turpis. Suspendisse pretium tortor non purus viverra, vel viverra elit vehicula. Morbi lacus metus, euismod at porttitor at, lobortis ac metus. Ut a molestie felis, id tristique est. Aliquam consectetur dui in fermentum commodo. Nam ut euismod erat. Donec mollis vulputate arcu, et hendrerit ante. Fusce aliquam, eros sit amet blandit tristique, massa enim interdum purus, a tempor felis ipsum sollicitudin felis. Fusce pulvinar commodo massa, id tincidunt nulla laoreet ut. Nullam id nisl nec metus scelerisque ultrices et ac mi. Nunc vel diam est. In euismod molestie venenatis. Fusce elementum commodo magna, nec vulputate urna tristique accumsan.",
		CommunityID: 1,
		UserID:      1,
	}

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("CreateThread", mockThreadEntity).Return(mockedEntity, nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)

		res, err := testThreadUseCase.CreateThread(mockThreadReq, 1)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockThreadRepo.On("CreateThread", mockThreadEntity).Return(entity.Thread{}, utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)

		res, err := testThreadUseCase.CreateThread(mockThreadReq, uint(1))
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestDeleteThread(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)
	mockNotifRepo := entityMocks.NewNotificationRepository(t)

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(mockedEntity, nil).Once()

		mockThreadRepo.On("DeleteThread", uint(1)).Return(nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)

		err := testThreadUseCase.DeleteThread(uint(1), uint(1), "")
		assert.NoError(t, err)
	})

	t.Run("internal-server-error-on-get-threads-by-ID", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(entity.Thread{}, utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)

		err := testThreadUseCase.DeleteThread(uint(1), uint(1), "")
		assert.Error(t, err)
	})

	t.Run("internal-server-error-on-delete-thread", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(mockedEntity, nil).Once()

		mockThreadRepo.On("DeleteThread", uint(1)).Return(utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)

		err := testThreadUseCase.DeleteThread(uint(1), uint(1), "")
		assert.Error(t, err)
	})

	t.Run("record-not-found", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(entity.Thread{}, utils.ErrNotFound).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)

		err := testThreadUseCase.DeleteThread(uint(1), uint(1), "")
		assert.Error(t, err)
	})

	t.Run("unauthorized-access", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(mockedEntity, nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)

		err := testThreadUseCase.DeleteThread(uint(1), uint(2), "")
		assert.Error(t, err)
	})

	t.Run("success-admin-deletion", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(mockedEntity, nil).Once()

		mockThreadRepo.On("DeleteThread", uint(1)).Return(nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)

		err := testThreadUseCase.DeleteThread(uint(1), uint(3), "Admin")
		assert.NoError(t, err)
	})
}

func TestUpdateThread(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)
	mockNotifRepo := entityMocks.NewNotificationRepository(t)

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(mockedEntity, nil).Twice()

		mockThreadRepo.On("UpdateThread", uint(1), mockedEntityMappedFromDTO).Return(nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		res, err := testThreadUseCase.UpdateThread(mockedThreadDTOReq, uint(1), uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("internal-server-error-on-get-threads-by-id", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(entity.Thread{}, utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		res, err := testThreadUseCase.UpdateThread(mockedThreadDTOReq, uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("record-not-found", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(entity.Thread{}, utils.ErrNotFound).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		res, err := testThreadUseCase.UpdateThread(mockedThreadDTOReq, uint(1), uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestGetThreadByID(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)
	mockNotifRepo := entityMocks.NewNotificationRepository(t)

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(mockedEntity, nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		res, err := testThreadUseCase.GetThreadByID(uint(1))
		assert.NotEmpty(t, res)
		assert.NoError(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(entity.Thread{}, utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		res, err := testThreadUseCase.GetThreadByID(uint(1))
		assert.Empty(t, res)
		assert.Error(t, err)
	})

	t.Run("record-not-found", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(entity.Thread{}, utils.ErrNotFound).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		res, err := testThreadUseCase.GetThreadByID(uint(1))
		assert.Empty(t, res)
		assert.Error(t, err)
	})
}

func TestLikeComment(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)
	mockNotifRepo := entityMocks.NewNotificationRepository(t)

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("LikeComment", mockedLikeCommentEntity).Return(nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		err := testThreadUseCase.LikeComment(uint(1), uint(1))
		assert.NoError(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockThreadRepo.On("LikeComment", mockedLikeCommentEntity).Return(utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		err := testThreadUseCase.LikeComment(uint(1), uint(1))
		assert.Error(t, err)
	})
}

func TestCreateCommentReport(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)
	mockNotifRepo := entityMocks.NewNotificationRepository(t)

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("CreateCommentReport", mockedCommentReportEntity).Return(nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		err := testThreadUseCase.CreateCommentReport(mockedCommentReportDTO)

		assert.NoError(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockThreadRepo.On("CreateCommentReport", mockedCommentReportEntity).Return(utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		err := testThreadUseCase.CreateCommentReport(mockedCommentReportDTO)

		assert.Error(t, err)
	})
}

func TestStoreSavedThread(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)
	mockNotifRepo := entityMocks.NewNotificationRepository(t)

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("StoreSavedThread", mockedSavedThreadEntity).Return(nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		err := testThreadUseCase.StoreSavedThread(mockedSavedThreadDTO)

		assert.NoError(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockThreadRepo.On("StoreSavedThread", mockedSavedThreadEntity).Return(utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		err := testThreadUseCase.StoreSavedThread(mockedSavedThreadDTO)

		assert.Error(t, err)
	})
}

func TestUnlikeComment(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)
	mockNotifRepo := entityMocks.NewNotificationRepository(t)

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("UnlikeComment", uint(1), uint(1)).Return(nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		err := testThreadUseCase.UnlikeComment(uint(1), uint(1))

		assert.NoError(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockThreadRepo.On("UnlikeComment", uint(1), uint(1)).Return(utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		err := testThreadUseCase.UnlikeComment(uint(1), uint(1))

		assert.Error(t, err)
	})
}

func TestGetSavedThread(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)
	mockNotifRepo := entityMocks.NewNotificationRepository(t)

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("GetSavedThread", uint(1)).Return(mockedDetailedThread, nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		thread, err := testThreadUseCase.GetSavedThread(uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, thread)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockThreadRepo.On("GetSavedThread", uint(1)).Return([]entity.ThreadWithDetails{}, utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		thread, err := testThreadUseCase.GetSavedThread(uint(1))

		assert.Error(t, err)
		assert.Empty(t, thread)
	})
}

func TestGetTrendingThreads(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)
	mockNotifRepo := entityMocks.NewNotificationRepository(t)

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("GetTrendingThreads", uint(1)).Return(mockedDetailedThread, nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		thread, err := testThreadUseCase.GetTrendingThreads(uint(1), -1)

		assert.NoError(t, err)
		assert.NotEmpty(t, thread)
	})

	t.Run("success-with-limit", func(t *testing.T) {
		mockThreadRepo.On("GetTrendingThreadsWithLimit", uint(1), 3).Return(mockedDetailedThread, nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		thread, err := testThreadUseCase.GetTrendingThreads(uint(1), 3)

		assert.NoError(t, err)
		assert.NotEmpty(t, thread)
	})

	t.Run("internal-server-error-with-limit", func(t *testing.T) {
		mockThreadRepo.On("GetTrendingThreadsWithLimit", uint(1), 3).Return([]entity.ThreadWithDetails{}, utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		thread, err := testThreadUseCase.GetTrendingThreads(uint(1), 3)

		assert.Error(t, err)
		assert.Empty(t, thread)
	})
}

func TestGetThreadsFromFollowedCommunity(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)
	mockNotifRepo := entityMocks.NewNotificationRepository(t)

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("GetThreadsFromFollowedCommunity", uint(1)).Return(mockedDetailedThread, nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		thread, err := testThreadUseCase.GetThreadsFromFollowedCommunity(uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, thread)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockThreadRepo.On("GetThreadsFromFollowedCommunity", uint(1)).Return([]entity.ThreadWithDetails{}, utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		thread, err := testThreadUseCase.GetThreadsFromFollowedCommunity(uint(1))

		assert.Error(t, err)
		assert.Empty(t, thread)
	})
}

func TestGetThreadsFromFollowedUsers(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)
	mockNotifRepo := entityMocks.NewNotificationRepository(t)

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("GetThreadsFromFollowedUsers", uint(1)).Return(mockedDetailedThread, nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		thread, err := testThreadUseCase.GetThreadsFromFollowedUsers(uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, thread)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockThreadRepo.On("GetThreadsFromFollowedUsers", uint(1)).Return([]entity.ThreadWithDetails{}, utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		thread, err := testThreadUseCase.GetThreadsFromFollowedUsers(uint(1))

		assert.Error(t, err)
		assert.Empty(t, thread)
	})
}

func TestGetThreads(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)
	mockNotifRepo := entityMocks.NewNotificationRepository(t)

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("GetThreads", "", uint(1)).Return(mockedDetailedThread, nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		thread, err := testThreadUseCase.GetThreads("", uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, thread)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockThreadRepo.On("GetThreads", "", uint(1)).Return([]entity.ThreadWithDetails{}, utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
		thread, err := testThreadUseCase.GetThreads("", uint(1))

		assert.Error(t, err)
		assert.Empty(t, thread)
	})
}

// func TestGetCommentsByThreadID(t *testing.T) {
// 	mockThreadRepo := mocks.NewThreadRepository(t)
// 	mockNotifRepo := entityMocks.NewNotificationRepository(t)

// 	t.Run("success", func(t *testing.T) {
// 		mockThreadRepo.On("GetCommentsByThreadID", uint(1)).Return(mockedDetailedThread, nil).Once()

// 		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, mockNotifRepo, nil)
// 		thread, err := testThreadUseCase.GetCommentsByThreadID(uint(1))

// 		assert.NoError(t, err)
// 		assert.NotEmpty(t, thread)
// 	})
// }
