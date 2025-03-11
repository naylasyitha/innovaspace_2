package dto

import (
	"github.com/google/uuid"
)

type Kelas struct {
	KelasId          string `json:"kelas_id"`
	Nama             string `json:"nama"`
	Deskripsi        string `json:"deskripsi"`
	Kategori         string `json:"kategori"`
	JumlahMateri     int    `json:"jumlah_materi"`
	CoverCourse      string `json:"cover_course"`
	TingkatKesulitan string `json:"tingkat_kesulitan"`
	Durasi           int    `json:"durasi"`
}

type KelasDetailResponse struct {
	KelasId          string   `json:"kelas_id"`
	Nama             string   `json:"nama"`
	Deskripsi        string   `json:"deskripsi"`
	Kategori         string   `json:"kategori"`
	JumlahMateri     int      `json:"jumlah_materi"`
	CoverCourse      string   `json:"cover_course"`
	TingkatKesulitan string   `json:"tingkat_kesulitan"`
	Durasi           int      `json:"durasi"`
	Materi           []Materi `json:"materi"`
}

type Materi struct {
	MateriId  uuid.UUID `json:"materi_id"`
	KelasId   string    `json:"kelas_id"`
	Judul     string    `json:"judul"`
	Deskripsi string    `json:"deskripsi"`
	IsFree    bool      `json:"is_free"`
	PathFile  string    `json:"path_file"`
}
