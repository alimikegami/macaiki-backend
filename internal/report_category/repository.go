package reportcategory

import "macaiki/internal/report_category/entity"

type ReportCategoryRepository interface {
	StoreReportCategory(reportCategory entity.ReportCategory) error
	GetAllReportCategory() ([]entity.ReportCategory, error)
	GetReportCategory(id uint) (entity.ReportCategory, error)
	UpdateReportCategory(reportCategory entity.ReportCategory) error
	DeleteReportCategory(reportCategory entity.ReportCategory) error
}
