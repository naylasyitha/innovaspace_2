package repository

import (
	"innovaspace/internal/domain/entity"

	"gorm.io/gorm"
)

type KelasMySQLItf interface {
	FindById(id string) (*entity.Kelas, error)
	GetAllKelas() ([]entity.Kelas, error)
}

type KelasMySQL struct {
	db *gorm.DB
}

func NewKelasMySQL(db *gorm.DB) KelasMySQLItf {
	return &KelasMySQL{db}
}

func (r *KelasMySQL) FindById(id string) (*entity.Kelas, error) {
	var kelas entity.Kelas
	err := r.db.Where("id = ?", id).First(&kelas).Error
	if err != nil {
		return nil, err

	}
	return &kelas, nil
}

func (r *KelasMySQL) GetAllKelas() ([]entity.Kelas, error) {
	var kelas []entity.Kelas
	err := r.db.Find(&kelas).Error
	if err != nil {
		return nil, err
	}
	return kelas, nil
}
