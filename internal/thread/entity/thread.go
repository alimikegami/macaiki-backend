package entity

import (
	communityentity "macaiki/internal/community/entity"
	reportCategoryEntity "macaiki/internal/report_category/entity"
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
	User        userEntity.User
	Community   communityentity.Community
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
	Thread   Thread
	User     userEntity.User
}

type Comment struct {
	gorm.Model
	Body      string
	UserID    uint
	ThreadID  uint
	CommentID uint
	Thread    Thread
	User      userEntity.User
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
	User      userEntity.User
	Comment   Comment
}

type ThreadReport struct {
	gorm.Model
	UserID           uint
	ThreadID         uint
	ReportCategoryID uint
	User             userEntity.User
	Thread           Thread
	ReportCategory   reportCategoryEntity.ReportCategory
}

type CommentReport struct {
	gorm.Model
	UserID           uint
	CommentID        uint
	ReportCategoryID uint
	User             userEntity.User
	Comment          Comment
	ReportCategory   reportCategoryEntity.ReportCategory
}

type SavedThread struct {
	gorm.Model
	UserID   uint `gorm:"index:unique_saved_thread,unique"`
	ThreadID uint `gorm:"index:unique_saved_thread,unique"`
	User     userEntity.User
	Thread   Thread
}
