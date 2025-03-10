package dto

import "github.com/google/uuid"

type CreateCommentRequest struct {
	ThreadId    uuid.UUID `json:"thread_id"`
	UserId      uuid.UUID `json:"user_id"`
	IsiKomentar string    `json:"isi_komentar" validate:"required"`
}

type UpdateCommentRequest struct {
	UserId      uuid.UUID `json:"user_id"`
	IsiKomentar string    `json:"isi_komentar" validate:"required"`
}

type CommentResponse struct {
	CommentId   uuid.UUID `json:"id"`
	ThreadId    uuid.UUID `json:"thread_id"`
	UserId      uuid.UUID `json:"user_id"`
	IsiKomentar string    `json:"isi_komentar" validate:"required"`
}
