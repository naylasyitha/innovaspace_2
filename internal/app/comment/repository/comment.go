package repository

import (
	"fmt"
	"innovaspace/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentMySQLItf interface {
	CreateComment(comment entity.Comment) error
	GetCommentById(commentId uuid.UUID) (entity.Comment, error)
	GetCommentsByThreadId(threadId uuid.UUID) ([]entity.Comment, error)
	UpdateComment(comment entity.Comment) error
	DeleteComment(commentId uuid.UUID) error
}

type CommentMySQL struct {
	db *gorm.DB
}

func NewCommentMySQL(db *gorm.DB) CommentMySQLItf {
	return &CommentMySQL{db}
}

func (r *CommentMySQL) CreateComment(comment entity.Comment) error {
	return r.db.Create(&comment).Error
}

func (r *CommentMySQL) GetCommentById(commentId uuid.UUID) (entity.Comment, error) {
	var comment entity.Comment
	err := r.db.Preload("Thread").First(&comment, "id = ?", commentId).Error
	return comment, err
}

func (r *CommentMySQL) GetCommentsByThreadId(threadId uuid.UUID) ([]entity.Comment, error) {
	var comments []entity.Comment
	err := r.db.Where("thread_id = ?", threadId).Find(&comments).Error
	return comments, err
}

func (r *CommentMySQL) UpdateComment(comment entity.Comment) error {
	updates := map[string]interface{}{}
	if comment.IsiKomentar != "" {
		updates["isi_komentar"] = comment.IsiKomentar
	}
	fmt.Println("Comment ID:", comment.Id)
	return r.db.Debug().Model(&entity.Comment{}).Where("id = ?", comment.Id).Updates(updates).Error
}

func (r *CommentMySQL) DeleteComment(commentId uuid.UUID) error {
	return r.db.Where("id = ?", commentId).Delete(&entity.Comment{}).Error
}

func (r *CommentMySQL) DeleteCommentsByThreadId(threadId uuid.UUID) error {
	return r.db.Where("thread_id = ?", threadId).Delete(&entity.Comment{}).Error
}
