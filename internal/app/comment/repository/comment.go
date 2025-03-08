package repository

import (
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
	err := r.db.Preload("Thread").First(&comment, "comment_id = ?", commentId).Error
	return comment, err
}

func (r *CommentMySQL) GetCommentsByThreadId(threadId uuid.UUID) ([]entity.Comment, error) {
	var comments []entity.Comment
	err := r.db.Where("thread_id = ?", threadId).Find(&comments).Error
	return comments, err
}

func (r *CommentMySQL) UpdateComment(comment entity.Comment) error {
	return r.db.Save(&comment).Error
}

func (r *CommentMySQL) DeleteComment(commentId uuid.UUID) error {
	return r.db.Where("comment_id = ?", commentId).Delete(&entity.Comment{}).Error
}
