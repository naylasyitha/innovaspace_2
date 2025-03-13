package dto

import (
	"github.com/google/uuid"
)

type Register struct {
	Email      string `json:"email" validate:"required,email"`
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required,min=8"`
	Nama       string `json:"nama" validate:"required"`
	Institusi  string `json:"institusi" validate:"required"`
	Preferensi string `json:"preferensi" validate:"required"`
}

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserParam struct {
	UserId   uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type ErrorInputResponse struct {
	FieldName string `json:"fieldName"`
	Message   string `json:"message"`
}

type GetProfile struct {
	Nama       string          `json:"nama"`
	Username   string          `json:"username"`
	Email      string          `json:"email"`
	Preferensi string          `json:"preferensi"`
	Institusi  string          `json:"institusi"`
	IsPremium  bool            `json:"is_premium"`
	MentorId   uuid.UUID       `json:"mentor_id"`
	Mentor     []ProfileMentor `json:"mentor"`
	Kelas      []ProfileKelas  `json:"kelas"`
}

type UpdateProfile struct {
	Nama       *string `json:"nama"`
	Username   *string `json:"username"`
	Email      *string `json:"email" validate:"email"`
	Preferensi *string `json:"preferensi"`
	Institusi  *string `json:"institusi"`
}

type SetMentor struct {
	MentorId uuid.UUID `json:"mentor_id"`
}

type ProfileMentor struct {
	Id           uuid.UUID `json:"id"`
	Nama         string    `json:"nama"`
	Deskripsi    string    `json:"deskripsi"`
	Preferensi   string    `json:"preferensi"`
	Spesialisasi string    `json:"spesialisasi"`
	Pendidikan   string    `json:"pendidikan"`
	Email        string    `json:"email"`
}

type ProfileKelas struct {
	KelasId          string `json:"kelas_id"`
	Nama             string `json:"nama"`
	Deskripsi        string `json:"deskripsi"`
	Kategori         string `json:"kategori"`
	JumlahMateri     int    `json:"jumlah_materi"`
	CoverCourse      string `json:"cover_course"`
	TingkatKesulitan string `json:"tingkat_kesulitan"`
	Durasi           int    `json:"durasi"`
	Persentase       int    `json:"persentase"`
}
