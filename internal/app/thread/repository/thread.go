package repository

import (
	"innovaspace/internal/domain/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ThreadMySQLItf interface {
	CreateThread(thread entity.Thread) error
	GetAllThreads() ([]entity.Thread, error)
	GetThreadById(threadId uuid.UUID) (entity.Thread, error)
	UpdateThread(thread *entity.Thread) error
	DeleteThread(threadId uuid.UUID) error
}

type ThreadMySQL struct {
	db *gorm.DB
}

func NewThreadMySQL(db *gorm.DB) ThreadMySQLItf {
	return &ThreadMySQL{db}
}

func (r ThreadMySQL) CreateThread(thread entity.Thread) error {
	return r.db.Create(thread).Error
}

func (r ThreadMySQL) GetAllThreads() ([]entity.Thread, error) {
	var threads []entity.Thread
	err := r.db.Find(&threads).Error
	return threads, err
}

func (r ThreadMySQL) GetThreadById(threadId uuid.UUID) (entity.Thread, error) {
	var thread entity.Thread
	err := r.db.First(&thread, "thread_id = ?", threadId).Error
	return thread, err
}

func (r *ThreadMySQL) UpdateThread(thread *entity.Thread) error {
	updates := map[string]interface{}{}
	if thread.Kategori != "" {
		updates["kategori"] = thread.Kategori
	}
	if thread.Isi != "" {
		updates["isi"] = thread.Isi
	}

	return r.db.Model(thread).Where("id = ?", thread.Id).Updates(updates).Error
}

func (r ThreadMySQL) DeleteThread(threadId uuid.UUID) error {
	return r.db.Delete(entity.Thread{}, "thread_id = ?", threadId).Error
}
