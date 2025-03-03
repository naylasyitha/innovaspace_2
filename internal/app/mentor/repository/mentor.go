package repository

import (
	"innovaspace/internal/domain/entity"

	"gorm.io/gorm"
)

type MentorMySQLItf interface {
	FindByPreferensi(preferensi string) ([]entity.Mentor, error)
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

// func (r *MentorMySQL) GetMentorsByPreferensi() ([]entity.Mentor, error){
//     var mentors []entity.Mentor
//     err := r.db.Table("mentors").Select("mentors.mentor_id, mentors.nama, mentors.preferensi").Joins("LEFT JOIN users ON mentors.preferensi = users.preferensi").Scan(&mentors).Error

//     if err != nil {
//         return nil, err
//     }
//     return mentors, nil
// }

// func (r *MentorMySQL) GetMentorsByUserPreferensi(userPreferensi string) ([]entity.Mentor, error){
//     var mentors []entity.Mentor
//     err := r.db.Where("preferensi = ?", userPreferensi).Find(&mentors).Error
//     if err != nil {
//         return nil, err
//     }
//     return mentors, nil
// }
