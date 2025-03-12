package repository

import (
	"innovaspace/internal/domain/entity"

	"gorm.io/gorm"
)

type EnrollMySQLItf interface {
	CreateEnroll(enroll entity.Enroll) error
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
