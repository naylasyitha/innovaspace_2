package entity

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	CommentId    uuid.UUID `json:"comment_id" gorm:"type:char(36);primaryKey"`
	ThreadID     uuid.UUID `json:"thread_id" gorm:"type:char(36);not null;index"`
	UserID       uuid.UUID `json:"user_id" gorm:"type:char(36);not null;index"`
	IsiKomentar  string    `json:"isi_komentar" gorm:"type:text"`
	CreatedDate  time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`

	// User   *User   `gorm:"foreignKey:UserID;references:UserID"`
	// Thread *Thread `gorm:"foreignKey:ThreadID;references:ThreadID"`
}
