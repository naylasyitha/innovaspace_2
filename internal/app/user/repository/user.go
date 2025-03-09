package repository

import (
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserMySQLItf interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	Get(user *entity.User, dto dto.UserParam) error
	UpdateMentor(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindById(userId uuid.UUID) (*entity.User, error)
	FindByPreferensi(userId uuid.UUID) (*entity.User, error)
	FindByUsername(username string) (*entity.User, error)
}

type UserMySQL struct {
	db *gorm.DB
}

func NewUserMySQL(db *gorm.DB) UserMySQLItf {
	return &UserMySQL{db}
}

func (r *UserMySQL) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserMySQL) Update(user *entity.User) error {
	updates := map[string]interface{}{}
	if user.Nama != "" {
		updates["nama"] = user.Nama
	}
	if user.Username != "" {
		updates["username"] = user.Username
	}
	if user.Email != "" {
		updates["email"] = user.Email
	}
	if user.Preferensi != "" {
		updates["preferensi"] = user.Preferensi
	}
	if user.Institusi != "" {
		updates["institusi"] = user.Institusi
	}

	return r.db.Model(user).Where("id = ?", user.Id).Updates(updates).Error
}

func (r *UserMySQL) Get(user *entity.User, userParam dto.UserParam) error {
	return r.db.First(user, userParam).Error
}

func (r *UserMySQL) UpdateMentor(user *entity.User) error {
	return r.db.Model(user).Where("id = ?", user.Id).Updates(map[string]interface{}{
		"mentor_id": user.MentorId,
	}).Error
}

func (r *UserMySQL) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserMySQL) FindById(userId uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "id = ?", userId).Error
	return &user, err
}

func (r *UserMySQL) FindByPreferensi(userId uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserMySQL) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
