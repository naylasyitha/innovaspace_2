package usecase

import (
	"errors"
	"innovaspace/internal/app/pembayaran/repository"
	UserRepo "innovaspace/internal/app/user/repository"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"
	"time"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PembayaranUsecaseItf interface {
	CreatePembayaran(userId uuid.UUID, total int, tipeBayar string, durasi int) (*dto.PaymentResponse, error)
	GetPembayaranById(id uuid.UUID) (*dto.PaymentResponse, error)
	GetPembayaranByUserId(userId uuid.UUID) ([]dto.PaymentResponse, error)
	UpdateStatusBayar(id uuid.UUID, status string) error
}

type PembayaranUsecase struct {
	pembayaranRepo repository.PembayaranMySQLItf
	userRepo       UserRepo.UserMySQLItf
	SnapClient     snap.Client
}

func NewPembayaranUsecase(paymentRepo repository.PembayaranMySQLItf, userRepo UserRepo.UserMySQLItf, snapClient snap.Client) PembayaranUsecaseItf {
	return &PembayaranUsecase{
		pembayaranRepo: paymentRepo,
		userRepo:       userRepo,
		SnapClient:     snapClient,
	}
}

func (u *PembayaranUsecase) CreatePembayaran(userId uuid.UUID, total int, tipeBayar string, durasi int) (*dto.PaymentResponse, error) {
	user, err := u.userRepo.FindById(userId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	orderId := uuid.New().String()
	snapRequest := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: int64(total),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Email,
			Email: user.Email,
		},
	}

	snapResponse, err := u.SnapClient.CreateTransaction(snapRequest)
	if err != nil {
		return nil, errors.New("failed to create snap transaction")
	}

	pembayaran := &entity.Pembayaran{
		Id:           uuid.New(),
		UserId:       userId,
		OrderId:      orderId,
		Total:        total,
		Status:       "pending",
		TipeBayar:    tipeBayar,
		Token:        snapResponse.Token,
		CreatedDate:  time.Now(),
		ModifiedDate: time.Now(),
	}

	if err := u.pembayaranRepo.CreatePembayaran(pembayaran); err != nil {
		return nil, errors.New("failed to save transaction to database")
	}
	premiumStart := time.Now()
	premiumEnd := time.Now().AddDate(1, durasi, 0)

	user.IsPremium = true
	user.PremiumStart = &premiumStart
	user.PremiumEnd = &premiumEnd
	if err := u.userRepo.Update(user); err != nil {
		return nil, errors.New("failed to update user premium status")
	}

	responsePembayaran := &dto.PaymentResponse{
		Id:          pembayaran.Id,
		OrderID:     pembayaran.OrderId,
		Total:       pembayaran.Total,
		Status:      pembayaran.Status,
		Token:       pembayaran.Token,
		CreatedDate: pembayaran.CreatedDate.Format(time.RFC3339),
	}

	return responsePembayaran, nil
}

func (u PembayaranUsecase) GetPembayaranById(id uuid.UUID) (*dto.PaymentResponse, error) {
	pembayaran, err := u.pembayaranRepo.GetPembayaranById(id)
	if err != nil {
		return nil, errors.New("transaction not found")
	}

	responsePembayaran := &dto.PaymentResponse{
		Id:          pembayaran.Id,
		OrderID:     pembayaran.OrderId,
		Total:       pembayaran.Total,
		Status:      pembayaran.Status,
		Token:       pembayaran.Token,
		CreatedDate: pembayaran.CreatedDate.Format(time.RFC3339),
	}

	return responsePembayaran, nil
}

func (u PembayaranUsecase) GetPembayaranByUserId(userId uuid.UUID) ([]dto.PaymentResponse, error) {
	pembayaran, err := u.pembayaranRepo.GetPembayaranByUserId(userId)
	if err != nil {
		return nil, errors.New("failed to fetch transactions")
	}

	var responsePembayaran []dto.PaymentResponse
	for _, pembayaran := range pembayaran {
		responsePembayaran = append(responsePembayaran, dto.PaymentResponse{
			Id:          pembayaran.Id,
			OrderID:     pembayaran.OrderId,
			Total:       pembayaran.Total,
			Status:      pembayaran.Status,
			Token:       pembayaran.Token,
			CreatedDate: pembayaran.CreatedDate.Format(time.RFC3339),
		})
	}
	return responsePembayaran, nil
}

func (u PembayaranUsecase) UpdateStatusBayar(id uuid.UUID, status string) error {
	return u.pembayaranRepo.UpdatePembayaran(id, status)
}
