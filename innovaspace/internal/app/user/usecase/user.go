package usecase

import (
	"errors"
	"innovaspace/internal/app/user/repository"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"
	"innovaspace/internal/infra/jwt"
	"innovaspace/internal/validation"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseItf interface {
	Register(register dto.Register) error
	Login(login dto.Login) (string, error)
}

type UserUsecase struct {
	userRepo  repository.UserMySQLItf
	jwt       jwt.JWT
	validator validation.InputValidation
}

func NewUserUsecase(userRepo repository.UserMySQLItf, validator validation.InputValidation) UserUsecaseItf {
	return &UserUsecase{
		userRepo:  userRepo,
		validator: validator,
	}
}

func (u *UserUsecase) Register(register dto.Register) error {
	var user entity.User
	if err := u.validator.Validate(register); err != nil {
		return err
	}

	if _, err := u.userRepo.FindByEmail(register.Email); err == nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user = entity.User{
		Email:        register.Email,
		Username:     register.Username,
		Password:     string(hashedPassword),
		Nama:         register.Nama,
		Institusi:    register.Institusi,
		BidangBisnis: register.BidangBisnis,
		Preferensi:   register.Preferensi,
	}

	err = u.userRepo.Create(&user)

	return err
}

func (u *UserUsecase) Login(login dto.Login) (string, error) {
	var user entity.User

	err := u.userRepo.Get(&user, dto.UserParam{Username: login.Username})
	if err != nil {
		return "", errors.New("invalid email or username")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := u.jwt.GenerateToken(user.UserId)
	if err != nil {
		return "", err
	}

	return token, nil
}
