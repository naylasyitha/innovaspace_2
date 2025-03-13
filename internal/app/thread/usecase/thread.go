package usecase

import (
	"errors"
	CommentRepo "innovaspace/internal/app/comment/repository"
	"innovaspace/internal/app/thread/repository"
	UserRepo "innovaspace/internal/app/user/repository"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"
	"time"

	"github.com/google/uuid"
)

type ThreadUsecaseItf interface {
	CreateThread(input dto.CreateThreadRequest) (dto.ThreadResponse, error)
	GetAllThreads() ([]dto.ThreadResponse, error)
	UpdateThread(threadId uuid.UUID, userId uuid.UUID, input dto.UpdateThreadRequest) error
	DeleteThread(threadId uuid.UUID, userId uuid.UUID) error
	GetThreadDetails(threadId uuid.UUID) (*dto.ThreadDetailResponse, error)
}

type ThreadUsecase struct {
	threadRepo  repository.ThreadMySQLItf
	commentRepo CommentRepo.CommentMySQLItf
	userRepo    UserRepo.UserMySQLItf
}

func NewThreadUsecase(threadRepo repository.ThreadMySQLItf, commentRepo CommentRepo.CommentMySQLItf, userRepo UserRepo.UserMySQLItf) ThreadUsecaseItf {
	return &ThreadUsecase{
		threadRepo:  threadRepo,
		commentRepo: commentRepo,
		userRepo:    userRepo,
	}
}

func (u ThreadUsecase) CreateThread(input dto.CreateThreadRequest) (dto.ThreadResponse, error) {
	thread := entity.Thread{
		Id:       uuid.New(),
		UserId:   input.UserId,
		Kategori: input.Kategori,
		Isi:      input.Isi,
	}

	if err := u.threadRepo.CreateThread(thread); err != nil {
		return dto.ThreadResponse{}, err
	}

	return dto.ThreadResponse{
		ThreadId: thread.Id,
		UserId:   thread.UserId,
		Kategori: thread.Kategori,
		Isi:      thread.Isi,
	}, nil
}

func (u ThreadUsecase) GetAllThreads() ([]dto.ThreadResponse, error) {
	threads, err := u.threadRepo.GetAllThreads()
	if err != nil {
		return nil, err
	}

	var response []dto.ThreadResponse
	var date time.Time
	for _, thread := range threads {
		if !thread.ModifiedDate.IsZero() {
			date = thread.ModifiedDate
		} else {
			date = thread.CreatedDate
		}

		username, err := u.userRepo.FindById(thread.UserId)
		if err != nil {
			return nil, err
		}

		formattedDate := date.Format("02/01/2006 15:04")
		response = append(response, dto.ThreadResponse{
			ThreadId: thread.Id,
			UserId:   thread.UserId,
			Username: username.Username,
			Kategori: thread.Kategori,
			Isi:      thread.Isi,
			Tanggal:  formattedDate,
		})
	}

	return response, nil
}

func (u ThreadUsecase) UpdateThread(threadId uuid.UUID, userId uuid.UUID, input dto.UpdateThreadRequest) error {
	thread, err := u.threadRepo.GetThreadById(threadId)
	if err != nil {
		return errors.New("thread not found")
	}

	if thread.UserId != userId {
		return errors.New("unauthorized")
	}

	if input.Kategori != "" {
		thread.Kategori = input.Kategori
	}
	if input.Isi != "" {
		thread.Isi = input.Isi
	}

	err = u.threadRepo.UpdateThread(&thread)
	if err != nil {
		return errors.New("failed to update thread")
	}

	return nil
}

func (u ThreadUsecase) DeleteThread(threadId uuid.UUID, userId uuid.UUID) error {
	thread, err := u.threadRepo.GetThreadById(threadId)
	if err != nil {
		return errors.New("thread not found")
	}

	if thread.UserId != userId {
		return errors.New("not allowed to delete thread")
	}

	if err := u.commentRepo.DeleteCommentsByThreadId(threadId); err != nil {
		return errors.New("failed to delete comment: " + err.Error())
	}

	return u.threadRepo.DeleteThread(threadId)
}

func (u *ThreadUsecase) GetThreadDetails(threadId uuid.UUID) (*dto.ThreadDetailResponse, error) {
	thread, err := u.threadRepo.GetThreadById(threadId)
	if err != nil {
		return nil, err
	}

	comments, err := u.commentRepo.GetCommentsByThreadId(threadId)
	if err != nil {
		return nil, err
	}

	var commentResponses []dto.Comment
	var dateComment time.Time
	for _, comment := range comments {
		if !comment.ModifiedDate.IsZero() {
			dateComment = comment.ModifiedDate
		} else {
			dateComment = comment.CreatedDate
		}
		formattedDate := dateComment.Format("02/01/2006 15:04")

		username, err := u.userRepo.FindById(comment.UserId)
		if err != nil {
			return nil, err
		}
		commentResponses = append(commentResponses, dto.Comment{
			CommentId:   comment.Id,
			ThreadId:    comment.ThreadId,
			UserId:      comment.UserId,
			Username:    username.Username,
			IsiKomentar: comment.IsiKomentar,
			Tanggal:     formattedDate,
		})
	}

	var dateThread time.Time
	if !thread.ModifiedDate.IsZero() {
		dateThread = thread.ModifiedDate
	} else {
		dateThread = thread.CreatedDate
	}

	formattedDate := dateThread.Format("02/01/2006 15:04")

	username, err := u.userRepo.FindById(thread.UserId)
	if err != nil {
		return nil, err
	}
	response := &dto.ThreadDetailResponse{
		ThreadId: thread.Id,
		UserId:   thread.UserId,
		Username: username.Username,
		Kategori: thread.Kategori,
		Isi:      thread.Isi,
		Tanggal:  formattedDate,
		Comments: commentResponses,
	}

	return response, nil
}
