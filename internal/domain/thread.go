package domain

import (
	"macaiki/internal/thread/dto"

	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	Title    string
	Body     string
	ImageURL string
	UserID   uint
	TopicID  uint
}

type ThreadUseCase interface {
	CreateThread(thread dto.ThreadRequest, userID uint) (dto.ThreadResponse, error)
	DeleteThread(threadID uint) error
	UpdateThread(thread dto.ThreadRequest, threadID uint, userID uint) (dto.ThreadResponse, error)
}

type ThreadRepository interface {
	CreateThread(thread Thread) (Thread, error)
	DeleteThread(threadID uint) error
	UpdateThread(threadID uint, thread Thread) error
	GetThreadByID(threadID uint) (Thread, error)
}
