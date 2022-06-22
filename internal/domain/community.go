package domain

import "gorm.io/gorm"

type Community struct {
	gorm.Model
	Name string
}
