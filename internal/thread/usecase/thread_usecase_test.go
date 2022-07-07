package usecase

import (
	"macaiki/internal/thread/dto"
	"macaiki/internal/thread/entity"
	"macaiki/internal/thread/mocks"
	"macaiki/pkg/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateThreadReport(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)

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

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, nil)

		err := testThreadUseCase.CreateThreadReport(mockThreadReportReq)
		assert.NoError(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockThreadRepo.On("CreateThreadReport", mockThreadReportEntity).Return(utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, nil)

		err := testThreadUseCase.CreateThreadReport(mockThreadReportReq)
		assert.Error(t, err)
	})
}

func TestCreateThread(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)

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

	createdMockThreadEntity := entity.Thread{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:       "lorem ipsum",
		Body:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi at lacinia lectus, ac commodo turpis. Suspendisse pretium tortor non purus viverra, vel viverra elit vehicula. Morbi lacus metus, euismod at porttitor at, lobortis ac metus. Ut a molestie felis, id tristique est. Aliquam consectetur dui in fermentum commodo. Nam ut euismod erat. Donec mollis vulputate arcu, et hendrerit ante. Fusce aliquam, eros sit amet blandit tristique, massa enim interdum purus, a tempor felis ipsum sollicitudin felis. Fusce pulvinar commodo massa, id tincidunt nulla laoreet ut. Nullam id nisl nec metus scelerisque ultrices et ac mi. Nunc vel diam est. In euismod molestie venenatis. Fusce elementum commodo magna, nec vulputate urna tristique accumsan.",
		ImageURL:    "",
		UserID:      1,
		CommunityID: 1,
	}

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("CreateThread", mockThreadEntity).Return(createdMockThreadEntity, nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, nil)

		res, err := testThreadUseCase.CreateThread(mockThreadReq, 1)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockThreadRepo.On("CreateThread", mockThreadEntity).Return(entity.Thread{}, utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, nil)

		res, err := testThreadUseCase.CreateThread(mockThreadReq, uint(1))
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestDeleteThread(t *testing.T) {
	mockThreadRepo := mocks.NewThreadRepository(t)

	t.Run("success", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(entity.Thread{
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
		}, nil).Once()

		mockThreadRepo.On("DeleteThread", uint(1)).Return(nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, nil)

		err := testThreadUseCase.DeleteThread(uint(1), uint(1), "")
		assert.NoError(t, err)
	})

	t.Run("internal-server-error-on-get-threads-by-ID", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(entity.Thread{}, utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, nil)

		err := testThreadUseCase.DeleteThread(uint(1), uint(1), "")
		assert.Error(t, err)
	})

	t.Run("internal-server-error-on-delete-thread", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(entity.Thread{
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
		}, nil).Once()

		mockThreadRepo.On("DeleteThread", uint(1)).Return(utils.ErrInternalServerError).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, nil)

		err := testThreadUseCase.DeleteThread(uint(1), uint(1), "")
		assert.Error(t, err)
	})

	t.Run("record-not-found", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(entity.Thread{}, utils.ErrNotFound).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, nil)

		err := testThreadUseCase.DeleteThread(uint(1), uint(1), "")
		assert.Error(t, err)
	})

	t.Run("unauthorized-access", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(entity.Thread{
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
		}, nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, nil)

		err := testThreadUseCase.DeleteThread(uint(1), uint(2), "")
		assert.Error(t, err)
	})

	t.Run("success-admin-deletion", func(t *testing.T) {
		mockThreadRepo.On("GetThreadByID", uint(1)).Return(entity.Thread{
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
		}, nil).Once()

		mockThreadRepo.On("DeleteThread", uint(1)).Return(nil).Once()

		testThreadUseCase := CreateNewThreadUseCase(mockThreadRepo, nil)

		err := testThreadUseCase.DeleteThread(uint(1), uint(3), "Admin")
		assert.NoError(t, err)
	})
}
