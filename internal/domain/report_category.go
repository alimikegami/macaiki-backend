package domain

import "macaiki/internal/report_category/dto"

type ReportCategory struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	UserReports []UserReport `gorm:"foreignKey:ReportCategoryID"`
}

type ReportCategoryUsecase interface {
	CreateReportCategory(reportCategory dto.ReportCategoryRequest, role string) error
	GetAllReportCategory() ([]dto.ReportCategoryResponse, error)
	GetReportCategory(id uint) (dto.ReportCategoryResponse, error)
	UpdateReportCategory(reportCategory dto.ReportCategoryRequest, id uint, role string) error
	DeleteReportCategory(id uint, role string) error
}

type ReportCategoryRepository interface {
	StoreReportCategory(reportCategory ReportCategory) error
	GetAllReportCategory() ([]ReportCategory, error)
	GetReportCategory(id uint) (ReportCategory, error)
	UpdateReportCategory(reportCategory ReportCategory) error
	DeleteReportCategory(reportCategory ReportCategory) error
}
