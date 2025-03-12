package dto

import "github.com/google/uuid"

type EnrollRequest struct {
	KelasId string    `json:"kelas_id"`
	UserId  uuid.UUID `json:"user_id"`
}

type EnrollResponse struct {
	Id      uuid.UUID `json:"enroll_id"`
	KelasId string    `json:"kelas_id"`
	UserId  uuid.UUID `json:"user_id"`
}
