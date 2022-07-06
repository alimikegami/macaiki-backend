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

type ThreadUpvote struct {
	gorm.Model
	UserID   uint `gorm:"index:unique_upvote,unique"`
	ThreadID uint `gorm:"index:unique_upvote,unique"`
	Thread   Thread
	User     userEntity.User
}

type ThreadDownvote struct {
	gorm.Model
	UserID   uint `gorm:"index:unique_downvote,unique"`
	ThreadID uint `gorm:"index:unique_downvote,unique"`
	Thread   Thread
	User     userEntity.User
}

type ThreadWithDetails struct {
	Thread
	userEntity.User
	UpvotesCount int
	IsUpvoted    int
	IsFollowed   int
	IsDownvoted  int
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
	Thread    Thread
}

type CommentDetails struct {
	Comment
	userEntity.User
	LikesCount int
	IsLiked    int
}

type CommentLikes struct {
	gorm.Model
	UserID    uint `gorm:"index:unique_likes,unique"`
	CommentID uint `gorm:"index:unique_likes,unique"`
}
