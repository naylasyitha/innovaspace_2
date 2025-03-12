package dto

import (
	"github.com/google/uuid"
)

type MateriResponse struct {
	MateriId    uuid.UUID `json:"materi_id"`
	KelasId     string    `json:"kelas_id"`
	JenisMateri string    `json:"jenis_materi"`
	Judul       string    `json:"judul"`
	Deskripsi   string    `json:"deskripsi"`
	PathFile    string    `json:"path_file"`
}
