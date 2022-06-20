package domain

import "macaiki/internal/report_category/dto"

type ReportCategory struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type ReportCategoryUsecase interface {
	GetAllReportCategory() ([]dto.ReportCategoryResponse, error)
	GetReportCategory(id uint) (dto.ReportCategoryResponse, error)
}

type ReportCategoryRepository interface {
	GetAllReportCategory() ([]ReportCategory, error)
	GetReportCategory(id uint) (ReportCategory, error)
}
