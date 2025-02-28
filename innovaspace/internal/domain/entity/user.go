package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId       uuid.UUID `json:"id" gorm:"type:varchar(36);primaryKey"`
	Email        string    `json:"email" gorm:"type:varchar(255);unique"`
	Username     string    `json:"username" gorm:"type:varchar(255);unique"`
	Password     string    `json:"password" gorm:"type:varchar(255)"`
	Nama         string    `json:"nama" gorm:"type:varchar(255)"`
	Institusi    string    `json:"institusi" gorm:"type:varchar(255)"`
	BidangBisnis string    `json:"bidang_bisnis" gorm:"type:varchar(255)"`
	Preferensi   string    `json:"preferensi" gorm:"type:varchar(255)"`
	IsPremium    bool      `json:"is_premium" gorm:"type:boolean"`
	PremiumStart time.Time `json:"premium_start" gorm:"type:timestamp"`
	PremiumEnd   time.Time `json:"premium_end"`
	CreatedBy    string    `json:"created_by" gorm:"type:varchar(255)"`
	CreatedDate  time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedBy   string    `json:"modified_by" gorm:"type:varchar(255)"`
	ModifiedDate time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
