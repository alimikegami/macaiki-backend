package user

import "macaiki/internal/user/entity"

type UserRepository interface {
	GetAllWithDetail(userID uint, search string) ([]entity.User, error)
	Store(user entity.User) error
	Get(id uint) (entity.User, error)
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
	GetReports() ([]entity.BriefReport, error)
}
