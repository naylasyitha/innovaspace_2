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
	Nama       string `json:"nama" validate:"required"`
	Username   string `json:"username" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Preferensi string `json:"preferensi" validate:"required"`
	Institusi  string `json:"institusi" validate:"required"`
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
