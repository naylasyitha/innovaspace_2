package repository

import (
	"innovaspace/internal/domain/entity"

	"gorm.io/gorm"
)

type ProgressMySQLItf interface {
	CreateProgress(progress entity.Progress) error
	// UpdateProgress(progress entity.Progress) error
	// GetProgressByUserAndMaterial(userId uuid.UUID, materiId uuid.UUID) (entity.Progress, error)
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

// func (r *ProgressMySQL) UpdateProgress(progress entity.Progress) error {
// 	return r.db.Save(&progress).Error
// }

// func (r *ProgressMySQL) GetProgressByUserAndMaterial(userId uuid.UUID, materiId uuid.UUID) (entity.Progress, error) {
// 	var progress entity.Progress
// 	err := r.db.Where("user_id = ? AND material_id = ?", userId, materiId).First(&progress).Error
// 	return progress, err
// }
