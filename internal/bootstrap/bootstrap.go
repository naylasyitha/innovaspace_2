package bootstrap

import (
	"fmt"
	MentorHandler "innovaspace/internal/app/mentor/interface/rest"
	MentorRepository "innovaspace/internal/app/mentor/repository"
	MentorUsecase "innovaspace/internal/app/mentor/usecase"
	UserHandler "innovaspace/internal/app/user/interface/rest"
	UserRepository "innovaspace/internal/app/user/repository"
	UserUsecase "innovaspace/internal/app/user/usecase"
	"innovaspace/internal/domain/entity"
	"innovaspace/internal/infra/env"
	"innovaspace/internal/infra/fiber"
	Seed "innovaspace/internal/infra/mysql"
	"innovaspace/internal/infra/storage"
	"innovaspace/internal/validation"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SupabaseStorage storage.StorageSupabase

func Start() error {
	config, err := env.New()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&entity.User{}, &entity.Mentor{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	app := fiber.New()

	v1 := app.Group("/api/v1")

	inputValidation := validation.NewInputValidation()

	UserRepository := UserRepository.NewUserMySQL(database)
	UserUsecase := UserUsecase.NewUserUsecase(UserRepository, *inputValidation, SupabaseStorage)
	UserHandler.NewUserHandler(v1, UserUsecase, *inputValidation)

	MentorRepository := MentorRepository.NewMentorMySQL(database)
	MentorUsecase := MentorUsecase.NewMentorUsecase(MentorRepository, UserRepository)
	MentorHandler.NewMentorHandler(v1, MentorUsecase, UserRepository)

	Seed.SeedMentors(database)

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))
}
