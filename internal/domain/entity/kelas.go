package entity

import (
	"time"
)

type Kelas struct {
	Id               string    `json:"id" gorm:"type:char(36);primaryKey"`
	Nama             string    `json:"nama" gorm:"type:varchar(255)"`
	Deskripsi        string    `json:"deskripsi" gorm:"type:text"`
	Kategori         string    `json:"kategori" gorm:"type:varchar(255)"`
	JumlahMateri     int       `json:"jumlah_materi" gorm:"type:int"`
	CoverCourse      string    `json:"cover_course" gorm:"type:text"`
	TingkatKesulitan string    `json:"tingkat_kesulitan"  gorm:"type:varchar(20)"`
	Durasi           int       `json:"durasi" gorm:"type:int"`
	IsPremium        bool      `json:"is_premium" gorm:"type:bool;default:false"`
	CreatedDate      time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate     time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`

	Materi      []Materi    `json:"-" gorm:"foreignKey:KelasId"`
	Enrollments []*Enroll   `json:"-" gorm:"foreignKey:KelasId"`
	Progress    []*Progress `json:"-" gorm:"foreignKey:KelasId"`
}
