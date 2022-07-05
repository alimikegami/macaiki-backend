package mysql

import (
	reportcategory "macaiki/internal/report_category"
	"macaiki/internal/report_category/entity"

	"gorm.io/gorm"
)

type ReportCategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewReportCategoryRepository(db *gorm.DB) reportcategory.ReportCategoryRepository {
	return &ReportCategoryRepositoryImpl{db}
}

func (rcr *ReportCategoryRepositoryImpl) StoreReportCategory(reportCategory entity.ReportCategory) error {
	tx := rcr.db.Create(&reportCategory)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}

func (rcr *ReportCategoryRepositoryImpl) GetAllReportCategory() ([]entity.ReportCategory, error) {
	reportCategories := []entity.ReportCategory{}

	tx := rcr.db.Find(&reportCategories)
	err := tx.Error
	if err != nil {
		return []entity.ReportCategory{}, err
	}

	return reportCategories, nil
}

func (rcr *ReportCategoryRepositoryImpl) GetReportCategory(id uint) (entity.ReportCategory, error) {
	reportCategory := entity.ReportCategory{}

	tx := rcr.db.Find(&reportCategory, id)
	err := tx.Error
	if err != nil {
		return entity.ReportCategory{}, err
	}

	return reportCategory, nil
}

func (rcr *ReportCategoryRepositoryImpl) UpdateReportCategory(reportCategory entity.ReportCategory) error {
	tx := rcr.db.Save(&reportCategory)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}

func (rcr *ReportCategoryRepositoryImpl) DeleteReportCategory(reportCategory entity.ReportCategory) error {
	tx := rcr.db.Delete(&reportCategory)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}
