package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateThreadRequest struct {
	UserId   uuid.UUID `json:"user_id"`
	Kategori string    `json:"kategori"`
	Isi      string    `json:"isi"`
}

type UpdateThreadRequest struct {
	Kategori string `json:"kategori"`
	Isi      string `json:"isi"`
}

type ThreadResponse struct {
	ThreadId uuid.UUID `json:"thread_id"`
	UserId   uuid.UUID `json:"user_id"`
	Kategori string    `json:"kategori"`
	Isi      string    `json:"isi"`
	Tanggal  string    `json:"tanggal"`
}

type ThreadDetailResponse struct {
	ThreadId uuid.UUID `json:"thread_id"`
	UserId   uuid.UUID `json:"user_id"`
	Kategori string    `json:"kategori"`
	Isi      string    `json:"isi"`
	Tanggal  time.Time `json:"tanggal"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	CommentId   uuid.UUID `json:"comment_id"`
	ThreadId    uuid.UUID `json:"thread_id"`
	UserId      uuid.UUID `json:"user_id"`
	IsiKomentar string    `json:"isi_komentar"`
	Tanggal     string    `json:"tanggal"`
}
