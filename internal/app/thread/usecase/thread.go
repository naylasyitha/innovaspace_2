package usecase

import (
	"errors"
	CommentRepo "innovaspace/internal/app/comment/repository"
	"innovaspace/internal/app/thread/repository"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"

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
}

func NewThreadUsecase(threadRepo repository.ThreadMySQLItf, commentRepo CommentRepo.CommentMySQLItf) ThreadUsecaseItf {
	return &ThreadUsecase{
		threadRepo:  threadRepo,
		commentRepo: commentRepo,
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
	for _, thread := range threads {
		response = append(response, dto.ThreadResponse{
			ThreadId: thread.Id,
			UserId:   thread.UserId,
			Kategori: thread.Kategori,
			Isi:      thread.Isi,
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
	for _, comment := range comments {
		commentResponses = append(commentResponses, dto.Comment{
			CommentId:   comment.Id,
			ThreadId:    comment.ThreadId,
			UserId:      comment.UserId,
			IsiKomentar: comment.IsiKomentar,
		})
	}

	response := &dto.ThreadDetailResponse{
		ThreadId: thread.Id,
		UserId:   thread.UserId,
		Kategori: thread.Kategori,
		Isi:      thread.Isi,
		Comments: commentResponses,
	}

	return response, nil
}
