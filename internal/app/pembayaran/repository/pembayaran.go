package repository

import (
	"innovaspace/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PembayaranMySQLItf interface {
	CreatePembayaran(*entity.Pembayaran) error
	GetPembayaranById(Id uuid.UUID) (*entity.Pembayaran, error)
	GetPembayaranByUserId(userId uuid.UUID) ([]entity.Pembayaran, error)
	GetPembayaranByOrderId(orderId string) (entity.Pembayaran, error)
	UpdatePembayaran(id uuid.UUID, status string) error
}

type PembayaranMySQL struct {
	db *gorm.DB
}

func NewPembayaranMySQL(db *gorm.DB) PembayaranMySQLItf {
	return &PembayaranMySQL{db}
}

func (r *PembayaranMySQL) CreatePembayaran(pembayaran *entity.Pembayaran) error {
	return r.db.Create(pembayaran).Error
}

func (r *PembayaranMySQL) GetPembayaranById(Id uuid.UUID) (*entity.Pembayaran, error) {
	var pembayaran entity.Pembayaran
	if err := r.db.First(&pembayaran, "id = ?", Id).Error; err != nil {
		return nil, err
	}
	return &pembayaran, nil
}

func (r *PembayaranMySQL) GetPembayaranByUserId(userId uuid.UUID) ([]entity.Pembayaran, error) {
	var pembayaran []entity.Pembayaran
	if err := r.db.Where("user_id = ?", userId).Find(&pembayaran).Error; err != nil {
		return nil, err
	}
	return pembayaran, nil
}

func (r *PembayaranMySQL) GetPembayaranByOrderId(orderId string) (entity.Pembayaran, error) {
	var pembayaran entity.Pembayaran
	if err := r.db.Where("order_id = ?", orderId).Find(&pembayaran).Error; err != nil {
		return entity.Pembayaran{}, err
	}
	return pembayaran, nil
}

func (r *PembayaranMySQL) UpdatePembayaran(id uuid.UUID, status string) error {
	return r.db.Model(&entity.Pembayaran{}).Where("id = ?", id).Update("status", status).Error
}
