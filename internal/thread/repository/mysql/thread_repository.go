package repository

import (
	"errors"
	"fmt"
	"macaiki/internal/domain"

	"gorm.io/gorm"
)

type ThreadRepositoryImpl struct {
	db *gorm.DB
}

func CreateNewThreadRepository(db *gorm.DB) domain.ThreadRepository {
	return &ThreadRepositoryImpl{db: db}
}

func (tr *ThreadRepositoryImpl) GetThreads() ([]domain.Thread, error) {
	var threads []domain.Thread
	res := tr.db.Find(&threads)

	if res.Error != nil {
		return []domain.Thread{}, res.Error
	}

	return threads, nil
}

func (tr *ThreadRepositoryImpl) SetThreadImage(imageURL string, threadID uint) error {
	fmt.Println(imageURL)
	res := tr.db.Model(&domain.Thread{}).Where("id = ?", threadID).Update("image_url", imageURL)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected < 1 {
		return errors.New("resource does not exists")
	}

	return nil
}

func (tr *ThreadRepositoryImpl) GetThreadByID(threadID uint) (domain.Thread, error) {
	var thread domain.Thread
	res := tr.db.First(&thread, threadID)
	if res.Error != nil {
		return thread, res.Error
	}

	if res.RowsAffected < 1 {
		return thread, errors.New("resource does not exists")
	}

	return thread, nil
}

func (tr *ThreadRepositoryImpl) CreateThread(thread domain.Thread) (domain.Thread, error) {
	res := tr.db.Create(&thread)
	if res.Error != nil {
		return domain.Thread{}, res.Error
	}

	return thread, nil
}

func (tr *ThreadRepositoryImpl) DeleteThread(threadID uint) error {
	if res := tr.db.Delete(&domain.Thread{}, threadID); res.Error != nil {
		return res.Error
	}
	return nil
}

func (tr *ThreadRepositoryImpl) UpdateThread(threadID uint, thread domain.Thread) error {
	res := tr.db.Model(&domain.Thread{}).Where("id", threadID).Updates(thread)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected < 1 {
		return errors.New("no affected rows")
	}

	return nil
}
