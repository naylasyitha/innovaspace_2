package entity

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	Id               uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Nama             string    `json:"nama" gorm:"type:varchar(255)"`
	Deskripsi        string    `json:"deskripsi" gorm:"type:text"`
	Kategori         string    `json:"kategori" gorm:"type:varchar(255)"`
	JumlahMateri     int
	JumlahAkses      int
	CoverCourse      string
	TingkatKesulitan string
	Durasi           string
	CreatedDate      time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate     time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`

	// Materi []Materi
	// Progres []Progres
}
