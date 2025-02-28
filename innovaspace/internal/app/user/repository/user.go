package repository

import (
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"

	"gorm.io/gorm"
)

type UserMySQLItf interface {
	FindByEmail(email string) (*entity.User, error)
	Create(user *entity.User) error
	Get(user *entity.User, dto dto.UserParam) error
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
