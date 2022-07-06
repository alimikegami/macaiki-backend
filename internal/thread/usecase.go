package thread

import (
	"macaiki/internal/thread/dto"
	"mime/multipart"
)

type ThreadUseCase interface {
	CreateThread(thread dto.ThreadRequest, userID uint) (dto.ThreadResponse, error)
	DeleteThread(threadID uint, userID uint, role string) error
	UpdateThread(thread dto.ThreadRequest, threadID uint, userID uint) (dto.ThreadResponse, error)
	GetThreadByID(threadID uint) (dto.ThreadResponse, error)
	SetThreadImage(img *multipart.FileHeader, threadID uint, userID uint) error
	UpvoteThread(threadID uint, userID uint) error
	UndoUpvoteThread(threadID, userID uint) error
	GetTrendingThreads(userID uint) ([]dto.DetailedThreadResponse, error)
	GetThreadsFromFollowedCommunity(userID uint) ([]dto.DetailedThreadResponse, error)
	GetThreadsFromFollowedUsers(userID uint) ([]dto.DetailedThreadResponse, error)
	AddThreadComment(dto.CommentRequest) error
	GetCommentsByThreadID(threadID uint) ([]dto.CommentResponse, error)
	GetThreads(keyword string, userID uint) ([]dto.DetailedThreadResponse, error)
	LikeComment(commentID, userID uint) error
	UnlikeComment(commentID, userID uint) error
	DownvoteThread(threadID uint, userID uint) error
	UndoDownvoteThread(threadID, userID uint) error
	DeleteComment(commentID uint, threadID uint, userID uint, role string) error
	CreateThreadReport(threadReport dto.ThreadReportRequest) error
	CreateCommentReport(commentReport dto.CommentReportRequest) error
	StoreSavedThread(savedThread dto.SavedThreadRequest) error
}
