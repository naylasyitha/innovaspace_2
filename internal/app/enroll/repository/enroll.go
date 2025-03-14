package repository

import (
	"innovaspace/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EnrollMySQLItf interface {
	CreateEnroll(enroll entity.Enroll) error
	FindByUserId(id uuid.UUID) ([]entity.Enroll, error)
}

type EnrollMySQL struct {
	db *gorm.DB
}

func NewEnrollMySQL(db *gorm.DB) EnrollMySQLItf {
	return &EnrollMySQL{db}
}

func (r *EnrollMySQL) CreateEnroll(enroll entity.Enroll) error {
	return r.db.Create(&enroll).Error
}

func (r *EnrollMySQL) FindByUserId(id uuid.UUID) ([]entity.Enroll, error) {
	var enrolls []entity.Enroll
	err := r.db.Model(entity.Enroll{}).Where("user_id = ?", id).Find(&enrolls).Error
	return enrolls, err
}
