package usecase

import (
	"errors"
	"innovaspace/internal/app/comment/repository"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"

	"github.com/google/uuid"
)

type CommentUsecaseItf interface {
	CreateComment(input dto.CreateCommentRequest) (dto.CommentResponse, error)
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

func (u CommentUsecase) CreateComment(input dto.CreateCommentRequest) (dto.CommentResponse, error) {
	comment := entity.Comment{
		Id:          uuid.New(),
		ThreadId:    input.ThreadId,
		UserId:      input.UserId,
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

func (u CommentUsecase) UpdateComment(userId uuid.UUID, commentId uuid.UUID, input dto.UpdateCommentRequest) (dto.CommentResponse, error) {
	comment, err := u.commentRepo.GetCommentById(commentId)
	if err != nil {
		return dto.CommentResponse{}, err
	}

	if comment.UserId != userId {
		return dto.CommentResponse{}, errors.New("unauthorized")
	}

	if input.IsiKomentar != "" {
		comment.IsiKomentar = input.IsiKomentar
	} else {
		return dto.CommentResponse{}, errors.New("comment required")
	}

	err = u.commentRepo.UpdateComment(comment)
	if err != nil {
		return dto.CommentResponse{}, errors.New("failed to update comment")
	}

	return dto.CommentResponse{
		CommentId:   comment.Id,
		IsiKomentar: comment.IsiKomentar,
		UserId:      comment.UserId,
		ThreadId:    comment.ThreadId,
	}, nil
}

func (u CommentUsecase) DeleteComment(userId uuid.UUID, commentId uuid.UUID) error {
	comment, err := u.commentRepo.GetCommentById(commentId)
	if err != nil {
		return errors.New("comment not found")
	}

	if comment.UserId != userId {
		return errors.New("not allowed to delete comment")
	}

	return u.commentRepo.DeleteComment(commentId)
}
