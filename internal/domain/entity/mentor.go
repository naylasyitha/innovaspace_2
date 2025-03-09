package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Mentor struct {
	Id              uuid.UUID      `json:"mentor_id" gorm:"type:char(36);primaryKey"`
	Email           string         `json:"email" gorm:"type:varchar(255);unique"`
	Username        string         `json:"username" gorm:"type:varchar(255);unique"`
	Password        string         `json:"password" gorm:"type:varchar(255)"`
	Nama            string         `json:"nama" gorm:"type:varchar(255)"`
	Deskripsi       string         `json:"deskripsi" gorm:"type:text"`
	Pendidikan      string         `json:"pendidikan" gorm:"type:varchar(255)"`
	Preferensi      string         `json:"preferensi" gorm:"type:varchar(255)"`
	PengalamanKerja datatypes.JSON `json:"pengalaman_kerja" gorm:"type:json"`
	Pencapaian      datatypes.JSON `json:"pencapaian" gorm:"type:json"`
	Keahlian        datatypes.JSON `json:"keahlian" gorm:"type:json"`
	TopikAjar       datatypes.JSON `json:"topik_ajar" gorm:"type:json"`
	Spesialisasi    string         `json:"spesialisasi" gorm:"type:varchar(255)"`
	ProfilMentor    string         `json:"profil_mentor" gorm:"type:text"`
	CreatedDate     time.Time      `json:"created_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	ModifiedDate    time.Time      `json:"modified_date" gorm:"type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
