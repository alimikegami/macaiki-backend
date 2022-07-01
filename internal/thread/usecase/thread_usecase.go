package usecase

import (
	"fmt"
	"macaiki/internal/thread"
	"macaiki/internal/thread/dto"
	"macaiki/internal/thread/entity"
	"macaiki/pkg/utils"
	"path/filepath"

	cloudstorage "macaiki/pkg/cloud_storage"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
)

type ThreadUseCaseImpl struct {
	tr    thread.ThreadRepository
	awsS3 *cloudstorage.S3
}

func AuthorizeThreadAccess(threadID uint, userID uint, tuc *ThreadUseCaseImpl) (bool, entity.Thread, error) {
	thread, err := tuc.tr.GetThreadByID(threadID)
	if err != nil {
		return false, entity.Thread{}, err
	}

	if thread.UserID != userID {
		return false, entity.Thread{}, nil
	}

	return true, thread, nil
}

func CreateNewThreadUseCase(tr thread.ThreadRepository, awsS3Instance *cloudstorage.S3) thread.ThreadUseCase {
	return &ThreadUseCaseImpl{tr: tr, awsS3: awsS3Instance}
}

func (tuc *ThreadUseCaseImpl) GetThreadByID(threadID uint) (dto.ThreadResponse, error) {
	var thread dto.ThreadResponse
	res, err := tuc.tr.GetThreadByID(threadID)

	if err != nil {
		return dto.ThreadResponse{}, utils.ErrInternalServerError
	}

	thread = dto.ThreadResponse{
		ID:          res.ID,
		Title:       res.Title,
		Body:        res.Body,
		CommunityID: res.CommunityID,
		ImageURL:    res.ImageURL,
		UserID:      res.UserID,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}

	return thread, nil
}

func (tuc *ThreadUseCaseImpl) CreateThread(thread dto.ThreadRequest, userID uint) (dto.ThreadResponse, error) {
	threadEntity := entity.Thread{
		Title:       thread.Title,
		Body:        thread.Body,
		UserID:      userID,
		CommunityID: thread.CommunityID,
	}

	res, err := tuc.tr.CreateThread(threadEntity)
	if err != nil {
		return dto.ThreadResponse{}, err
	}
	return dto.ThreadResponse{
		ID:          res.ID,
		Title:       res.Title,
		Body:        res.Body,
		CommunityID: res.CommunityID,
		ImageURL:    res.ImageURL,
		UserID:      res.UserID,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}, nil
}

func (tuc *ThreadUseCaseImpl) SetThreadImage(img *multipart.FileHeader, threadID uint, userID uint) error {
	flag, thread, err := AuthorizeThreadAccess(threadID, userID, tuc)
	if err != nil {
		return err
	}

	if !flag {
		return utils.ErrUnauthorizedAccess
	}

	if thread.ImageURL != "" {
		err = tuc.awsS3.DeleteImage(thread.ImageURL, "thread")
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	uniqueFilename := uuid.New()
	result, err := tuc.awsS3.UploadImage(uniqueFilename.String(), "thread", img)
	if err != nil {
		fmt.Printf("failed to upload file, %v", err)
		return err
	}

	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))

	err = tuc.tr.SetThreadImage(uniqueFilename.String()+filepath.Ext(img.Filename), threadID)

	return err
}

func (tuc *ThreadUseCaseImpl) DeleteThread(threadID uint, userID uint) error {
	flag, _, err := AuthorizeThreadAccess(threadID, userID, tuc)
	if err != nil {
		return err
	}

	if !flag {
		return utils.ErrUnauthorizedAccess
	}

	err = tuc.tr.DeleteThread(threadID)
	return err
}

func (tuc *ThreadUseCaseImpl) UpdateThread(thread dto.ThreadRequest, threadID uint, userID uint) (dto.ThreadResponse, error) {
	flag, _, err := AuthorizeThreadAccess(threadID, userID, tuc)
	if err != nil {
		return dto.ThreadResponse{}, err
	}

	if !flag {
		return dto.ThreadResponse{}, utils.ErrUnauthorizedAccess
	}

	threadEntity := entity.Thread{
		Title:       thread.Title,
		Body:        thread.Body,
		CommunityID: thread.CommunityID,
	}

	err = tuc.tr.UpdateThread(threadID, threadEntity)
	if err != nil {
		if err.Error() == "no affected rows" {
			return dto.ThreadResponse{}, utils.ErrBadParamInput
		}
		return dto.ThreadResponse{}, utils.ErrInternalServerError
	}

	res, err := tuc.tr.GetThreadByID(threadID)

	if err != nil {
		return dto.ThreadResponse{}, utils.ErrInternalServerError
	}

	threadResponse := dto.ThreadResponse{
		ID:          res.ID,
		Title:       res.Title,
		Body:        res.Body,
		CommunityID: res.CommunityID,
		ImageURL:    res.ImageURL,
		UserID:      res.UserID,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}

	return threadResponse, err
}

func (tuc *ThreadUseCaseImpl) LikeThread(threadID uint, userID uint) error {
	threadLikes := entity.ThreadLikes{
		ThreadID: threadID,
		UserID:   userID,
	}
	err := tuc.tr.LikeThread(threadLikes)

	return err
}

