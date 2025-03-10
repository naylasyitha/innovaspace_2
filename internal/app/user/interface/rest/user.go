package rest

import (
	"innovaspace/internal/app/user/usecase"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/middleware"
	"innovaspace/internal/validation"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	usecase    usecase.UserUsecaseItf
	validator  validation.InputValidation
	middleware middleware.MiddlewareItf
}

func NewUserHandler(routerGroup fiber.Router, userUsecase usecase.UserUsecaseItf, Validator validation.InputValidation, middleware middleware.MiddlewareItf) {
	userHandler := UserHandler{
		usecase:    userUsecase,
		validator:  Validator,
		middleware: middleware,
	}

	routerGroup = routerGroup.Group("/users")
	routerGroup.Post("/register", userHandler.Register)
	routerGroup.Post("/login", userHandler.Login)
	routerGroup.Get("/get-profile/:id", userHandler.middleware.Authentication, userHandler.GetProfile)
	routerGroup.Patch("/update/:id", userHandler.middleware.Authentication, userHandler.Update)
	routerGroup.Patch("/set-mentor/:id", userHandler.middleware.Authentication, userHandler.SetMentor)
	routerGroup.Patch("/update-mentor/:id", userHandler.middleware.Authentication, userHandler.UpdateMentor)
}

func errorResponse(ctx *fiber.Ctx, status int, message string, errors interface{}) error {
	return ctx.Status(status).JSON(fiber.Map{
		"success": false,
		"message": message,
		"errors":  errors,
	})
}

func successResponse(ctx *fiber.Ctx, data interface{}) error {
	return ctx.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	var register dto.Register

	if err := ctx.BodyParser(&register); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Permintaan tidak valid",
			"errors":  "Format request tidak sesuai",
		})
	}

	if err := h.validator.Validate(register); err != nil {
		log.Printf("Validation error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"success": false,
			"message": "Validasi gagal",
			"errors":  err.Error(),
		})
	}

	createdUser, err := h.usecase.Register(register)
	if err != nil {
		errorMap := fiber.Map{"general": err.Error()}
		status := fiber.StatusBadRequest

		if strings.Contains(err.Error(), "duplicate") {
			status = fiber.StatusConflict
			errorMap = fiber.Map{
				"email": "Email sudah terdaftar",
			}
		}
		if strings.Contains(err.Error(), "username already exists") {
			status = fiber.StatusConflict
			errorMap = fiber.Map{
				"username": "Username sudah digunakan",
			}
		}

		return ctx.Status(status).JSON(fiber.Map{
			"success": false,
			"message": "Registrasi gagal",
			"errors":  errorMap,
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Registrasi berhasil",
		"data": fiber.Map{
			"user": fiber.Map{
				"id":           createdUser.Id,
				"username":     createdUser.Username,
				"email":        createdUser.Email,
				"nama":         createdUser.Nama,
				"institusi":    createdUser.Institusi,
				"profile_pict": createdUser.UserPict,
			},
		},
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	var login dto.Login

	if err := ctx.BodyParser(&login); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Permintaan tidak valid",
			"errors":  "Format request tidak sesuai",
		})
	}

	token, err := h.usecase.Login(login)
	if err != nil {
		status := fiber.StatusInternalServerError
		message := "Terjadi kesalahan sistem"
		errors := "Silakan coba beberapa saat lagi"

		return ctx.Status(status).JSON(fiber.Map{
			"success": false,
			"message": message,
			"errors":  errors,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Login berhasil",
		"data": fiber.Map{
			"token": token,
		},
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Parameter tidak valid",
			"errors":  "ID wajib diisi",
		})
	}

	ParseId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Parameter tidak valid",
			"errors":  "Format ID tidak valid",
		})
	}

	user, err := h.usecase.GetProfileById(ParseId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Profil tidak ditemukan",
			"errors":  err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "Profil berhasil ditemukan",
		"data": fiber.Map{
			"profile": user,
		},
	})
}

func (h *UserHandler) Update(ctx *fiber.Ctx) error {
	userId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Parameter tidak valid",
			"errors":  "Format ID tidak valid",
		})
	}

	var request dto.UpdateProfile
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Permintaan tidak valid",
			"errors":  "Format request tidak sesuai",
		})
	}

	err = h.usecase.UpdateUser(userId, request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Gagal memperbarui data",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Data pengguna berhasil diperbarui",
		"data":    fiber.Map{},
	})
}

func (h *UserHandler) SetMentor(ctx *fiber.Ctx) error {
	userId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Parameter tidak valid",
			"errors":  "Format ID tidak valid",
		})
	}

	var request dto.SetMentor
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Permintaan tidak valid",
			"errors":  "Format request tidak sesuai",
		})
	}

	if request.MentorId == uuid.Nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Parameter tidak valid",
			"errors":  "Mentor ID wajib diisi",
		})
	}

	err = h.usecase.SetMentor(userId, request)
	if err != nil {
		status := fiber.StatusInternalServerError
		message := "Gagal menetapkan mentor"
		errors := fiber.Map{"general": err.Error()}

		if err.Error() == "user has mentor" {
			status = fiber.StatusConflict
			message = "Pengguna sudah memiliki mentor"
			errors = fiber.Map{"mentor": "Tidak bisa menetapkan mentor baru"}
		}

		return ctx.Status(status).JSON(fiber.Map{
			"success": false,
			"message": message,
			"errors":  errors,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Mentor berhasil ditambahkan",
		"data":    fiber.Map{},
	})
}

func (h *UserHandler) UpdateMentor(ctx *fiber.Ctx) error {
	userId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Parameter tidak valid",
			"errors":  "Format ID tidak valid",
		})
	}

	var request dto.SetMentor
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Permintaan tidak valid",
			"errors":  "Format request tidak sesuai",
		})
	}

	if request.MentorId == uuid.Nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Parameter tidak valid",
			"errors":  "Mentor ID wajib diisi",
		})
	}

	err = h.usecase.UpdateMentor(userId, request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Gagal memperbarui mentor",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Mentor berhasil diperbarui",
		"data":    fiber.Map{},
	})
}
