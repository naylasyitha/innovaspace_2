package entity

import (
	"time"

	"github.com/google/uuid"
)

type Enroll struct {
	Id           uuid.UUID `json:"enroll_id" gorm:"type:char(36);primaryKey"`
	KelasId      string    `json:"kelas_id" gorm:"type:char(36);not null;index"`
	UserId       uuid.UUID `json:"user_id" gorm:"type:char(36);not null;index"`
	CreatedDate  time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`

	User  User  `json:"user" gorm:"foreignKey:UserId"`
	Kelas Kelas `json:"kelas" gorm:"foreignKey:KelasId"`
}
