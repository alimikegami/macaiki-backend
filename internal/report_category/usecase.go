package reportcategory

import "macaiki/internal/report_category/dto"

type ReportCategoryUsecase interface {
	CreateReportCategory(reportCategory dto.ReportCategoryRequest, role string) error
	GetAllReportCategory() ([]dto.ReportCategoryResponse, error)
	GetReportCategory(id uint) (dto.ReportCategoryResponse, error)
	UpdateReportCategory(reportCategory dto.ReportCategoryRequest, id uint, role string) error
	DeleteReportCategory(id uint, role string) error
}
