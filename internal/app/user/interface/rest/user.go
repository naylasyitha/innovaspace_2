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
				"email":    "Email sudah terdaftar",
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
			"error": "invalid request body",
		})
	}

	token, err := h.usecase.Login(login)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})

	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success login user",
		"token":   token,
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "missing ID",
		})
	}

	ParseId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	user, err := h.usecase.GetProfileById(ParseId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(user)
}

func (h *UserHandler) Update(ctx *fiber.Ctx) error {
	userId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	var request dto.UpdateProfile
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err = h.usecase.UpdateUser(userId, request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}

func (h *UserHandler) SetMentor(ctx *fiber.Ctx) error {
	userId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var request dto.SetMentor
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if request.MentorId == uuid.Nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Mentor ID required",
		})
	}

	err = h.usecase.SetMentor(userId, request)
	if err != nil {
		if err.Error() == "user has mentor" {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User already has mentor"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "mentor set succesfully",
	})
}

func (h *UserHandler) UpdateMentor(ctx *fiber.Ctx) error {
	userId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var request dto.SetMentor
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err = h.usecase.UpdateMentor(userId, request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Mentor updated succesfulyy",
	})
}
