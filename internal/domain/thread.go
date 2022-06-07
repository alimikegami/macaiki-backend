package domain

import (
	"macaiki/internal/thread/dto"

	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	Title   string
	Body    string
	UserID  uint
	TopicID uint
}

type ThreadUseCase interface {
	CreateThread(thread dto.ThreadRequest, userID uint) error
	DeleteThread(threadID uint) error
}

type ThreadRepository interface {
	CreateThread(thread Thread) error
	DeleteThread(threadID uint) error
}