func (tuc *ThreadUseCaseImpl) GetTrendingThreads(userID uint) ([]dto.DetailedThreadResponse, error) {
	var threads []dto.DetailedThreadResponse
	res, err := tuc.tr.GetTrendingThreads(userID)

	if err != nil {
		return []dto.DetailedThreadResponse{}, utils.ErrInternalServerError
	}

	for _, thread := range res {
		threads = append(threads, dto.DetailedThreadResponse{
			ID:                    thread.Thread.ID,
			Title:                 thread.Title,
			Body:                  thread.Body,
			CommunityID:           thread.CommunityID,
			ImageURL:              thread.ImageURL,
			UserID:                thread.UserID,
			UserName:              thread.User.Name,
			UserProfession:        thread.User.Profession,
			UserProfilePictureURL: thread.User.ProfileImageUrl,
			CreatedAt:             thread.Thread.CreatedAt,
			UpdatedAt:             thread.Thread.UpdatedAt,
			LikesCount:            thread.LikesCount,
			IsLiked:               thread.IsLiked,
			IsFollowed:            thread.IsFollowed,
		})
	}

	return threads, nil
}

func (tuc *ThreadUseCaseImpl) GetThreadsFromFollowedCommunity(userID uint) ([]dto.DetailedThreadResponse, error) {
	var threads []dto.DetailedThreadResponse
	res, err := tuc.tr.GetThreadsFromFollowedCommunity(userID)

	if err != nil {
		return []dto.DetailedThreadResponse{}, utils.ErrInternalServerError
	}

	for _, thread := range res {
		threads = append(threads, dto.DetailedThreadResponse{
			ID:                    thread.Thread.ID,
			Title:                 thread.Title,
			Body:                  thread.Body,
			CommunityID:           thread.CommunityID,
			ImageURL:              thread.ImageURL,
			UserID:                thread.UserID,
			UserName:              thread.User.Name,
			UserProfession:        thread.User.Profession,
			UserProfilePictureURL: thread.User.ProfileImageUrl,
			CreatedAt:             thread.Thread.CreatedAt,
			UpdatedAt:             thread.Thread.UpdatedAt,
			LikesCount:            thread.LikesCount,
			IsLiked:               thread.IsLiked,
			IsFollowed:            thread.IsFollowed,
		})
	}

	return threads, nil
}

func (tuc *ThreadUseCaseImpl) GetThreadsFromFollowedUsers(userID uint) ([]dto.DetailedThreadResponse, error) {
	var threads []dto.DetailedThreadResponse
	res, err := tuc.tr.GetThreadsFromFollowedUsers(userID)

	if err != nil {
		return []dto.DetailedThreadResponse{}, utils.ErrInternalServerError
	}

	for _, thread := range res {
		threads = append(threads, dto.DetailedThreadResponse{
			ID:                    thread.Thread.ID,
			Title:                 thread.Title,
			Body:                  thread.Body,
			CommunityID:           thread.CommunityID,
			ImageURL:              thread.ImageURL,
			UserID:                thread.UserID,
			UserName:              thread.User.Name,
			UserProfession:        thread.User.Profession,
			UserProfilePictureURL: thread.User.ProfileImageUrl,
			CreatedAt:             thread.Thread.CreatedAt,
			UpdatedAt:             thread.Thread.UpdatedAt,
			LikesCount:            thread.LikesCount,
			IsLiked:               thread.IsLiked,
			IsFollowed:            thread.IsFollowed,
		})
	}

	return threads, nil
}

func (tuc *ThreadUseCaseImpl) AddThreadComment(comment dto.CommentRequest) error {
	err := tuc.tr.AddThreadComment(entity.Comment{
		Body:      comment.Body,
		UserID:    comment.UserID,
		ThreadID:  comment.ThreadID,
		CommentID: comment.CommentID,
	})

	return err
}

func (tuc *ThreadUseCaseImpl) GetCommentsByThreadID(threadID uint) ([]dto.CommentResponse, error) {
	var commentsResp []dto.CommentResponse

	comments, err := tuc.tr.GetCommentsByThreadID(threadID)

	if err != nil {
		return []dto.CommentResponse{}, err
	}

	for _, comment := range comments {
		commentsResp = append(commentsResp, dto.CommentResponse{
			ID:                    comment.Comment.ID,
			Body:                  comment.Body,
			ThreadID:              comment.ThreadID,
			UserID:                comment.UserID,
			Username:              comment.User.Name,
			UserProfilePictureURL: comment.User.ProfileImageUrl,
			CreatedAt:             comment.Comment.CreatedAt,
			LikesCount:            comment.LikesCount,
		})
	}

	return commentsResp, nil
}

func (tuc *ThreadUseCaseImpl) GetThreads(keyword string, userID uint) ([]dto.DetailedThreadResponse, error) {
	var threads []dto.DetailedThreadResponse
	res, err := tuc.tr.GetThreads(keyword, userID)

	if err != nil {
		return []dto.DetailedThreadResponse{}, utils.ErrInternalServerError
	}

	for _, thread := range res {
		threads = append(threads, dto.DetailedThreadResponse{
			ID:                    thread.Thread.ID,
			Title:                 thread.Title,
			Body:                  thread.Body,
			CommunityID:           thread.CommunityID,
			ImageURL:              thread.ImageURL,
			UserID:                thread.UserID,
			UserName:              thread.User.Name,
			UserProfession:        thread.User.Profession,
			UserProfilePictureURL: thread.User.ProfileImageUrl,
			CreatedAt:             thread.Thread.CreatedAt,
			UpdatedAt:             thread.Thread.UpdatedAt,
			LikesCount:            thread.LikesCount,
			IsLiked:               thread.IsLiked,
			IsFollowed:            thread.IsFollowed,
		})
	}

	return threads, nil
}

func (tuc *ThreadUseCaseImpl) LikeComment(commentID, userID uint) error {
	err := tuc.tr.LikeComment(entity.CommentLikes{
		UserID:    userID,
		CommentID: commentID,
	})

	return err
}
