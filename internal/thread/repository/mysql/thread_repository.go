package repository

import (
	"fmt"
	"macaiki/internal/thread"
	"macaiki/internal/thread/entity"
	"macaiki/pkg/utils"
	"strings"

	"gorm.io/gorm"
)

type ThreadRepositoryImpl struct {
	db *gorm.DB
}

func CreateNewThreadRepository(db *gorm.DB) thread.ThreadRepository {
	return &ThreadRepositoryImpl{db: db}
}

func (tr *ThreadRepositoryImpl) SetThreadImage(imageURL string, threadID uint) error {
	fmt.Println(imageURL)
	res := tr.db.Model(&entity.Thread{}).Where("id = ?", threadID).Update("image_url", imageURL)

	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error.Error() == "record not found" {
			return utils.ErrNotFound
		}
		return utils.ErrInternalServerError
	}

	return nil
}

func (tr *ThreadRepositoryImpl) GetThreadByID(threadID uint) (entity.Thread, error) {
	var thread entity.Thread
	res := tr.db.First(&thread, threadID)
	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error.Error() == "record not found" {
			return entity.Thread{}, utils.ErrNotFound
		}
		return entity.Thread{}, utils.ErrInternalServerError
	}

	return thread, nil
}

func (tr *ThreadRepositoryImpl) CreateThread(thread entity.Thread) (entity.Thread, error) {
	res := tr.db.Create(&thread)
	if res.Error != nil {
		fmt.Println(res.Error)
		return entity.Thread{}, utils.ErrInternalServerError
	}

	return thread, nil
}

func (tr *ThreadRepositoryImpl) DeleteThread(threadID uint) error {
	res := tr.db.Delete(&entity.Thread{}, threadID)
	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error.Error() == "record not found" {
			return utils.ErrNotFound
		}
		return utils.ErrInternalServerError
	}

	return nil
}

func (tr *ThreadRepositoryImpl) UpdateThread(threadID uint, thread entity.Thread) error {
	res := tr.db.Model(&entity.Thread{}).Where("id", threadID).Updates(thread)
	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error.Error() == "record not found" {
			return utils.ErrNotFound
		}
		return utils.ErrInternalServerError
	}

	return nil
}

func (tr *ThreadRepositoryImpl) UpvoteThread(threadUpvote entity.ThreadUpvote) error {
	res := tr.db.Create(&threadUpvote)

	if res.Error != nil {
		fmt.Println(res.Error)
		if strings.HasPrefix(res.Error.Error(), "Error 1452: Cannot add or update a child row") {
			return utils.ErrNotFound
		} else if strings.HasPrefix(res.Error.Error(), "Error 1062: Duplicate entry") {
			return utils.ErrDuplicateEntry
		}
		return utils.ErrInternalServerError
	}

	return nil
}

func (tr *ThreadRepositoryImpl) GetTrendingThreads(userID uint) ([]entity.ThreadWithDetails, error) {
	var threads []entity.ThreadWithDetails
	// TODO: retrieve name, profile URL, etc
	res := tr.db.Raw("SELECT t.*, NOT ISNULL(t3.id) AS is_upvoted, NOT ISNULL(t4.user_id) AS is_followed, NOT ISNULL(t5.id) AS is_downvoted, users.name, users.profile_image_url, users.profession FROM threads t LEFT JOIN (SELECT thread_id, COUNT(*) AS upvotes_count FROM thread_upvotes tu WHERE DATEDIFF(NOW(), tu.created_at) < 7  GROUP BY thread_id) AS t2 ON t.id = t2.thread_id LEFT JOIN (SELECT * FROM thread_upvotes tu WHERE tu.user_id = ?) AS t3 ON t.id = t3.thread_id LEFT JOIN (SELECT user_id FROM user_followers uf WHERE uf.follower_id = ?) AS t4 ON t4.user_id = t.user_id LEFT JOIN users ON users.id = t.user_id LEFT JOIN (SELECT id, thread_id FROM thread_downvotes td WHERE td.user_id = ?) AS t5 ON t5.thread_id = t.id ORDER BY upvotes_count DESC;", userID, userID, userID).Scan(&threads)
	if res.Error != nil {
		return []entity.ThreadWithDetails{}, res.Error
	}

	return threads, nil
}

