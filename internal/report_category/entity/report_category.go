package entity

import (
	userEntity "macaiki/internal/user/entity"
)

type ReportCategory struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	UserReports []userEntity.UserReport `gorm:"foreignKey:ReportCategoryID"`
}
