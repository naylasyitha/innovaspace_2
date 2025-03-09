package repository

import (
	"innovaspace/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MentorMySQLItf interface {
	FindByPreferensi(preferensi string) ([]entity.Mentor, error)
	FindById(id uuid.UUID) (*entity.Mentor, error)
	GetAllMentors() ([]entity.Mentor, error)
}

type MentorMySQL struct {
	db *gorm.DB
}

func NewMentorMySQL(db *gorm.DB) MentorMySQLItf {
	return &MentorMySQL{db}
}

func (r MentorMySQL) FindByPreferensi(preferensi string) ([]entity.Mentor, error) {
	var mentors []entity.Mentor
	err := r.db.Where("preferensi = ?", preferensi).Find(&mentors).Error
	return mentors, err
}

func (r MentorMySQL) FindById(id uuid.UUID) (*entity.Mentor, error) {
	var mentor entity.Mentor
	if err := r.db.Where("id = ?", id).Find(&mentor).Error; err != nil {
		return nil, err
	}
	return &mentor, nil
}

func (r MentorMySQL) GetAllMentors() ([]entity.Mentor, error) {
	var mentors []entity.Mentor
	err := r.db.Find(&mentors).Error
	if err != nil {
		return nil, err
	}
	return mentors, nil
}
