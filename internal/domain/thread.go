package domain

import (
	"macaiki/internal/thread/dto"
	"mime/multipart"

	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	Title       string
	Body        string
	ImageURL    string
	UserID      uint
	CommunityID uint
}

type ThreadLikes struct {
	gorm.Model
	UserID   uint
	ThreadID uint
	Thread   Thread
	User     User
}

type ThreadFollower struct {
	gorm.Model
	ThreadID uint
	UserID   uint
}

type ThreadWithDetails struct {
	Thread
	User
	LikesCount int
}

type ThreadUseCase interface {
	CreateThread(thread dto.ThreadRequest, userID uint) (dto.ThreadResponse, error)
	DeleteThread(threadID uint) error
	GetThreads() ([]dto.ThreadResponse, error)
	UpdateThread(thread dto.ThreadRequest, threadID uint, userID uint) (dto.ThreadResponse, error)
	GetThreadByID(threadID uint) (dto.ThreadResponse, error)
	SetThreadImage(img *multipart.FileHeader, threadID uint) error
	LikeThread(threadID uint, userID uint) error
	GetTrendingThreads() ([]dto.DetailedThreadResponse, error)
	GetThreadsFromFollowedCommunity(userID uint) ([]dto.DetailedThreadResponse, error)
	GetThreadsFromFollowedUsers(userID uint) ([]dto.DetailedThreadResponse, error)
}

type ThreadRepository interface {
	CreateThread(thread Thread) (Thread, error)
	DeleteThread(threadID uint) error
	GetThreads() ([]Thread, error)
	UpdateThread(threadID uint, thread Thread) error
	GetThreadByID(threadID uint) (Thread, error)
	SetThreadImage(imageURL string, threadID uint) error
	LikeThread(threadLikes ThreadLikes) error
	GetTrendingThreads() ([]ThreadWithDetails, error)
	GetThreadsFromFollowedCommunity(userID uint) ([]ThreadWithDetails, error)
	GetThreadsFromFollowedUsers(userID uint) ([]ThreadWithDetails, error)
}
