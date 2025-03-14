package usecase

import (
	"errors"
	"fmt"
	EnrollRepo "innovaspace/internal/app/enroll/repository"
	KelasRepo "innovaspace/internal/app/kelas/repository"
	MentorRepo "innovaspace/internal/app/mentor/repository"
	ProgressRepo "innovaspace/internal/app/progress/repository"
	"innovaspace/internal/app/user/repository"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"
	"innovaspace/internal/infra/jwt"
	"innovaspace/internal/validation"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseItf interface {
	Register(register dto.Register) (entity.User, error)
	Login(login dto.Login) (string, error)
	UpdateUser(userId uuid.UUID, updateData dto.UpdateProfile) error
	GetProfileById(userId uuid.UUID) (dto.GetProfile, error)
	SetMentor(userId uuid.UUID, input dto.SetMentor) error
	UpdateMentor(userId uuid.UUID, input dto.SetMentor) error
}

type UserUsecase struct {
	userRepo     repository.UserMySQLItf
	mentorRepo   MentorRepo.MentorMySQLItf
	enrollRepo   EnrollRepo.EnrollMySQLItf
	progressRepo ProgressRepo.ProgressMySQLItf
	kelasRepo    KelasRepo.KelasMySQLItf
	jwt          jwt.JWT
	validator    validation.InputValidation
}

func NewUserUsecase(userRepo repository.UserMySQLItf, mentorRepo MentorRepo.MentorMySQLItf, enrollRepo EnrollRepo.EnrollMySQLItf, progressRepo ProgressRepo.ProgressMySQLItf, kelasRepo KelasRepo.KelasMySQLItf, validator validation.InputValidation, jwt jwt.JWT) UserUsecaseItf {
	return &UserUsecase{
		userRepo:     userRepo,
		mentorRepo:   mentorRepo,
		enrollRepo:   enrollRepo,
		progressRepo: progressRepo,
		kelasRepo:    kelasRepo,
		validator:    validator,
		jwt:          jwt,
	}
}

func (u *UserUsecase) Register(register dto.Register) (entity.User, error) {
	var user entity.User
	if err := u.validator.Validate(register); err != nil {
		return entity.User{}, err
	}

	if _, err := u.userRepo.FindByEmail(register.Email); err == nil {
		return entity.User{}, errors.New("email already exists")
	}

	if _, err := u.userRepo.FindByUsername(register.Username); err == nil {
		return entity.User{}, errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}

	user = entity.User{
		Id:         uuid.New(),
		Email:      register.Email,
		Username:   register.Username,
		Password:   string(hashedPassword),
		Nama:       register.Nama,
		Institusi:  register.Institusi,
		Preferensi: register.Preferensi,
	}

	err = u.userRepo.Create(&user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return entity.User{}, errors.New("duplicate entry")
		}
		return entity.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (u *UserUsecase) Login(login dto.Login) (string, error) {
	var user entity.User

	err := u.userRepo.Get(&user, dto.UserParam{Username: login.Username})
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := u.jwt.GenerateToken(user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserUsecase) UpdateUser(userId uuid.UUID, updateData dto.UpdateProfile) error {
	user, err := u.userRepo.FindById(userId)
	if err != nil {
		return errors.New("user not found")
	}

	if updateData.Nama != nil {
		user.Nama = *updateData.Nama
	}
	if updateData.Username != nil {
		user.Username = *updateData.Username
	}
	if updateData.Email != nil {
		user.Email = *updateData.Email
	}
	if updateData.Preferensi != nil {
		user.Preferensi = *updateData.Preferensi
	}
	if updateData.Institusi != nil {
		user.Institusi = *updateData.Institusi
	}

	err = u.userRepo.Update(user)
	if err != nil {
		return errors.New("failed to update user")
	}

	return nil
}

func (u *UserUsecase) GetProfileById(id uuid.UUID) (dto.GetProfile, error) {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return dto.GetProfile{}, err
	}

	var mentors []dto.ProfileMentor
	var mentorID *uuid.UUID

	if user.MentorId != nil && *user.MentorId != uuid.Nil {
		mentorID = user.MentorId
		mentor, err := u.mentorRepo.FindById(*user.MentorId)
		if err != nil {
			return dto.GetProfile{}, err
		}

		mentors = append(mentors, dto.ProfileMentor{
			Id:           mentor.Id,
			Nama:         mentor.Nama,
			Deskripsi:    mentor.Deskripsi,
			Preferensi:   mentor.Preferensi,
			Spesialisasi: mentor.Spesialisasi,
			Pendidikan:   mentor.Pendidikan,
			Email:        mentor.Email,
		})
	}

	enrollList, err := u.enrollRepo.FindByUserId(id)
	if err != nil {
		return dto.GetProfile{}, err
	}

	var resp []dto.ProfileKelas
	for _, enroll := range enrollList {
		kelas, err := u.kelasRepo.FindById(enroll.KelasId)
		if err != nil {
			return dto.GetProfile{}, err
		}

		progressList, err := u.progressRepo.GetProgressByUserAndKelas(id, enroll.KelasId)

		if err != nil {
			return dto.GetProfile{}, err
		}

		totalProgress := len(progressList)

		var persenProgress float64
		if kelas.JumlahMateri > 0 {
			persenProgress = (float64(totalProgress) / float64(kelas.JumlahMateri)) * 100
		} else {
			persenProgress = 0.0
		}

		resp = append(resp, dto.ProfileKelas{
			KelasId:          kelas.Id,
			Nama:             kelas.Nama,
			Deskripsi:        kelas.Deskripsi,
			Kategori:         kelas.Kategori,
			JumlahMateri:     kelas.JumlahMateri,
			CoverCourse:      kelas.CoverCourse,
			TingkatKesulitan: kelas.TingkatKesulitan,
			Durasi:           kelas.Durasi,
			Persentase:       int(persenProgress),
		})
	}

	return dto.GetProfile{
		Nama:       user.Nama,
		Username:   user.Username,
		Email:      user.Email,
		Preferensi: user.Preferensi,
		Institusi:  user.Institusi,
		MentorId:   getUUIDValue(mentorID),
		Mentor:     mentors,
		Kelas:      resp,
	}, nil
}

func (u UserUsecase) SetMentor(userId uuid.UUID, input dto.SetMentor) error {
	user, err := u.userRepo.FindById(userId)
	if err != nil {
		return errors.New("user not found")
	}

	if user.MentorId == nil {
		user.MentorId = &input.MentorId
		err = u.userRepo.UpdateMentor(user)
		if err != nil {
			return fmt.Errorf("failed to update mentor in user: %w", err)
		}
	} else {
		return errors.New("user has mentor")
	}

	return nil
}

func (u UserUsecase) UpdateMentor(userId uuid.UUID, input dto.SetMentor) error {
	user, err := u.userRepo.FindById(userId)
	if err != nil {
		return errors.New("user not found")
	}

	user.MentorId = &input.MentorId
	err = u.userRepo.UpdateMentor(user)
	if err != nil {
		return fmt.Errorf("failed to update mentor in user: %w", err)
	}

	return nil
}

func getUUIDValue(id *uuid.UUID) uuid.UUID {
	if id != nil {
		return *id
	}
	return uuid.Nil
}
