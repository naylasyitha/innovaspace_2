package entity

import (
	"time"

	"github.com/google/uuid"
)

type Kelas struct {
	Id               string    `json:"id" gorm:"type:char(36);primaryKey"`
	UserId           uuid.UUID `json:"user_id" gorm:"type:char(36);not null;index"`
	Nama             string    `json:"nama" gorm:"type:varchar(255)"`
	Deskripsi        string    `json:"deskripsi" gorm:"type:text"`
	Kategori         string    `json:"kategori" gorm:"type:varchar(255)"`
	JumlahMateri     int       `json:"jumlah_materi" gorm:"type:int"`
	CoverCourse      string    `json:"cover_course" gorm:"type:text"`
	TingkatKesulitan string    `json:"tingkat_kesulitan"  gorm:"type:varchar(20)"`
	Durasi           int       `json:"durasi" gorm:"type:int"`
	CreatedDate      time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate     time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`

	User   User     `json:"user" gorm:"foreignKey:UserId"`
	Materi []Materi `json:"-" gorm:"foreignKey:KelasId"`
}
