package bootstrap

import (
	"fmt"
	CommentHandler "innovaspace/internal/app/comment/interface/rest"
	CommentRepository "innovaspace/internal/app/comment/repository"
	CommentUsecase "innovaspace/internal/app/comment/usecase"
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
	"innovaspace/internal/infra/jwt"
	Seed "innovaspace/internal/infra/mysql"
	"innovaspace/internal/infra/storage"
	"innovaspace/internal/middleware"
	CORS "innovaspace/internal/middleware"
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

	err = database.AutoMigrate(&entity.Mentor{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	err = database.AutoMigrate(&entity.User{})
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

	jwt := jwt.NewJWT()
	middleware := middleware.NewMiddleware(jwt)
	CORS.CorsMiddleware(app)

	v1 := app.Group("/api/v1")

	inputValidation := validation.NewInputValidation()

	UserRepository := UserRepository.NewUserMySQL(database)
	UserUsecase := UserUsecase.NewUserUsecase(UserRepository, *inputValidation, *jwt)
	UserHandler.NewUserHandler(v1, UserUsecase, *inputValidation, middleware)

	MentorRepository := MentorRepository.NewMentorMySQL(database)
	MentorUsecase := MentorUsecase.NewMentorUsecase(MentorRepository, UserRepository)
	MentorHandler.NewMentorHandler(v1, MentorUsecase, UserRepository, middleware)

	CommentRepository := CommentRepository.NewCommentMySQL(database)
	CommentUsecase := CommentUsecase.NewCommentUsecase(CommentRepository)
	CommentHandler.NewCommentHandler(v1, CommentUsecase, middleware)

	ThreadRepository := ThreadRepository.NewThreadMySQL(database)
	ThreadUsecase := ThreadUsecase.NewThreadUsecase(ThreadRepository, CommentRepository)
	ThreadHandler.NewThreadHandler(v1, ThreadUsecase, middleware)

	Seed.SeedMentors(database)

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))
}
