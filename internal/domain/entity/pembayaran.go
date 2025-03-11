package entity

import (
	"time"

	"github.com/google/uuid"
)

type Pembayaran struct {
	Id        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	UserId    uuid.UUID `json:"user_id" gorm:"type:char(36);not null;index"`
	OrderId   string    `json:"order_id" gorm:"type:varchar(255);not null;unique"`
	Total     int       `json:"total" gorm:"type:varchar(25);not null"`
	Status    string    `json:"status" gorm:"type:varchar(50);not null"`
	TipeBayar string    `json:"tipe_bayar" gorm:"type:varchar(50)"`
	// SnapUrl      string    `json:"snap_url" gorm:"type:text"`
	Token        string    `json:"token" gorm:"type:varchar(255);not null"`
	CreatedDate  time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
