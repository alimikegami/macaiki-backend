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

type ThreadWithDetails struct {
	Thread
	User
	LikesCount int
}

type ThreadFollower struct {
	gorm.Model
	ThreadID uint
	UserID   uint
}

type Comment struct {
	gorm.Model
	Body      string
	UserID    uint
	ThreadID  uint
	CommentID uint
}

type CommentDetails struct {
	Comment
	User
	LikesCount int
	IsLiked    int
}

type ThreadUseCase interface {
	CreateThread(thread dto.ThreadRequest, userID uint) (dto.ThreadResponse, error)
	DeleteThread(threadID uint, userID uint) error
	GetThreads() ([]dto.ThreadResponse, error)
	UpdateThread(thread dto.ThreadRequest, threadID uint, userID uint) (dto.ThreadResponse, error)
	GetThreadByID(threadID uint) (dto.ThreadResponse, error)
	SetThreadImage(img *multipart.FileHeader, threadID uint, userID uint) error
	LikeThread(threadID uint, userID uint) error
	GetTrendingThreads(userID uint) ([]dto.DetailedThreadResponse, error)
	GetThreadsFromFollowedCommunity(userID uint) ([]dto.DetailedThreadResponse, error)
	GetThreadsFromFollowedUsers(userID uint) ([]dto.DetailedThreadResponse, error)
	AddThreadComment(dto.CommentRequest) error
	GetCommentsByThreadID(threadID uint) ([]dto.CommentResponse, error)
}

type ThreadRepository interface {
	CreateThread(thread Thread) (Thread, error)
	DeleteThread(threadID uint) error
	GetThreads() ([]Thread, error)
	UpdateThread(threadID uint, thread Thread) error
	GetThreadByID(threadID uint) (Thread, error)
	SetThreadImage(imageURL string, threadID uint) error
	LikeThread(threadLikes ThreadLikes) error
	GetTrendingThreads(userID uint) ([]ThreadWithDetails, error)
	GetThreadsFromFollowedCommunity(userID uint) ([]ThreadWithDetails, error)
	GetThreadsFromFollowedUsers(userID uint) ([]ThreadWithDetails, error)
	AddThreadComment(comment Comment) error
	GetCommentsByThreadID(threadID uint) ([]CommentDetails, error)
}
