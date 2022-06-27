package mysql

import (
	"macaiki/internal/domain"

	"gorm.io/gorm"
)

type ReportCategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewReportCategoryRepository(db *gorm.DB) domain.ReportCategoryRepository {
	return &ReportCategoryRepositoryImpl{db}
}

func (rcr *ReportCategoryRepositoryImpl) StoreReportCategory(reportCategory domain.ReportCategory) error {
	tx := rcr.db.Create(&reportCategory)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}

func (rcr *ReportCategoryRepositoryImpl) GetAllReportCategory() ([]domain.ReportCategory, error) {
	reportCategories := []domain.ReportCategory{}

	tx := rcr.db.Find(&reportCategories)
	err := tx.Error
	if err != nil {
		return []domain.ReportCategory{}, err
	}

	return reportCategories, nil
}

func (rcr *ReportCategoryRepositoryImpl) GetReportCategory(id uint) (domain.ReportCategory, error) {
	reportCategory := domain.ReportCategory{}

	tx := rcr.db.Find(&reportCategory, id)
	err := tx.Error
	if err != nil {
		return domain.ReportCategory{}, err
	}

	return reportCategory, nil
}

func (rcr *ReportCategoryRepositoryImpl) UpdateReportCategory(reportCategory domain.ReportCategory) error {
	tx := rcr.db.Save(&reportCategory)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}

func (rcr *ReportCategoryRepositoryImpl) DeleteReportCategory(reportCategory domain.ReportCategory) error {
	tx := rcr.db.Delete(&reportCategory)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}
