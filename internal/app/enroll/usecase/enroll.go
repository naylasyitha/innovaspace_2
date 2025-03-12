package usecase

import (
	"errors"
	"innovaspace/internal/app/enroll/repository"
	KelasRepo "innovaspace/internal/app/kelas/repository"
	UserRepo "innovaspace/internal/app/user/repository"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"

	"github.com/google/uuid"
)

type EnrollUsecaseItf interface {
	CreateEnroll(input dto.EnrollRequest) (dto.EnrollResponse, error)
}

type EnrollUsecase struct {
	enrollRepo repository.EnrollMySQLItf
	userRepo   UserRepo.UserMySQLItf
	kelasRepo  KelasRepo.KelasMySQLItf
}

func NewEnrollUsecase(enrollRepo repository.EnrollMySQLItf, userRepo UserRepo.UserMySQLItf, kelasRepo KelasRepo.KelasMySQLItf) EnrollUsecaseItf {
	return &EnrollUsecase{
		enrollRepo: enrollRepo,
		userRepo:   userRepo,
		kelasRepo:  kelasRepo,
	}
}

func (u *EnrollUsecase) CreateEnroll(input dto.EnrollRequest) (dto.EnrollResponse, error) {
	user, err := u.userRepo.FindById(input.UserId)
	if err != nil {
		return dto.EnrollResponse{}, err
	}

	class, err := u.kelasRepo.FindById(input.KelasId)
	if err != nil {
		return dto.EnrollResponse{}, err
	}

	if class.IsPremium && !user.IsPremium {
		return dto.EnrollResponse{}, errors.New("user must be a premium member to enroll this class")
	}

	enroll := entity.Enroll{
		Id:      uuid.New(),
		KelasId: input.KelasId,
		UserId:  input.UserId,
	}

	err = u.enrollRepo.CreateEnroll(enroll)
	if err != nil {
		return dto.EnrollResponse{}, err
	}

	return dto.EnrollResponse{
		Id:      enroll.Id,
		KelasId: enroll.KelasId,
		UserId:  enroll.UserId,
	}, nil
}
