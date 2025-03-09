package entity

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Id           uuid.UUID `json:"comment_id" gorm:"type:char(36);primaryKey"`
	ThreadId     uuid.UUID `json:"thread_id" gorm:"type:char(36);not null;index"`
	UserId       uuid.UUID `json:"user_id" gorm:"type:char(36);not null;index"`
	IsiKomentar  string    `json:"isi_komentar" gorm:"type:text"`
	CreatedDate  time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`

	User   User   `json:"user" gorm:"foreignKey:UserId"`
	Thread Thread `json:"thread" gorm:"foreignKey:ThreadId"`
}
