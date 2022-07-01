package usecase

import (
	reportcategory "macaiki/internal/report_category"
	"macaiki/internal/report_category/dto"
	"macaiki/internal/report_category/entity"
	"macaiki/pkg/utils"

	"github.com/go-playground/validator/v10"
)

type ReportCategoryUsecaseImpl struct {
	rcRepo    reportcategory.ReportCategoryRepository
	validator *validator.Validate
}

func NewReportCategoryUsecase(rcRepo reportcategory.ReportCategoryRepository, validator *validator.Validate) reportcategory.ReportCategoryUsecase {
	return &ReportCategoryUsecaseImpl{rcRepo, validator}
}

func (rcu *ReportCategoryUsecaseImpl) CreateReportCategory(reportCategory dto.ReportCategoryRequest, role string) error {
	if role != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	if err := rcu.validator.Struct(reportCategory); err != nil {
		return utils.ErrBadParamInput
	}

	reportCategoryEntity := entity.ReportCategory{
		Name: reportCategory.Name,
	}

	err := rcu.rcRepo.StoreReportCategory(reportCategoryEntity)
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
}

func (rcu *ReportCategoryUsecaseImpl) GetAllReportCategory() ([]dto.ReportCategoryResponse, error) {
	reportCategories, err := rcu.rcRepo.GetAllReportCategory()
	if err != nil {
		return []dto.ReportCategoryResponse{}, utils.ErrInternalServerError
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
		return dto.ReportCategoryResponse{}, utils.ErrInternalServerError
	}
	if reportCategory.ID == 0 {
		return dto.ReportCategoryResponse{}, utils.ErrNotFound
	}

	return dto.ReportCategoryResponse{ID: reportCategory.ID, Name: reportCategory.Name}, nil
}

func (rcu *ReportCategoryUsecaseImpl) UpdateReportCategory(reportCategory dto.ReportCategoryRequest, id uint, role string) error {
	if role != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	if err := rcu.validator.Struct(reportCategory); err != nil {
		return utils.ErrBadParamInput
	}

	reportCategoryDB, err := rcu.rcRepo.GetReportCategory(id)
	if err != nil {
		return utils.ErrInternalServerError
	}
	if reportCategoryDB.ID == 0 {
		return utils.ErrNotFound
	}

	reportCategoryDB.Name = reportCategory.Name
	err = rcu.rcRepo.UpdateReportCategory(reportCategoryDB)
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
}

func (rcu *ReportCategoryUsecaseImpl) DeleteReportCategory(id uint, role string) error {
	if role != "Admin" {
		return utils.ErrUnauthorizedAccess
	}

	reportCategory, err := rcu.rcRepo.GetReportCategory(id)
	if err != nil {
		return utils.ErrInternalServerError
	}
	if reportCategory.ID == 0 {
		return utils.ErrNotFound
	}

	err = rcu.rcRepo.UpdateReportCategory(reportCategory)
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
}
