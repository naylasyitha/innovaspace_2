package usecase

import (
	"errors"
	"innovaspace/internal/app/kelas/repository"
	MateriRepo "innovaspace/internal/app/materi/repository"
	"innovaspace/internal/domain/dto"

	"github.com/google/uuid"
)

type KelasUsecaseItf interface {
	GetAllKelas() ([]dto.Kelas, error)
	GetKelasDetails(kelasId uuid.UUID) (*dto.KelasDetailResponse, error)
}

type KelasUsecase struct {
	kelasRepo  repository.KelasMySQLItf
	materiRepo MateriRepo.MateriMySQLItf
}

func NewKelasUsecase(kelasRepo repository.KelasMySQLItf, materiRepo MateriRepo.MateriMySQLItf) KelasUsecaseItf {
	return &KelasUsecase{
		kelasRepo:  kelasRepo,
		materiRepo: materiRepo,
	}
}

func (u *KelasUsecase) GetAllKelas() ([]dto.Kelas, error) {
	kelasList, err := u.kelasRepo.GetAllKelas()
	if err != nil {
		return nil, errors.New("gagal mengambil data kelas: " + err.Error())
	}

	var response []dto.Kelas
	for _, kelas := range kelasList {
		response = append(response, dto.Kelas{
			KelasId:          kelas.Id,
			Nama:             kelas.Nama,
			Deskripsi:        kelas.Deskripsi,
			Kategori:         kelas.Kategori,
			JumlahMateri:     kelas.JumlahMateri,
			CoverCourse:      kelas.CoverCourse,
			TingkatKesulitan: kelas.TingkatKesulitan,
			Durasi:           kelas.Durasi,
		})
	}
	return response, nil
}

func (u *KelasUsecase) GetKelasDetails(kelasId uuid.UUID) (*dto.KelasDetailResponse, error) {
	kelas, err := u.kelasRepo.FindById(kelasId)
	if err != nil {
		return nil, err
	}

	materies, err := u.materiRepo.GetMateriByKelasId(kelasId)
	if err != nil {
		return nil, err
	}

	var materiResponses []dto.Materi
	for _, materi := range materies {
		materiResponses = append(materiResponses, dto.Materi{
			MateriId:  materi.Id,
			KelasId:   materi.KelasId,
			Judul:     materi.Judul,
			Deskripsi: materi.Deskripsi,
			IsFree:    materi.IsFree,
			PathFile:  materi.PathFile,
		})
	}

	response := &dto.KelasDetailResponse{
		KelasId:          kelas.Id,
		Nama:             kelas.Nama,
		Deskripsi:        kelas.Deskripsi,
		Kategori:         kelas.Kategori,
		JumlahMateri:     kelas.JumlahMateri,
		CoverCourse:      kelas.CoverCourse,
		TingkatKesulitan: kelas.TingkatKesulitan,
		Durasi:           kelas.Durasi,
		Materi:           materiResponses,
	}

	return response, nil
}
