package dto

import "github.com/google/uuid"

type ProgressRequest struct {
	MateriId uuid.UUID `json:"materi_id"`
	UserId   uuid.UUID `json:"user_id"`
	Jawaban  string    `json:"jawaban"`
}

type ProgressResponse struct {
	Id          uuid.UUID `json:"progress_id"`
	MateriId    uuid.UUID `json:"materi_id"`
	UserId      uuid.UUID `json:"user_id"`
	Jawaban     string    `json:"jawaban"`
	IsCompleted bool      `json:"is_completed"`
}
