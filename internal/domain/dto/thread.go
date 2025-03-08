package dto

import "github.com/google/uuid"

type CreateThreadRequest struct {
	UserId   uuid.UUID `json:"user_id"`
	Kategori string    `json:"kategori"`
	Isi      string    `json:"isi"`
}

type UpdateThreadRequest struct {
	UserId   uuid.UUID `json:"user_id"`
	Kategori string    `json:"kategori"`
	Isi      string    `json:"isi"`
}

type ThreadResponse struct {
	ThreadId uuid.UUID `json:"thread_id"`
	UserId   uuid.UUID `json:"user_id"`
	Kategori string    `json:"kategori"`
	Isi      string    `json:"isi"`
}