func (tr *ThreadRepositoryImpl) GetThreadsFromFollowedCommunity(userID uint) ([]entity.ThreadWithDetails, error) {
	var threads []entity.ThreadWithDetails

	res := tr.db.Raw("SELECT t.*, NOT ISNULL(t4.id) AS is_upvoted, NOT ISNULL(t5.user_id) AS is_followed, NOT ISNULL(t6.id) AS is_downvoted, users.name, users.profile_image_url, users.profession FROM threads t LEFT JOIN (SELECT thread_id, COUNT(*) AS upvotes_count FROM thread_upvotes tu GROUP BY thread_id) AS t2 ON t.id = t2.thread_id INNER JOIN (SELECT * FROM followed_communities fc WHERE fc.user_id = ?) AS t3 ON t.community_id = t3.community_id INNER JOIN users ON users.id = t.user_id LEFT JOIN (SELECT * FROM thread_upvotes tu WHERE tu.user_id = ?) AS t4 ON t.id = t4.thread_id LEFT JOIN (SELECT user_id FROM user_followers uf WHERE uf.follower_id= ?) AS t5 ON t5.user_id = t.user_id LEFT JOIN (SELECT id, thread_id FROM thread_downvotes td WHERE td.user_id = ?) AS t6 ON t6.thread_id = t.id;", userID, userID, userID, userID).Scan(&threads)

	if res.Error != nil {
		fmt.Println(res.Error)
		return []entity.ThreadWithDetails{}, utils.ErrInternalServerError
	}

	return threads, nil
}

func (tr *ThreadRepositoryImpl) GetThreadsFromFollowedUsers(userID uint) ([]entity.ThreadWithDetails, error) {
	var threads []entity.ThreadWithDetails

	res := tr.db.Raw("SELECT t.*, NOT ISNULL(t4.id) AS is_upvoted, NOT ISNULL(t3.user_id) AS is_followed, NOT ISNULL(t5.id) AS is_downvoted, users.name, users.profile_image_url, users.profession  FROM threads t LEFT JOIN (SELECT thread_id, COUNT(*) AS upvotes_count FROM thread_upvotes tu GROUP BY thread_id) AS t2 ON t.id = t2.thread_id INNER JOIN (SELECT user_id FROM user_followers uf WHERE uf.follower_id= ?) AS t3 ON t3.user_id = t.user_id LEFT JOIN users ON users.id = t.user_id LEFT JOIN (SELECT * FROM thread_upvotes tu WHERE tu.user_id = ?) AS t4 ON t.id = t4.thread_id LEFT JOIN (SELECT id, thread_id FROM thread_downvotes td WHERE td.user_id = ?) AS t5 ON t5.thread_id = t.id;", userID, userID, userID).Scan(&threads)

	if res.Error != nil {
		fmt.Println(res.Error)
		return []entity.ThreadWithDetails{}, utils.ErrInternalServerError
	}

	return threads, nil
}

func (tr *ThreadRepositoryImpl) AddThreadComment(comment entity.Comment) error {
	res := tr.db.Create(&comment)

	if res.Error != nil {
		fmt.Println(res.Error)
		return utils.ErrInternalServerError
	}

	return nil
}

func (tr *ThreadRepositoryImpl) GetCommentsByThreadID(threadID uint) ([]entity.CommentDetails, error) {
	var comments []entity.CommentDetails
	res := tr.db.Raw("SELECT comments.*, users.*, t2.likes_count FROM comments LEFT JOIN (SELECT comment_id, COUNT(*) AS likes_count FROM comment_likes cl GROUP BY comment_id) AS t2 ON comments.id = t2.comment_id INNER JOIN users ON comments.user_id = users.id WHERE comments.thread_id = ?", threadID).Scan(&comments)

	if res.Error != nil {
		fmt.Println(res.Error)
		return []entity.CommentDetails{}, utils.ErrInternalServerError
	}

	return comments, nil
}

