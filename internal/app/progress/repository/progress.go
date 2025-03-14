package repository

import (
	"innovaspace/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProgressMySQLItf interface {
	CreateProgress(progress entity.Progress) error
	GetProgressByUserAndKelas(userId uuid.UUID, kelasId string) ([]entity.Progress, error)
}

type ProgressMySQL struct {
	db *gorm.DB
}

func NewProgressMySQL(db *gorm.DB) ProgressMySQLItf {
	return &ProgressMySQL{db}
}

func (r *ProgressMySQL) CreateProgress(progress entity.Progress) error {
	return r.db.Create(&progress).Error
}

func (r *ProgressMySQL) GetProgressByUserAndKelas(userId uuid.UUID, kelasId string) ([]entity.Progress, error) {
	var progress []entity.Progress
	err := r.db.Where("user_id = ? AND kelas_id = ?", userId, kelasId).First(&progress).Error
	return progress, err
}
