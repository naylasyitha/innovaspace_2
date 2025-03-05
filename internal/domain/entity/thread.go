package entity

import (
	"time"

	"github.com/google/uuid"
)

type Thread struct {
	ThreadId     uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	UserId       uuid.UUID `json:"user_id" gorm:"type:uuid;foreignKey"`
	Kategori     string    `json:"kategori" gorm:"type:varchar(25)"`
	Isi          string    `json:"isi" gorm:"type:text"`
	CreatedDate  time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
