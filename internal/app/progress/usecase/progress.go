package usecase

import (
	"errors"
	MateriRepo "innovaspace/internal/app/materi/repository"
	"innovaspace/internal/app/progress/repository"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"

	"github.com/google/uuid"
)

type ProgressUsecaseItf interface {
	CreateProgress(input dto.ProgressRequest) (dto.ProgressResponse, error)
}

type ProgressUsecase struct {
	progressRepo repository.ProgressMySQLItf
	materiRepo   MateriRepo.MateriMySQLItf
}

func NewProgressUsecase(progressRepo repository.ProgressMySQLItf, materiRepo MateriRepo.MateriMySQLItf) ProgressUsecaseItf {
	return &ProgressUsecase{
		progressRepo: progressRepo,
		materiRepo:   materiRepo,
	}
}

func (u *ProgressUsecase) CreateProgress(input dto.ProgressRequest) (dto.ProgressResponse, error) {
	materi, err := u.materiRepo.FindById(input.MateriId)
	if err != nil {
		return dto.ProgressResponse{}, err
	}

	if materi.JenisMateri == "Study Case" && input.Jawaban == "" {
		return dto.ProgressResponse{}, errors.New("jawaban diperlukan untuk study case")
	}

	progress := entity.Progress{
		Id:       uuid.New(),
		MateriId: input.MateriId,
		KelasId:  input.KelasId,
		UserId:   input.UserId,
		Jawaban:  input.Jawaban,
	}

	err = u.progressRepo.CreateProgress(progress)
	if err != nil {
		return dto.ProgressResponse{}, err
	}

	return dto.ProgressResponse{
		Id:          progress.Id,
		MateriId:    progress.MateriId,
		UserId:      progress.UserId,
		Jawaban:     progress.Jawaban,
		IsCompleted: true,
	}, nil
}
