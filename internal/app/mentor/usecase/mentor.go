package usecase

import (
	"errors"
	"innovaspace/internal/app/mentor/repository"
	UserRepo "innovaspace/internal/app/user/repository"
	"innovaspace/internal/domain/dto"

	"github.com/google/uuid"
)

type MentorUsecaseItf interface {
	GetMentorDetails(id uuid.UUID) (dto.MentorsDetails, error)
	GetMentorsByPreferensi(preferensi string) ([]dto.Mentor, error)
	GetAllMentors() ([]dto.Mentor, error)
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

func (u MentorUsecase) GetMentorDetails(id uuid.UUID) (dto.MentorsDetails, error) {
	mentor, err := u.mentorRepo.FindById(id)
	if err != nil {
		return dto.MentorsDetails{}, errors.New("mentor not found")
	}

	return dto.MentorsDetails{
		ProfilMentor:    mentor.ProfilMentor,
		Nama:            mentor.Nama,
		Deskripsi:       mentor.Deskripsi,
		Spesialisasi:    mentor.Spesialisasi,
		Pendidikan:      mentor.Pendidikan,
		PengalamanKerja: string(mentor.PengalamanKerja),
		Pencapaian:      string(mentor.Pencapaian),
		Keahlian:        string(mentor.Keahlian),
		TopikAjar:       string(mentor.TopikAjar),
		Email:           mentor.Email,
	}, nil
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
			ProfilMentor: mentor.ProfilMentor,
		})
	}

	return response, nil
}

func (u MentorUsecase) GetAllMentors() ([]dto.Mentor, error) {
	mentors, err := u.mentorRepo.GetAllMentors()
	if err != nil {
		return nil, err
	}

	var response []dto.Mentor
	for _, mentor := range mentors {
		response = append(response, dto.Mentor{
			Nama:         mentor.Nama,
			Spesialisasi: mentor.Spesialisasi,
			Email:        mentor.Email,
			ProfilMentor: mentor.ProfilMentor,
		})
	}
	return response, nil
}
