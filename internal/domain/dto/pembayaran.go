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
	PaymentUrl  string    `json:"payment_url"`
	CreatedDate string    `json:"created_date"`
}

type MidtransWebhookRequest struct {
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	TransactionID     string `json:"transaction_id"`
	StatusMessage     string `json:"status_message"`
	StatusCode        string `json:"status_code"`
	SignatureKey      string `json:"signature_key"`
	SettlementTime    string `json:"settlement_time"`
	PaymentType       string `json:"payment_type"`
	OrderID           string `json:"order_id"`
	MerchantID        string `json:"merchant_id"`
	GrossAmount       string `json:"gross_amount"`
	FraudStatus       string `json:"fraud_status"`
	Currency          string `json:"currency"`
}

type ProfilePembayaran struct {
	CreatedDate string    `json:"created_date"`
	OrderId     uuid.UUID `json:"order_id"`
	TipeBayar   string    `json:"tipe_bayar"`
	Total       int       `json:"total"`
	Status      string    `json:"status"`
}
