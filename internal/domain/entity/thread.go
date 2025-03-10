package entity

import (
	"time"

	"github.com/google/uuid"
)

type Thread struct {
	Id           uuid.UUID `json:"thread_id" gorm:"type:char(36);primaryKey"`
	UserId       uuid.UUID `json:"user_id" gorm:"type:char(36);not null;index"`
	Kategori     string    `json:"kategori" gorm:"type:varchar(25)"`
	Isi          string    `json:"isi" gorm:"type:text"`
	CreatedDate  time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`

	User    User      `json:"user" gorm:"foreignKey:UserId"`
	Comment []Comment `json:"-" gorm:"foreignKey:ThreadId;constraint:OnDelete:CASCADE;"`
}
