package repository

import (
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

func (tr *ThreadRepositoryImpl) CreateThread(thread domain.Thread) error {
	res := tr.db.Create(&thread)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (tr *ThreadRepositoryImpl) DeleteThread(threadID uint) error {
	if res := tr.db.Delete(&domain.Thread{}, threadID); res.Error != nil {
		return res.Error
	}
	return nil
}

// func (tr *ThreadRepositoryImpl) UpdateThread(thread domain.Thread) error {
// 	if res := tr.db.Update(); res.Error != nil {
// 		return res.Error
// 	}
// 	return nil
// }
