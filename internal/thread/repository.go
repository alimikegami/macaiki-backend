package thread

import "macaiki/internal/thread/entity"

type ThreadRepository interface {
	CreateThread(thread entity.Thread) (entity.Thread, error)
	DeleteThread(threadID uint) error
	UpdateThread(threadID uint, thread entity.Thread) error
	GetThreadByID(threadID uint) (entity.Thread, error)
	SetThreadImage(imageURL string, threadID uint) error
	UpvoteThread(threadUpvote entity.ThreadUpvote) error
	UndoUpvoteThread(threadID, userID uint) error
	GetTrendingThreads(userID uint) ([]entity.ThreadWithDetails, error)
	GetThreadsFromFollowedCommunity(userID uint) ([]entity.ThreadWithDetails, error)
	GetThreadsFromFollowedUsers(userID uint) ([]entity.ThreadWithDetails, error)
	AddThreadComment(comment entity.Comment) error
	GetCommentsByThreadID(threadID uint) ([]entity.CommentDetails, error)
	GetThreads(keyword string, userID uint) ([]entity.ThreadWithDetails, error)
	LikeComment(commentLikes entity.CommentLikes) error
	UnlikeComment(commentID, userID uint) error
	DownvoteThread(downvote entity.ThreadDownvote) error
	UndoDownvoteThread(threadID, userID uint) error
	GetThreadDownvotes(threadID, userID uint) (entity.ThreadDownvote, error)
	GetThreadUpvotes(threadID, userID uint) (entity.ThreadUpvote, error)
	DeleteComment(commentID uint) error
	GetCommentByID(commentID uint) (entity.Comment, error)
	CreateThreadReport(threadReport entity.ThreadReport) error
}
