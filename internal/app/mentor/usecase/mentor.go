package usecase

import (
	"innovaspace/internal/app/mentor/repository"
	UserRepo "innovaspace/internal/app/user/repository"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"
	"log"

	"github.com/google/uuid"
)

type MentorUsecaseItf interface {
	GetMentorsDetails(userId uuid.UUID) ([]entity.Mentor, error)
	GetMentorsByPreferensi(preferensi string) ([]dto.Mentor, error)
}

type MentorUsecase struct {
	mentorRepo repository.MentorMySQLItf
	userRepo   UserRepo.UserMySQLItf
}

func NewMentorUsecase(mentorRepo repository.MentorMySQLItf, userRepo UserRepo.UserMySQLItf) MentorUsecaseItf {
	return &MentorUsecase{
		mentorRepo: mentorRepo,
		userRepo:   userRepo,
	}
}

func (u MentorUsecase) GetMentorsDetails(userId uuid.UUID) ([]entity.Mentor, error) {
	user, err := u.userRepo.FindById(userId)
	if err != nil {
		log.Println("User not found:", err)
		return nil, err
	}

	mentors, err := u.mentorRepo.FindByPreferensi(user.Preferensi)
	if err != nil {
		log.Println("Error fetching mentors:", err)
		return nil, err
	}

	return mentors, nil
}

func (u MentorUsecase) GetMentorsByPreferensi(preferensi string) ([]dto.Mentor, error) {
	mentors, err := u.mentorRepo.FindByPreferensi(preferensi)
	if err != nil {
		return nil, err
	}

	var response []dto.Mentor
	for _, mentor := range mentors {
		response = append(response, dto.Mentor{
			Nama:         mentor.Nama,
			Spesialisasi: mentor.Spesialisasi,
			Email:        mentor.Email,
		})
	}

	return response, nil
}
