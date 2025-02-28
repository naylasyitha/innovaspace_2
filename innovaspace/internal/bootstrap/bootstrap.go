package bootstrap

import (
	"fmt"
	UserHandler "innovaspace/internal/app/user/interface/rest"
	UserRepository "innovaspace/internal/app/user/repository"
	UserUsecase "innovaspace/internal/app/user/usecase"
	"innovaspace/internal/domain/entity"
	"innovaspace/internal/infra/env"
	"innovaspace/internal/infra/fiber"
	"innovaspace/internal/validation"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

	err = database.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	app := fiber.New()

	v1 := app.Group("/api/v1")

	inputValidation := validation.NewInputValidation()

	UserRepository := UserRepository.NewUserMySQL(database)
	UserUsecase := UserUsecase.NewUserUsecase(UserRepository, *inputValidation)
	UserHandler.NewUserHandler(v1, UserUsecase, *inputValidation)

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))
}
