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

func (tuc *ThreadUseCaseImpl) GetThreads() ([]dto.ThreadResponse, error) {
	var threads []dto.ThreadResponse
	res, err := tuc.tr.GetThreads()

	if err != nil {
		return []dto.ThreadResponse{}, domain.ErrInternalServerError
	}

	for _, thread := range res {
		threads = append(threads, dto.ThreadResponse{
			ID:        thread.ID,
			Title:     thread.Title,
			Body:      thread.Body,
			TopicID:   thread.TopicID,
			ImageURL:  thread.ImageURL,
			UserID:    thread.UserID,
			CreatedAt: thread.CreatedAt,
			UpdatedAt: thread.UpdatedAt,
		})
	}

	return threads, nil
}

func (tuc *ThreadUseCaseImpl) CreateThread(thread dto.ThreadRequest, userID uint) (dto.ThreadResponse, error) {
	threadEntity := domain.Thread{
		Title:   thread.Title,
		Body:    thread.Body,
		UserID:  userID,
		TopicID: thread.TopicID,
	}

	res, err := tuc.tr.CreateThread(threadEntity)
	if err != nil {
		return dto.ThreadResponse{}, err
	}
	return dto.ThreadResponse{
		ID:        res.ID,
		Title:     res.Title,
		Body:      res.Body,
		TopicID:   res.TopicID,
		ImageURL:  res.ImageURL,
		UserID:    res.UserID,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (tuc *ThreadUseCaseImpl) DeleteThread(threadID uint) error {
	// TODO: add validation logic to make sure the only user that can delete a thread is either the admin or the user who created the thread
	err := tuc.tr.DeleteThread(threadID)
	return err
}

func (tuc *ThreadUseCaseImpl) UpdateThread(thread dto.ThreadRequest, threadID uint, userID uint) (dto.ThreadResponse, error) {
	// TODO: add validation logic to make sure the only user that can update a thread is the user who created the thread
	threadEntity := domain.Thread{
		Title:   thread.Title,
		Body:    thread.Body,
		TopicID: thread.TopicID,
	}

	err := tuc.tr.UpdateThread(threadID, threadEntity)

	if err.Error() == "no affected rows" {
		return dto.ThreadResponse{}, domain.ErrBadParamInput
	} else if err != nil {
		return dto.ThreadResponse{}, domain.ErrInternalServerError
	}

	res, err := tuc.tr.GetThreadByID(threadID)

	if err != nil {
		return dto.ThreadResponse{}, domain.ErrInternalServerError
	}

	threadResponse := dto.ThreadResponse{
		ID:        res.ID,
		Title:     res.Title,
		Body:      res.Body,
		TopicID:   res.TopicID,
		ImageURL:  res.ImageURL,
		UserID:    res.UserID,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	return threadResponse, err
}
