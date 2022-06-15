package domain

import (
	"macaiki/internal/thread/dto"
	"mime/multipart"

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
	GetThreads() ([]dto.ThreadResponse, error)
	UpdateThread(thread dto.ThreadRequest, threadID uint, userID uint) (dto.ThreadResponse, error)
	GetThreadByID(threadID uint) (dto.ThreadResponse, error)
	SetThreadImage(img *multipart.FileHeader, threadID uint) error
}

type ThreadRepository interface {
	CreateThread(thread Thread) (Thread, error)
	DeleteThread(threadID uint) error
	GetThreads() ([]Thread, error)
	UpdateThread(threadID uint, thread Thread) error
	GetThreadByID(threadID uint) (Thread, error)
	SetThreadImage(imageURL string, threadID uint) error
}
