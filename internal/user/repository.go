package user

import "macaiki/internal/user/entity"

type UserRepository interface {
	GetAllWithDetail(userID uint, search string) ([]entity.User, error)
	Store(user entity.User) error
	Get(id uint) (entity.User, error)
	GetWithDetail(id, tokenID uint) (entity.User, error)
	Update(userDB *entity.User, user entity.User) (entity.User, error)
	Delete(id uint) error
	GetByEmail(email string) (entity.User, error)
	GetByUsername(username string) (entity.User, error)

	GetFollowerNumber(id uint) (int, error)
	GetFollowingNumber(id uint) (int, error)
	GetThreadsNumber(id uint) (int, error)
	Follow(user, userFollower entity.User) (entity.User, error)
	Unfollow(user, userFollower entity.User) (entity.User, error)
	GetFollower(userID, getFollowingUserID uint) ([]entity.User, error)
	GetFollowing(userID, getFollowingUserID uint) ([]entity.User, error)
	SetUserImage(id uint, imageURL string, tableName string) error

	StoreReport(userReport entity.UserReport) error
	StoreOTP(VerifyEmail entity.VerificationEmail) error
	GetOTP(email string) (entity.VerificationEmail, error)
	GetReports() ([]entity.BriefReport, error)
	GetUserReport(reportID uint) (entity.UserReport, error)

	GetDashboardAnalytics() (entity.AdminDashboardAnalytics, error)
	GetReportedThread(threadReportID uint) (entity.ReportedThread, error)
	GetReportedCommunity(communityReportID uint) (entity.ReportedCommunity, error)
	GetReportedComment(commentReportID uint) (entity.ReportedComment, error)
	GetReportedUser(userReportID uint) (entity.ReportedUser, error)

	DeleteUserReport(userReportID uint) error
	DeleteThreadReport(threadReportID uint) error
	DeleteCommunityReport(communityReportID uint) error
	DeleteCommentReport(commentReportID uint) error
}
