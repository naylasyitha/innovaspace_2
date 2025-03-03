package dto

type Mentor struct {
	Email        string `json:"email" validate:"required,email"`
	Nama         string `json:"nama" validate:"required"`
	Spesialisasi string `json:"spesialisasi" validate:"required"`
}
