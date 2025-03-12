package rest

import (
	"innovaspace/internal/app/enroll/usecase"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type EnrollHandler struct {
	EnrollUsecase usecase.EnrollUsecaseItf
	middleware    middleware.MiddlewareItf
}

func NewEnrollHandler(routerGroup fiber.Router, enrollUsecase usecase.EnrollUsecaseItf, middleware middleware.MiddlewareItf) {
	EnrollHandler := &EnrollHandler{
		EnrollUsecase: enrollUsecase,
		middleware:    middleware,
	}

	routerGroup = routerGroup.Group("/enroll")
	routerGroup.Post("/", EnrollHandler.middleware.Authentication, EnrollHandler.CreateEnroll)
}

func errorResponse(ctx *fiber.Ctx, status int, message, detail string) error {
	return ctx.Status(status).JSON(fiber.Map{
		"success": false,
		"message": message,
		"errors":  detail,
	})
}

func successResponse(ctx *fiber.Ctx, status int, message string, data interface{}) error {
	return ctx.Status(status).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func (h *EnrollHandler) CreateEnroll(ctx *fiber.Ctx) error {
	var input dto.EnrollRequest
	if err := ctx.BodyParser(&input); err != nil {
		return errorResponse(ctx, fiber.StatusBadRequest,
			"Permintaan tidak valid", "Format request tidak sesuai")
	}

	userId := ctx.Locals("userId").(uuid.UUID)
	input.UserId = userId

	enroll, err := h.EnrollUsecase.CreateEnroll(input)
	if err != nil {
		return errorResponse(ctx, fiber.StatusInternalServerError,
			"Gagal enroll kelas", err.Error())
	}

	return successResponse(ctx, fiber.StatusCreated, "Berhasil enroll kelas",
		fiber.Map{
			"enroll_id": enroll.Id,
			"kelas_id":  enroll.KelasId,
			"user_id":   enroll.UserId,
		})
}
