package dto

import (
	"github.com/google/uuid"
)

type Register struct {
	Email        string `json:"email" validate:"required,email"`
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required,min=8"`
	Nama         string `json:"nama" validate:"required"`
	Institusi    string `json:"institusi" validate:"required"`
	BidangBisnis string `json:"bidang_bisnis" validate:"required"`
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