func (tr *ThreadRepositoryImpl) GetThreads(keyword string, userID uint) ([]entity.ThreadWithDetails, error) {
	var threads []entity.ThreadWithDetails

	res := tr.db.Raw("SELECT combined.*, upvotes_count, NOT ISNULL(t4.id) AS is_upvoted, NOT ISNULL(t3.user_id) AS is_followed, NOT ISNULL(t5.id) AS is_downvoted, users.name, users.profile_image_url, users.profession FROM (SELECT * FROM threads t WHERE t.body LIKE ? OR t.title LIKE ? UNION SELECT t.* FROM comments c LEFT JOIN threads t ON t.id = c.thread_id WHERE c.body LIKE ?) AS combined LEFT JOIN (SELECT thread_id, COUNT(*) AS upvotes_count FROM thread_upvotes tu GROUP BY thread_id) AS t2 ON combined.id = t2.thread_id LEFT JOIN (SELECT user_id FROM user_followers uf WHERE uf.follower_id= ?) AS t3 ON t3.user_id = combined.user_id LEFT JOIN users ON users.id = combined.user_id LEFT JOIN (SELECT * FROM thread_upvotes tu WHERE tu.user_id = ?) AS t4 ON combined.id = t4.thread_id LEFT JOIN (SELECT id, thread_id, user_id FROM thread_downvotes td WHERE td.user_id = ?) AS t5 ON t5.thread_id = combined.id;", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", userID, userID, userID).Scan(&threads)

	if res.Error != nil {
		fmt.Println(res.Error)
		return []entity.ThreadWithDetails{}, utils.ErrInternalServerError
	}

	return threads, nil
}

func (tr *ThreadRepositoryImpl) LikeComment(commentLikes entity.CommentLikes) error {
	res := tr.db.Create(&commentLikes)

	if res.Error != nil {
		fmt.Println(res.Error)
		if strings.HasPrefix(res.Error.Error(), "Error 1452: Cannot add or update a child row") {
			return utils.ErrNotFound
		}
		return utils.ErrInternalServerError
	}

	return nil
}

func (tr *ThreadRepositoryImpl) DownvoteThread(downvote entity.ThreadDownvote) error {
	res := tr.db.Create(&downvote)

	if res.Error != nil {
		fmt.Println(res.Error)
		if strings.HasPrefix(res.Error.Error(), "Error 1452: Cannot add or update a child row") {
			return utils.ErrNotFound
		} else if strings.HasPrefix(res.Error.Error(), "Error 1062: Duplicate entry") {
			return utils.ErrDuplicateEntry
		}
		return utils.ErrInternalServerError
	}

	return nil
}

func (tr *ThreadRepositoryImpl) UndoDownvoteThread(threadID, userID uint) error {
	res := tr.db.Unscoped().Delete(&entity.ThreadDownvote{}, "thread_id = ? AND user_id = ?", threadID, userID)

	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error.Error() == "record not found" {
			return utils.ErrNotFound
		}
		return utils.ErrInternalServerError
	}

	return nil
}

func (tr *ThreadRepositoryImpl) UnlikeComment(commentID, userID uint) error {
	res := tr.db.Unscoped().Delete(&entity.CommentLikes{}, "thread_id = ? AND user_id = ?", commentID, userID)

	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error.Error() == "record not found" {
			return utils.ErrNotFound
		}
		return utils.ErrInternalServerError
	}

	return nil
}

func (tr *ThreadRepositoryImpl) UndoUpvoteThread(commentID, userID uint) error {
	res := tr.db.Unscoped().Delete(&entity.ThreadUpvote{}, "thread_id = ? AND user_id = ?", commentID, userID)

	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error.Error() == "record not found" {
			return utils.ErrNotFound
		}
		return utils.ErrInternalServerError
	}

	return nil
}

func (tr *ThreadRepositoryImpl) GetThreadDownvotes(threadID, userID uint) (entity.ThreadDownvote, error) {
	var threadDownvotes entity.ThreadDownvote
	res := tr.db.First(&threadDownvotes, "thread_id = ? AND user_id = ?", threadID, userID)

	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error.Error() == "record not found" {
			return threadDownvotes, utils.ErrNotFound
		}
		return threadDownvotes, utils.ErrInternalServerError
	}

	return threadDownvotes, nil
}

func (tr *ThreadRepositoryImpl) GetThreadUpvotes(threadID, userID uint) (entity.ThreadUpvote, error) {
	var threadUpvote entity.ThreadUpvote
	res := tr.db.First(&threadUpvote, "thread_id = ? AND user_id = ?", threadID, userID)

	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error.Error() == "record not found" {
			return threadUpvote, utils.ErrNotFound
		}

		return threadUpvote, utils.ErrInternalServerError
	}

	return threadUpvote, nil
}
