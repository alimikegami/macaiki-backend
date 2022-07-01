package thread

import (
	"macaiki/internal/thread/dto"
	"mime/multipart"
)

type ThreadUseCase interface {
	CreateThread(thread dto.ThreadRequest, userID uint) (dto.ThreadResponse, error)
	DeleteThread(threadID uint, userID uint) error
	UpdateThread(thread dto.ThreadRequest, threadID uint, userID uint) (dto.ThreadResponse, error)
	GetThreadByID(threadID uint) (dto.ThreadResponse, error)
	SetThreadImage(img *multipart.FileHeader, threadID uint, userID uint) error
	LikeThread(threadID uint, userID uint) error
	GetTrendingThreads(userID uint) ([]dto.DetailedThreadResponse, error)
	GetThreadsFromFollowedCommunity(userID uint) ([]dto.DetailedThreadResponse, error)
	GetThreadsFromFollowedUsers(userID uint) ([]dto.DetailedThreadResponse, error)
	AddThreadComment(dto.CommentRequest) error
	GetCommentsByThreadID(threadID uint) ([]dto.CommentResponse, error)
	GetThreads(keyword string, userID uint) ([]dto.DetailedThreadResponse, error)
}
