package usecase

import (
	"fmt"
	"macaiki/internal/domain"
	"macaiki/internal/thread/dto"

	cloudstorage "macaiki/pkg/cloud_storage"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
)

type ThreadUseCaseImpl struct {
	tr    domain.ThreadRepository
	awsS3 *cloudstorage.S3
}

func CreateNewThreadUseCase(tr domain.ThreadRepository, awsS3Instance *cloudstorage.S3) domain.ThreadUseCase {
	return &ThreadUseCaseImpl{tr: tr, awsS3: awsS3Instance}
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

func (tuc *ThreadUseCaseImpl) GetThreadByID(threadID uint) (dto.ThreadResponse, error) {
	var thread dto.ThreadResponse
	res, err := tuc.tr.GetThreadByID(threadID)

	if err != nil {
		return dto.ThreadResponse{}, domain.ErrInternalServerError
	}

	thread = dto.ThreadResponse{
		ID:        res.ID,
		Title:     res.Title,
		Body:      res.Body,
		TopicID:   res.TopicID,
		ImageURL:  res.ImageURL,
		UserID:    res.UserID,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	return thread, nil
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

func (tuc *ThreadUseCaseImpl) SetThreadImage(img *multipart.FileHeader, threadID uint) error {
	uniqueFilename := uuid.New()
	result, err := tuc.awsS3.UploadImage(uniqueFilename.String(), "thread", img)
	if err != nil {
		fmt.Printf("failed to upload file, %v", err)
		return err
	}

	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))

	err = tuc.tr.SetThreadImage(aws.StringValue(&result.Location), threadID)

	return err
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
