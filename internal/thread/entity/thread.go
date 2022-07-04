package entity

import (
	userEntity "macaiki/internal/user/entity"

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
	User     userEntity.User
}

type ThreadDownvote struct {
	gorm.Model
	UserID   uint
	ThreadID uint
	Thread   Thread
	User     userEntity.User
}

type ThreadWithDetails struct {
	Thread
	userEntity.User
	LikesCount int
	IsLiked    int
	IsFollowed int
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
	userEntity.User
	LikesCount int
	IsLiked    int
}

type CommentLikes struct {
	gorm.Model
	UserID    uint
	CommentID uint
}
