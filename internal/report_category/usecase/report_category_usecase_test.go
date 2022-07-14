package usecase

import (
	"macaiki/internal/report_category/dto"
	"macaiki/internal/report_category/entity"
	"macaiki/internal/report_category/mocks"
	"macaiki/pkg/utils"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var (
	mockedReportCategoryDTO = dto.ReportCategoryRequest{
		Name: "Violence",
	}

	mockedReportCategoryEntity = entity.ReportCategory{
		Name: "Violence",
	}

	mockedReportCategoryReturnedEntities = []entity.ReportCategory{
		{
			ID:   1,
			Name: "Violence",
		},
	}

	mockedReportCategoryReturnedEntity = entity.ReportCategory{
		ID:   uint(1),
		Name: "Violence",
	}
)

func TestCreateReportCategory(t *testing.T) {
	mockedReportCategoryRepo := mocks.NewReportCategoryRepository(t)

	t.Run("success", func(t *testing.T) {
		mockedReportCategoryRepo.On("StoreReportCategory", mockedReportCategoryEntity).Return(nil).Once()

		testReportCategoryUseCase := NewReportCategoryUsecase(mockedReportCategoryRepo, validator.New())
		err := testReportCategoryUseCase.CreateReportCategory(mockedReportCategoryDTO, "Admin")

		assert.NoError(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockedReportCategoryRepo.On("StoreReportCategory", mockedReportCategoryEntity).Return(utils.ErrInternalServerError).Once()

		testReportCategoryUseCase := NewReportCategoryUsecase(mockedReportCategoryRepo, validator.New())
		err := testReportCategoryUseCase.CreateReportCategory(mockedReportCategoryDTO, "Admin")

		assert.Error(t, err)
	})
}

func TestGetAllReportCategory(t *testing.T) {
	mockedReportCategoryRepo := mocks.NewReportCategoryRepository(t)

	t.Run("success", func(t *testing.T) {
		mockedReportCategoryRepo.On("GetAllReportCategory").Return(mockedReportCategoryReturnedEntities, nil).Once()

		testReportCategoryUseCase := NewReportCategoryUsecase(mockedReportCategoryRepo, validator.New())
		res, err := testReportCategoryUseCase.GetAllReportCategory()

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockedReportCategoryRepo.On("GetAllReportCategory").Return([]entity.ReportCategory{}, utils.ErrInternalServerError).Once()

		testReportCategoryUseCase := NewReportCategoryUsecase(mockedReportCategoryRepo, validator.New())
		res, err := testReportCategoryUseCase.GetAllReportCategory()

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestGetReportCategory(t *testing.T) {
	mockedReportCategoryRepo := mocks.NewReportCategoryRepository(t)

	t.Run("success", func(t *testing.T) {
		mockedReportCategoryRepo.On("GetReportCategory", uint(1)).Return(mockedReportCategoryReturnedEntity, nil).Once()

		testReportCategoryUseCase := NewReportCategoryUsecase(mockedReportCategoryRepo, validator.New())
		res, err := testReportCategoryUseCase.GetReportCategory(uint(1))

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockedReportCategoryRepo.On("GetReportCategory", uint(1)).Return(entity.ReportCategory{}, utils.ErrInternalServerError).Once()

		testReportCategoryUseCase := NewReportCategoryUsecase(mockedReportCategoryRepo, validator.New())
		res, err := testReportCategoryUseCase.GetReportCategory(uint(1))

		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestDeleteReportCategory(t *testing.T) {
	mockedReportCategoryRepo := mocks.NewReportCategoryRepository(t)

	t.Run("success", func(t *testing.T) {
		mockedReportCategoryRepo.On("GetReportCategory", uint(1)).Return(mockedReportCategoryReturnedEntity, nil).Once()
		mockedReportCategoryRepo.On("DeleteReportCategory", mockedReportCategoryReturnedEntity).Return(nil).Once()

		testReportCategoryUseCase := NewReportCategoryUsecase(mockedReportCategoryRepo, validator.New())
		err := testReportCategoryUseCase.DeleteReportCategory(uint(1), "Admin")

		assert.NoError(t, err)
	})

	t.Run("internal-server-error", func(t *testing.T) {
		mockedReportCategoryRepo.On("GetReportCategory", uint(1)).Return(entity.ReportCategory{}, utils.ErrInternalServerError).Once()

		testReportCategoryUseCase := NewReportCategoryUsecase(mockedReportCategoryRepo, validator.New())
		err := testReportCategoryUseCase.DeleteReportCategory(uint(1), "Admin")

		assert.Error(t, err)
	})
}
