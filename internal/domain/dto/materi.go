package dto

import (
	"github.com/google/uuid"
)

type MateriResponse struct {
	MateriId  uuid.UUID `json:"materi_id"`
	KelasId   string    `json:"kelas_id"`
	Judul     string    `json:"judul"`
	Deskripsi string    `json:"deskripsi"`
	IsFree    bool      `json:"is_free"`
	PathFile  string    `json:"path_file"`
}
