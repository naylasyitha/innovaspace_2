package entity

import (
	"time"

	"github.com/google/uuid"
)

type Mentor struct {
	MentorID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Email           string    `json:"email" gorm:"type:varchar(255);unique"`
	Username        string    `json:"username" gorm:"type:varchar(255);unique"`
	Password        string    `json:"password" gorm:"type:varchar(255)"`
	Nama            string    `json:"nama" gorm:"type:varchar(255)"`
	Deskripsi       string    `json:"deskripsi" gorm:"type:text"`
	Pendidikan      string    `json:"pendidikan" gorm:"type:varchar(255)"`
	Preferensi      string    `json:"preferensi" gorm:"type:varchar(255)"`
	PengalamanKerja string    `json:"pengalaman_kerja" gorm:"type:text"`
	Pencapaian      string    `json:"pencapaian" gorm:"type:text"`
	Keahlian        string    `json:"keahlian" gorm:"type:text"`
	TopikAjar       string    `json:"topik_ajar" gorm:"type:text"`
	Spesialisasi    string    `json:"spesialisasi" gorm:"type:varchar(255)"`
	CreatedDate     time.Time `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate    time.Time `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
