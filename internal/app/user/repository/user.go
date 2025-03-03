package repository

import (
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserMySQLItf interface {
	FindByEmail(email string) (*entity.User, error)
	FindById(userId uuid.UUID) (*entity.User, error)
	Create(user *entity.User) error
	Get(user *entity.User, dto dto.UserParam) error
	FindByPreferensi(userId uuid.UUID) (*entity.User, error)
}

type UserMySQL struct {
	db *gorm.DB
}

func NewUserMySQL(db *gorm.DB) UserMySQLItf {
	return &UserMySQL{db}
}

func (r *UserMySQL) Create(user *entity.User) error {
	return r.db.Debug().Create(user).Error
}

func (r *UserMySQL) Get(user *entity.User, userParam dto.UserParam) error {
	return r.db.Debug().First(user, userParam).Error
}

func (r *UserMySQL) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Debug().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserMySQL) FindById(userId uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "user_id = ?", userId).Error
	// err := r.db.Where("id = ?", userId).First(&user).Error
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
