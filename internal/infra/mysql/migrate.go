package mysql

import (
	"innovaspace/internal/domain/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		entity.Mentor{},
	); err != nil {
		return err
	}

	if err := db.AutoMigrate(
		entity.User{},
	); err != nil {
		return err
	}

	if err := db.AutoMigrate(
		entity.Thread{},
	); err != nil {
		return err
	}

	if err := db.AutoMigrate(
		entity.Comment{},
	); err != nil {
		return err
	}

	if err := db.AutoMigrate(
		entity.Kelas{},
	); err != nil {
		return err
	}

	if err := db.AutoMigrate(
		entity.Pembayaran{},
	); err != nil {
		return err
	}

	if err := db.AutoMigrate(
		entity.Materi{},
	); err != nil {
		return err
	}

	if err := db.AutoMigrate(
		entity.Enroll{},
	); err != nil {
		return err
	}

	if err := db.AutoMigrate(
		entity.Progress{},
	); err != nil {
		return err
	}

	return nil
}
