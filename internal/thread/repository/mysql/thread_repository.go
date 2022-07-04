package repository

import (
	"errors"
	"fmt"
	"macaiki/internal/thread"
	"macaiki/internal/thread/entity"

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
		return res.Error
	}

	if res.RowsAffected < 1 {
		return errors.New("resource does not exists")
	}

	return nil
}

func (tr *ThreadRepositoryImpl) GetThreadByID(threadID uint) (entity.Thread, error) {
	var thread entity.Thread
	res := tr.db.First(&thread, threadID)
	if res.Error != nil {
		return thread, res.Error
	}

	if res.RowsAffected < 1 {
		return thread, errors.New("resource does not exists")
	}

	return thread, nil
}

func (tr *ThreadRepositoryImpl) CreateThread(thread entity.Thread) (entity.Thread, error) {
	res := tr.db.Create(&thread)
	if res.Error != nil {
		return entity.Thread{}, res.Error
	}

	return thread, nil
}

func (tr *ThreadRepositoryImpl) DeleteThread(threadID uint) error {
	if res := tr.db.Delete(&entity.Thread{}, threadID); res.Error != nil {
		return res.Error
	}
	return nil
}

func (tr *ThreadRepositoryImpl) UpdateThread(threadID uint, thread entity.Thread) error {
	res := tr.db.Model(&entity.Thread{}).Where("id", threadID).Updates(thread)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected < 1 {
		return errors.New("no affected rows")
	}

	return nil
}

func (tr *ThreadRepositoryImpl) LikeThread(threadLikes entity.ThreadLikes) error {
	res := tr.db.Create(&threadLikes)

	return res.Error
}

func (tr *ThreadRepositoryImpl) GetTrendingThreads(userID uint) ([]entity.ThreadWithDetails, error) {
	var threads []entity.ThreadWithDetails
	// TODO: retrieve name, profile URL, etc
	res := tr.db.Raw("SELECT t.*, NOT ISNULL(t3.id) AS is_liked, NOT ISNULL(t4.user_id) AS is_followed, users.name, users.profile_image_url, users.profession FROM threads t LEFT JOIN (SELECT thread_id, COUNT(*) AS likes_count FROM thread_likes tl WHERE DATEDIFF(NOW(), tl.created_at) < 7  GROUP BY thread_id) AS t2 ON t.id = t2.thread_id LEFT JOIN (SELECT * FROM thread_likes tl WHERE tl.user_id = ?) AS t3 ON t.id = t3.thread_id LEFT JOIN (SELECT user_id FROM user_followers uf WHERE uf.follower_id = ?) AS t4 ON t4.user_id = t.user_id LEFT JOIN users ON users.id = t.user_id ORDER BY likes_count DESC;", userID, userID).Scan(&threads)
	if res.Error != nil {
		return []entity.ThreadWithDetails{}, res.Error
	}

	return threads, nil
}

func (tr *ThreadRepositoryImpl) GetThreadsFromFollowedCommunity(userID uint) ([]entity.ThreadWithDetails, error) {
	var threads []entity.ThreadWithDetails

	res := tr.db.Raw("SELECT t.*, NOT ISNULL(t4.id) AS is_liked, NOT ISNULL(t5.user_id) AS is_followed, users.name, users.profile_image_url, users.profession FROM threads t LEFT JOIN (SELECT thread_id, COUNT(*) AS likes_count FROM thread_likes tl GROUP BY thread_id) AS t2 ON t.id = t2.thread_id INNER JOIN (SELECT * FROM followed_communities fc WHERE fc.user_id = ?) AS t3 ON t.community_id = t3.community_id INNER JOIN users ON users.id = t.user_id LEFT JOIN (SELECT * FROM thread_likes tl WHERE tl.user_id = ?) AS t4 ON t.id = t4.thread_id LEFT JOIN (SELECT user_id FROM user_followers uf WHERE uf.follower_id= ?) AS t5 ON t5.user_id = t.user_id;", userID, userID, userID).Scan(&threads)

	if res.Error != nil {
		return []entity.ThreadWithDetails{}, res.Error
	}

	return threads, nil
}

func (tr *ThreadRepositoryImpl) GetThreadsFromFollowedUsers(userID uint) ([]entity.ThreadWithDetails, error) {
	var threads []entity.ThreadWithDetails

	res := tr.db.Raw("SELECT t.*, NOT ISNULL(t4.id) AS is_liked, NOT ISNULL(t3.user_id) AS is_followed, users.name, users.profile_image_url, users.profession  FROM threads t LEFT JOIN (SELECT thread_id, COUNT(*) AS likes_count FROM thread_likes tl GROUP BY thread_id) AS t2 ON t.id = t2.thread_id INNER JOIN (SELECT user_id FROM user_followers uf WHERE uf.follower_id= ?) AS t3 ON t3.user_id = t.user_id LEFT JOIN users ON users.id = t.user_id LEFT JOIN (SELECT * FROM thread_likes tl WHERE tl.user_id = ?) AS t4 ON t.id = t4.thread_id;", userID, userID).Scan(&threads)

	if res.Error != nil {
		return []entity.ThreadWithDetails{}, res.Error
	}

	return threads, nil
}

func (tr *ThreadRepositoryImpl) AddThreadComment(comment entity.Comment) error {
	res := tr.db.Create(&comment)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (tr *ThreadRepositoryImpl) GetCommentsByThreadID(threadID uint) ([]entity.CommentDetails, error) {
	var comments []entity.CommentDetails
	res := tr.db.Raw("SELECT comments.*, users.*, t2.likes_count FROM comments LEFT JOIN (SELECT comment_id, COUNT(*) AS likes_count FROM comment_likes cl GROUP BY comment_id) AS t2 ON comments.id = t2.comment_id INNER JOIN users ON comments.user_id = users.id WHERE comments.thread_id = ?", threadID).Scan(&comments)

	if res.Error != nil {
		return []entity.CommentDetails{}, res.Error
	}

	return comments, nil
}

func (tr *ThreadRepositoryImpl) GetThreads(keyword string, userID uint) ([]entity.ThreadWithDetails, error) {
	var threads []entity.ThreadWithDetails

	res := tr.db.Raw("SELECT combined.*, likes_count, NOT ISNULL(t4.id) AS is_liked, NOT ISNULL(t3.user_id) AS is_followed, users.name, users.profile_image_url, users.profession FROM (SELECT * FROM threads t WHERE t.body LIKE ? OR t.title LIKE ? UNION SELECT t.* FROM comments c LEFT JOIN threads t ON t.id = c.thread_id WHERE c.body LIKE ?) AS combined LEFT JOIN (SELECT thread_id, COUNT(*) AS likes_count FROM thread_likes tl GROUP BY thread_id) AS t2 ON combined.id = t2.thread_id LEFT JOIN (SELECT user_id FROM user_followers uf WHERE uf.follower_id= ?) AS t3 ON t3.user_id = combined.user_id LEFT JOIN users ON users.id = combined.user_id LEFT JOIN (SELECT * FROM thread_likes tl WHERE tl.user_id = ?) AS t4 ON combined.id = t4.thread_id;", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", userID, userID).Scan(&threads)

	if res.Error != nil {
		return []entity.ThreadWithDetails{}, res.Error
	}

	return threads, nil
}

func (tr *ThreadRepositoryImpl) LikeComment(commentLikes entity.CommentLikes) error {
	res := tr.db.Create(&commentLikes)

	return res.Error
}

func (tr *ThreadRepositoryImpl) DownvoteThread(downvote entity.ThreadDownvote) error {
	res := tr.db.Create(&downvote)

	return res.Error
}
