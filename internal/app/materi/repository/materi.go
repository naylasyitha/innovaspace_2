package repository

import (
	"innovaspace/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MateriMySQLItf interface {
	FindById(id uuid.UUID) (*entity.Materi, error)
	GetMateriByKelasId(kelasId string) ([]entity.Materi, error)
}

type MateriMySQL struct {
	db *gorm.DB
}

func NewMateriMySQL(db *gorm.DB) MateriMySQLItf {
	return &MateriMySQL{db}
}

func (r *MateriMySQL) FindById(id uuid.UUID) (*entity.Materi, error) {
	var materi entity.Materi
	err := r.db.Where("id = ?", id).First(&materi).Error
	if err != nil {
		return nil, err
	}
	return &materi, nil
}

func (r *MateriMySQL) GetMateriByKelasId(kelasId string) ([]entity.Materi, error) {
	var materi []entity.Materi
	err := r.db.Where("kelas_id = ?", kelasId).Find(&materi).Error
	if err != nil {
		return nil, err
	}
	return materi, nil
}
