package usecase

import (
	"macaiki/internal/domain"
	"macaiki/internal/report_category/dto"

	"github.com/go-playground/validator/v10"
)

type ReportCategoryUsecaseImpl struct {
	rcRepo    domain.ReportCategoryRepository
	validator *validator.Validate
}

func NewReportCategoryUsecase(rcRepo domain.ReportCategoryRepository, validator *validator.Validate) domain.ReportCategoryUsecase {
	return &ReportCategoryUsecaseImpl{rcRepo, validator}
}

func (rcu *ReportCategoryUsecaseImpl) GetAllReportCategory() ([]dto.ReportCategoryResponse, error) {
	reportCategories, err := rcu.rcRepo.GetAllReportCategory()
	if err != nil {
		return []dto.ReportCategoryResponse{}, domain.ErrInternalServerError
	}

	dtoReportCategories := []dto.ReportCategoryResponse{}
	for _, val := range reportCategories {
		dtoReportCategories = append(dtoReportCategories, dto.ReportCategoryResponse{ID: val.ID, Name: val.Name})
	}

	return dtoReportCategories, nil
}

func (rcu *ReportCategoryUsecaseImpl) GetReportCategory(id uint) (dto.ReportCategoryResponse, error) {
	reportCategory, err := rcu.rcRepo.GetReportCategory(id)
	if err != nil {
		return dto.ReportCategoryResponse{}, domain.ErrInternalServerError
	}
	if reportCategory.ID == 0 {
		return dto.ReportCategoryResponse{}, domain.ErrNotFound
	}

	return dto.ReportCategoryResponse{ID: reportCategory.ID, Name: reportCategory.Name}, nil
}
