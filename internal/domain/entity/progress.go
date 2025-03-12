package entity

import (
	"time"

	"github.com/google/uuid"
)

type Progress struct {
	Id           uuid.UUID `json:"progress_id" gorm:"type:char(36);primaryKey"`
	MateriId     uuid.UUID `json:"materi_id" gorm:"type:char(36);not null;index"`
	UserId       uuid.UUID `json:"user_id" gorm:"type:char(36);not null;index"`
	Jawaban      string    `json:"jawaban" gorm:"type:text"`
	IsCompleted  bool      `json:"is_completed" gorm:"type:boolean"`
	CreatedDate  time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`

	User   User   `json:"user" gorm:"foreignKey:UserId"`
	Materi Materi `json:"materi" gorm:"foreignKey:MateriId"`
}
