package dto

import "github.com/google/uuid"

type PaymentRequest struct {
	UserId    uuid.UUID `json:"user_id" validate:"required"`
	Total     int       `json:"total" validate:"required"`
	TipeBayar string    `json:"tipe_bayar" validate:"required"`
	Durasi    int       `json:"durasi" validate:"required"`
}

type PaymentResponse struct {
	Id          uuid.UUID `json:"id"`
	OrderID     string    `json:"order_id"`
	Total       int       `json:"total"`
	Status      string    `json:"status"`
	Token       string    `json:"token"`
	CreatedDate string    `json:"created_at"`
}
