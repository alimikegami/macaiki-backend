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
		return []dto.ReportCategoryResponse{}, err
	}

	dtoReportCategories := []dto.ReportCategoryResponse{}
	for _, val := range reportCategories {
		dtoReportCategories = append(dtoReportCategories, dto.ReportCategoryResponse{ID: val.ID, Name: val.Name})
	}

	return dtoReportCategories, nil
}
