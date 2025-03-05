package repository

import (
	"innovaspace/internal/domain/entity"

	"gorm.io/gorm"
)

type MentorMySQLItf interface {
	FindByPreferensi(preferensi string) ([]entity.Mentor, error)
	FindByUsername(username string) (entity.Mentor, error)
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

func (r MentorMySQL) FindByUsername(username string) (entity.Mentor, error) {
	var mentors entity.Mentor
	err := r.db.Where("username = ?", username).Find(&mentors).Error
	return mentors, err
}

func (r MentorMySQL) GetAllMentors() ([]entity.Mentor, error) {
	var mentors []entity.Mentor
	err := r.db.Find(&mentors).Error
	if err != nil {
		return nil, err
	}
	return mentors, nil
}
