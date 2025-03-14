package bootstrap

import (
	"fmt"
	CommentHandler "innovaspace/internal/app/comment/interface/rest"
	CommentRepository "innovaspace/internal/app/comment/repository"
	CommentUsecase "innovaspace/internal/app/comment/usecase"
	EnrollHandler "innovaspace/internal/app/enroll/interface/rest"
	EnrollRepository "innovaspace/internal/app/enroll/repository"
	EnrollUsecase "innovaspace/internal/app/enroll/usecase"
	KelasHandler "innovaspace/internal/app/kelas/interface/rest"
	KelasRepository "innovaspace/internal/app/kelas/repository"
	KelasUsecase "innovaspace/internal/app/kelas/usecase"
	MateriRepository "innovaspace/internal/app/materi/repository"
	MentorHandler "innovaspace/internal/app/mentor/interface/rest"
	MentorRepository "innovaspace/internal/app/mentor/repository"
	MentorUsecase "innovaspace/internal/app/mentor/usecase"
	PembayaranHandler "innovaspace/internal/app/pembayaran/interface/rest"
	PembayaranRepository "innovaspace/internal/app/pembayaran/repository"
	PembayaranUsecase "innovaspace/internal/app/pembayaran/usecase"
	ProgressHandler "innovaspace/internal/app/progress/interface/rest"
	ProgressRepository "innovaspace/internal/app/progress/repository"
	ProgressUsecase "innovaspace/internal/app/progress/usecase"
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
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SupabaseStorage storage.StorageSupabase
var SnapClient snap.Client

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

	err = database.AutoMigrate(&entity.Kelas{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	err = database.AutoMigrate(&entity.Pembayaran{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	err = database.AutoMigrate(&entity.Materi{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	err = database.AutoMigrate(&entity.Enroll{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	err = database.AutoMigrate(&entity.Progress{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	app := fiber.New()

	jwt := jwt.NewJWT()
	middlewareService := middleware.NewMiddleware(jwt)
	CORS.CorsMiddleware(app)
	snapClient := InitSnapClient()

	v1 := app.Group("/api/v1")

	inputValidation := validation.NewInputValidation()

	userRepo := UserRepository.NewUserMySQL(database)
	mentorRepo := MentorRepository.NewMentorMySQL(database)
	commentRepo := CommentRepository.NewCommentMySQL(database)
	threadRepo := ThreadRepository.NewThreadMySQL(database)
	materiRepo := MateriRepository.NewMateriMySQL(database)
	kelasRepo := KelasRepository.NewKelasMySQL(database)
	pembayaranRepo := PembayaranRepository.NewPembayaranMySQL(database)
	enrollRepo := EnrollRepository.NewEnrollMySQL(database)
	progressRepo := ProgressRepository.NewProgressMySQL(database)

	mentorUsecase := MentorUsecase.NewMentorUsecase(mentorRepo, userRepo)
	commentUsecase := CommentUsecase.NewCommentUsecase(commentRepo)
	threadUsecase := ThreadUsecase.NewThreadUsecase(threadRepo, commentRepo, userRepo)
	kelasUsecase := KelasUsecase.NewKelasUsecase(kelasRepo, materiRepo)
	pembayaranUsecase := PembayaranUsecase.NewPembayaranUsecase(pembayaranRepo, userRepo, snapClient)
	enrollUsecase := EnrollUsecase.NewEnrollUsecase(enrollRepo, userRepo, kelasRepo)
	progressUsecase := ProgressUsecase.NewProgressUsecase(progressRepo, materiRepo)
	userUsecase := UserUsecase.NewUserUsecase(userRepo, mentorRepo, enrollRepo, progressRepo, kelasRepo, *inputValidation, *jwt)

	MentorHandler.NewMentorHandler(v1, mentorUsecase, userRepo)
	CommentHandler.NewCommentHandler(v1, commentUsecase, middlewareService)
	ThreadHandler.NewThreadHandler(v1, threadUsecase, middlewareService)
	KelasHandler.NewKelasHandler(v1, kelasUsecase)
	PembayaranHandler.NewPembayaranHandler(v1, pembayaranUsecase, middlewareService)
	EnrollHandler.NewEnrollHandler(v1, enrollUsecase, middlewareService)
	ProgressHandler.NewProgressHandler(v1, progressUsecase, middlewareService)
	UserHandler.NewUserHandler(v1, userUsecase, *inputValidation, middlewareService)

	Seed.SeedMentors(database)
	Seed.SeedKelas(database)
	Seed.SeedMateri(database)

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))
}

func InitSnapClient() snap.Client {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans.Environment = midtrans.Sandbox

	snapClient := snap.Client{}
	snapClient.New(midtrans.ServerKey, midtrans.Sandbox)

	return snapClient
}
