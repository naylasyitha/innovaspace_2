package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId       uuid.UUID `json:"user_id" gorm:"type:char(36);primaryKey"`
	MentorID     uuid.UUID `json:"mentor_id" gorm:"type:char(36);not null;index"`
	Email        string    `json:"email" gorm:"type:varchar(255);unique"`
	Username     string    `json:"username" gorm:"type:varchar(255);unique"`
	Password     string    `json:"password" gorm:"type:varchar(255)"`
	Nama         string    `json:"nama" gorm:"type:varchar(255)"`
	Institusi    string    `json:"institusi" gorm:"type:varchar(255)"`
	Preferensi   string    `json:"preferensi" gorm:"type:varchar(255)"`
	UserPict     string    `json:"user_pict" gorm:"type:varchar(255)"`
	IsPremium    bool      `json:"is_premium" gorm:"type:boolean;default:false"`
	PremiumStart time.Time `json:"premium_start" gorm:"type:timestamp"`
	PremiumEnd   time.Time `json:"premium_end"`
	HasMentor    bool      `json:"has_mentor" gorm:"type:boolean;default:false"`
	CreatedDate  time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`

	// Mentor *Mentor `gorm:"foreignKey:MentorID;references:MentorID"`
}
