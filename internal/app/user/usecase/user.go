package usecase

import (
	"errors"
	"innovaspace/internal/app/user/repository"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"
	"innovaspace/internal/infra/jwt"
	"innovaspace/internal/infra/storage"
	"innovaspace/internal/validation"
	"mime/multipart"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseItf interface {
	Register(register dto.Register, file *multipart.FileHeader) error
	Login(login dto.Login) (string, error)
	UpdateUser(userId uuid.UUID, updateData dto.GetProfile) error
	GetProfileByUsername(username string) (dto.GetProfile, error)
}

type UserUsecase struct {
	userRepo  repository.UserMySQLItf
	jwt       jwt.JWT
	validator validation.InputValidation
	storage   storage.StorageSupabase
}

func NewUserUsecase(userRepo repository.UserMySQLItf, validator validation.InputValidation, storage storage.StorageSupabase) UserUsecaseItf {
	return &UserUsecase{
		userRepo:  userRepo,
		validator: validator,
		storage:   storage,
	}
}

func (u *UserUsecase) Register(register dto.Register, file *multipart.FileHeader) error {
	var user entity.User
	if err := u.validator.Validate(register); err != nil {
		return err
	}

	if _, err := u.userRepo.FindByEmail(register.Email); err == nil {
		return errors.New("email already exists")
	}

	if _, err := u.userRepo.FindByUsername(register.Username); err == nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user = entity.User{
		UserId:     uuid.New(),
		Email:      register.Email,
		Username:   register.Username,
		Password:   string(hashedPassword),
		Nama:       register.Nama,
		Institusi:  register.Institusi,
		Preferensi: register.Preferensi,
	}

	// if file != nil {
	// 	imageUrl, err := storage.NewStorageSupabase().UploadProfilePicture(user.UserId.String(), file)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	user.UserPict = imageUrl
	// }

	err = u.userRepo.Create(&user)

	return err
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

	token, err := u.jwt.GenerateToken(user.UserId)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserUsecase) UpdateUser(userId uuid.UUID, updateData dto.GetProfile) error {
	user, err := u.userRepo.FindById(userId)
	if err != nil {
		return errors.New("user not found")
	}

	if updateData.Nama != "" {
		user.Nama = updateData.Nama
	}
	if updateData.Username != "" {
		user.Username = updateData.Username
	}
	if updateData.Email != "" {
		user.Email = updateData.Email
	}
	if updateData.Preferensi != "" {
		user.Preferensi = updateData.Preferensi
	}

	err = u.userRepo.Update(user)
	if err != nil {
		return errors.New("Failed to update user")
	}

	return nil
}

func (u *UserUsecase) GetProfileByUsername(username string) (dto.GetProfile, error) {
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return dto.GetProfile{}, err
	}

	return dto.GetProfile{
		Nama:       user.Nama,
		Username:   user.Username,
		Email:      user.Email,
		Preferensi: user.Preferensi,
		Institusi:  user.Institusi,
		// UserPict: user.UserPict,
	}, nil
}
