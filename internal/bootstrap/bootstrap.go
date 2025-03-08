package bootstrap

import (
	"fmt"
	MentorHandler "innovaspace/internal/app/mentor/interface/rest"
	MentorRepository "innovaspace/internal/app/mentor/repository"
	MentorUsecase "innovaspace/internal/app/mentor/usecase"
	ThreadHandler "innovaspace/internal/app/thread/interface/rest"
	ThreadRepository "innovaspace/internal/app/thread/repository"
	ThreadUsecase "innovaspace/internal/app/thread/usecase"
	UserHandler "innovaspace/internal/app/user/interface/rest"
	UserRepository "innovaspace/internal/app/user/repository"
	UserUsecase "innovaspace/internal/app/user/usecase"
	"innovaspace/internal/domain/entity"
	"innovaspace/internal/infra/env"
	"innovaspace/internal/infra/fiber"
	Seed "innovaspace/internal/infra/mysql"
	"innovaspace/internal/infra/storage"
	"innovaspace/internal/middleware"
	"innovaspace/internal/validation"
	"log"
	"net/http"

	// "github.com/gofiber/fiber/v2/middleware/cors"
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

	err = database.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	err = database.AutoMigrate(&entity.Mentor{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	err = database.AutoMigrate(&entity.Thread{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	err = database.AutoMigrate(&entity.Comment{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	app := fiber.New()

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://be-intern.bccdev.id/nayla/")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

		if r.Method == "OPTIONS" {
			w.Write([]byte("allowed"))
			return
		}
	})

	v1 := app.Group("/api/v1")

	inputValidation := validation.NewInputValidation()

	UserRepository := UserRepository.NewUserMySQL(database)
	UserUsecase := UserUsecase.NewUserUsecase(UserRepository, *inputValidation, SupabaseStorage)
	UserHandler.NewUserHandler(v1, UserUsecase, *inputValidation, &middleware.Middleware{})

	MentorRepository := MentorRepository.NewMentorMySQL(database)
	MentorUsecase := MentorUsecase.NewMentorUsecase(MentorRepository, UserRepository)
	MentorHandler.NewMentorHandler(v1, MentorUsecase, UserRepository)

	ThreadRepository := ThreadRepository.NewThreadMySQL(database)
	ThreadUsecase := ThreadUsecase.NewThreadUsecase(ThreadRepository)
	ThreadHandler.NewThreadHandler(v1, ThreadUsecase)

	Seed.SeedMentors(database)

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))
}
