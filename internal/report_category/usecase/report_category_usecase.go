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

func (rcu *ReportCategoryUsecaseImpl) CreateReportCategory(reportCategory dto.ReportCategoryRequest, role string) error {
	if role != "Admin" {
		return domain.ErrUnauthorizedAccess
	}

	if err := rcu.validator.Struct(reportCategory); err != nil {
		return domain.ErrBadParamInput
	}

	reportCategoryEntity := domain.ReportCategory{
		Name: reportCategory.Name,
	}

	err := rcu.rcRepo.StoreReportCategory(reportCategoryEntity)
	if err != nil {
		return domain.ErrInternalServerError
	}

	return nil
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

func (rcu *ReportCategoryUsecaseImpl) UpdateReportCategory(reportCategory dto.ReportCategoryRequest, id uint, role string) error {
	if role != "Admin" {
		return domain.ErrUnauthorizedAccess
	}

	if err := rcu.validator.Struct(reportCategory); err != nil {
		return domain.ErrBadParamInput
	}

	reportCategoryDB, err := rcu.rcRepo.GetReportCategory(id)
	if err != nil {
		return domain.ErrInternalServerError
	}
	if reportCategoryDB.ID == 0 {
		return domain.ErrNotFound
	}

	reportCategoryDB.Name = reportCategory.Name
	err = rcu.rcRepo.UpdateReportCategory(reportCategoryDB)
	if err != nil {
		return domain.ErrInternalServerError
	}

	return nil
}
func (rcu *ReportCategoryUsecaseImpl) DeleteReportCategory(id uint, role string) error {
	if role != "Admin" {
		return domain.ErrUnauthorizedAccess
	}

	reportCategory, err := rcu.rcRepo.GetReportCategory(id)
	if err != nil {
		return domain.ErrInternalServerError
	}
	if reportCategory.ID == 0 {
		return domain.ErrNotFound
	}

	err = rcu.rcRepo.UpdateReportCategory(reportCategory)
	if err != nil {
		return domain.ErrInternalServerError
	}

	return nil
}
