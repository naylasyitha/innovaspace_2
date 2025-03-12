package rest

import (
	"innovaspace/internal/app/progress/usecase"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProgressHandler struct {
	ProgressUsecase usecase.ProgressUsecaseItf
	middleware      middleware.MiddlewareItf
}

func NewProgressHandler(routerGroup fiber.Router, progressUsecase usecase.ProgressUsecaseItf, middleware middleware.MiddlewareItf) {
	ProgressHandler := &ProgressHandler{
		ProgressUsecase: progressUsecase,
		middleware:      middleware,
	}

	routerGroup = routerGroup.Group("/progress")
	routerGroup.Post("/", ProgressHandler.middleware.Authentication, ProgressHandler.CreateProgress)
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

func (h *ProgressHandler) CreateProgress(ctx *fiber.Ctx) error {
	var input dto.ProgressRequest
	if err := ctx.BodyParser(&input); err != nil {
		return errorResponse(ctx, fiber.StatusBadRequest,
			"Permintaan tidak valid", "Format request tidak sesuai")
	}

	userId := ctx.Locals("userId").(uuid.UUID)
	input.UserId = userId

	progress, err := h.ProgressUsecase.CreateProgress(input)
	if err != nil {
		return errorResponse(ctx, fiber.StatusInternalServerError,
			"Gagal set progress belajar", err.Error())
	}

	return successResponse(ctx, fiber.StatusCreated, "Berhasil set progress belajar",
		fiber.Map{
			"progress_id":  progress.Id,
			"materi_id":    progress.MateriId,
			"user_id":      progress.UserId,
			"is_completed": progress.IsCompleted,
		})
}
