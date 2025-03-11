package entity

import (
	"time"

	"github.com/google/uuid"
)

type Materi struct {
	Id           uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	KelasId      string    `json:"kelas_id" gorm:"type:char(36);not null;index"`
	UserId       uuid.UUID `json:"user_id" gorm:"type:char(36);not null;index"`
	Judul        string    `json:"judul" gorm:"type:varchar(255)"`
	Deskripsi    string    `json:"deskripsi" gorm:"type:text"`
	IsFree       bool      `json:"is_free" gorm:"type:bool"`
	PathFile     string    `json:"path_file" gorm:"type:text"`
	CreatedDate  time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`

	Kelas Kelas `json:"kelas" gorm:"foreignKey:KelasId"`
	User  User  `json:"user" gorm:"foreignKey:UserId"`
}
