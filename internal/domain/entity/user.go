package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Email        string    `json:"email" gorm:"type:varchar(255);unique"`
	Username     string    `json:"username" gorm:"type:varchar(255);unique"`
	Password     string    `json:"password" gorm:"type:varchar(255)"`
	Nama         string    `json:"nama" gorm:"type:varchar(255)"`
	Institusi    string    `json:"institusi" gorm:"type:varchar(255)"`
	Preferensi   string    `json:"preferensi" gorm:"type:varchar(255)"`
	IsPremium    bool      `json:"is_premium" gorm:"type:boolean;default:false"`
	PremiumStart time.Time `json:"premium_start" gorm:"type:timestamp"`
	PremiumEnd   time.Time `json:"premium_end"`
	CreatedDate  time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
