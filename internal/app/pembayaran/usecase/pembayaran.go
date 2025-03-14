package usecase

import (
	"errors"
	"fmt"
	"innovaspace/internal/app/pembayaran/repository"
	UserRepo "innovaspace/internal/app/user/repository"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/domain/entity"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PembayaranUsecaseItf interface {
	CreatePembayaran(userId uuid.UUID, tipeBayar string, durasi int) (*dto.PaymentResponse, error)
	GetPembayaranByUserId(userId uuid.UUID) ([]dto.PaymentResponse, error)
	UpdateStatusBayar(orderId string, status string) error
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

func calculate(durasi int) int {
	switch durasi {
	case 3:
		return 150000
	case 6:
		return 300000
	case 12:
		return 500000
	default:
		return 0
	}
}

func (u *PembayaranUsecase) CreatePembayaran(userId uuid.UUID, tipeBayar string, durasi int) (*dto.PaymentResponse, error) {
	if durasi != 3 && durasi != 6 && durasi != 12 {
		return nil, errors.New("invalid duration")
	}

	total := calculate(durasi)
	if total == 0 {
		return nil, errors.New("failed to calculate total")
	}

	user, err := u.userRepo.FindById(userId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	orderId := uuid.New().String()

	pembayaran := &entity.Pembayaran{
		Id:           uuid.New(),
		UserId:       userId,
		OrderId:      orderId,
		Total:        total,
		Status:       "pending",
		TipeBayar:    tipeBayar,
		CreatedDate:  time.Now(),
		ModifiedDate: time.Now(),
	}

	if err := u.pembayaranRepo.CreatePembayaran(pembayaran); err != nil {
		return nil, errors.New("failed to save transaction to database")
	}

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
	if snapResponse == nil {
		fmt.Println("Error snapResponse nil")
		return nil, errors.New("snap response is nil")
	}

	premiumStart := time.Now()
	premiumEnd := time.Now().AddDate(0, durasi, 0)

	user.IsPremium = true
	user.PremiumStart = &premiumStart
	user.PremiumEnd = &premiumEnd
	if err := u.userRepo.Update(user); err != nil {
		return nil, errors.New("failed to update user premium status")
	}

	fmt.Println("Mengupdate user:", user)
	if err := u.userRepo.Update(user); err != nil {
		log.Fatalf("Gagal mengupdate user: %v", err)
	}

	responsePembayaran := &dto.PaymentResponse{
		Id:          pembayaran.Id,
		OrderID:     pembayaran.OrderId,
		Total:       pembayaran.Total,
		Status:      pembayaran.Status,
		Token:       pembayaran.Token,
		PaymentUrl:  pembayaran.PaymentUrl,
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

func (u PembayaranUsecase) UpdateStatusBayar(orderId, status string) error {
	pembayaran, err := u.pembayaranRepo.GetPembayaranByOrderId(orderId)
	if err != nil {
		return errors.New("transaction not found")
	}

	tx := u.pembayaranRepo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic detected:", r)
			tx.Rollback()
		}
	}()

	err = u.pembayaranRepo.UpdatePembayaran(pembayaran.Id, status)
	if err != nil {
		tx.Rollback()
		return errors.New("failed to update payment status")
	}

	if status == "settlement" || status == "capture" {
		user, err := u.userRepo.FindById(pembayaran.UserId)
		if err != nil {
			tx.Rollback()
			return errors.New("user not found")
		}

		user.IsPremium = true
		now := time.Now()
		premiumEnd := now.AddDate(0, pembayaran.Durasi, 0) // Gunakan variabel yang sama
		user.PremiumStart = &now
		user.PremiumEnd = &premiumEnd

		if err := u.userRepo.UpdateWithTx(tx, user); err != nil {
			tx.Rollback()
			return errors.New("failed to update user premium status")
		}
	}

	tx.Commit()
	return nil
}
