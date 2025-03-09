package usecase

import (
	"errors"
	"innovaspace/internal/app/comment/repository"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"

	"github.com/google/uuid"
)

type CommentUsecaseItf interface {
	CreateComment(userId uuid.UUID, input dto.CreateCommentRequest) (dto.CommentResponse, error)
	GetCommentsByThread(threadId uuid.UUID) ([]dto.CommentResponse, error)
	UpdateComment(userId uuid.UUID, commentId uuid.UUID, input dto.UpdateCommentRequest) (dto.CommentResponse, error)
	DeleteComment(userId uuid.UUID, commentId uuid.UUID) error
}

type CommentUsecase struct {
	commentRepo repository.CommentMySQLItf
}

func NewCommentUsecase(commentRepo repository.CommentMySQLItf) CommentUsecaseItf {
	return &CommentUsecase{
		commentRepo: commentRepo,
	}
}

func (u CommentUsecase) CreateComment(userId uuid.UUID, input dto.CreateCommentRequest) (dto.CommentResponse, error) {
	comment := entity.Comment{
		Id:          uuid.New(),
		ThreadId:    input.ThreadId,
		UserId:      userId,
		IsiKomentar: input.IsiKomentar,
	}

	err := u.commentRepo.CreateComment(comment)
	if err != nil {
		return dto.CommentResponse{}, err
	}

	return dto.CommentResponse{
		CommentId:   comment.Id,
		ThreadId:    comment.ThreadId,
		UserId:      comment.UserId,
		IsiKomentar: comment.IsiKomentar,
	}, nil
}

func (u CommentUsecase) GetCommentsByThread(threadId uuid.UUID) ([]dto.CommentResponse, error) {
	comments, err := u.commentRepo.GetCommentsByThreadId(threadId)
	if err != nil {
		return nil, err
	}
	var response []dto.CommentResponse
	for _, comment := range comments {
		response = append(response, dto.CommentResponse{
			CommentId:   comment.Id,
			ThreadId:    comment.ThreadId,
			UserId:      comment.UserId,
			IsiKomentar: comment.IsiKomentar,
		})
	}

	return response, nil
}

func (u CommentUsecase) UpdateComment(userId uuid.UUID, commentId uuid.UUID, input dto.UpdateCommentRequest) (dto.CommentResponse, error) {
	comment, err := u.commentRepo.GetCommentById(commentId)
	if err != nil {
		return dto.CommentResponse{}, err
	}

	if comment.UserId != userId {
		return dto.CommentResponse{}, errors.New("unauthorized")
	}

	comment.IsiKomentar = input.IsiKomentar
	err = u.commentRepo.UpdateComment(comment)
	if err != nil {
		return dto.CommentResponse{}, err
	}

	return dto.CommentResponse{
		CommentId:   comment.Id,
		ThreadId:    comment.ThreadId,
		UserId:      comment.UserId,
		IsiKomentar: comment.IsiKomentar,
	}, nil
}

func (u CommentUsecase) DeleteComment(userId uuid.UUID, commentId uuid.UUID) error {
	comment, err := u.commentRepo.GetCommentById(commentId)
	if err != nil {
		return err
	}

	if comment.UserId != userId {
		return errors.New("unauthorized")
	}

	return u.commentRepo.DeleteComment(commentId)
}
