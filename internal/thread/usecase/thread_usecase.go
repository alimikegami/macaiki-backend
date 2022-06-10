package usecase

import (
	"macaiki/internal/domain"
	"macaiki/internal/thread/dto"
)

type ThreadUseCaseImpl struct {
	tr domain.ThreadRepository
}

func CreateNewThreadUseCase(tr domain.ThreadRepository) domain.ThreadUseCase {
	return &ThreadUseCaseImpl{tr: tr}
}

func (tuc *ThreadUseCaseImpl) CreateThread(thread dto.ThreadRequest, userID uint) error {
	threadEntity := domain.Thread{
		Title:   thread.Title,
		Body:    thread.Body,
		UserID:  userID,
		TopicID: thread.TopicID,
	}

	err := tuc.tr.CreateThread(threadEntity)
	return err
}

func (tuc *ThreadUseCaseImpl) DeleteThread(threadID uint) error {
	// TODO: add validation logic to make sure the only user that can delete a thread is either the admin or the user who created the thread
	err := tuc.tr.DeleteThread(threadID)
	return err
}
