package rest

import (
	"innovaspace/internal/app/pembayaran/usecase"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PembayaranHandler struct {
	PembayaranUsecase usecase.PembayaranUsecaseItf
	middleware        middleware.MiddlewareItf
}

func NewPembayaranHandler(routerGroup fiber.Router, pembayaranUsecase usecase.PembayaranUsecaseItf, middleware middleware.MiddlewareItf) {
	PembayaranHandler := PembayaranHandler{
		PembayaranUsecase: pembayaranUsecase,
		middleware:        middleware,
	}

	routerGroup = routerGroup.Group("/pembayaran")
	routerGroup.Post("/create-pembayaran", PembayaranHandler.middleware.Authentication, PembayaranHandler.CreatePembayaran)
	routerGroup.Get("/:id", PembayaranHandler.GetPembayaranById)
	routerGroup.Post("/status-pembayaran", PembayaranHandler.MidtransHandler)
	// routerGroup.Get("/pembayaran-user", PembayaranHandler.GetPembayaranByUserID)
}

func (h PembayaranHandler) CreatePembayaran(ctx *fiber.Ctx) error {
	var request dto.PaymentRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userId, ok := ctx.Locals("userId").(uuid.UUID)
	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user ID",
		})
	}

	pembayaran, err := h.PembayaranUsecase.CreatePembayaran(userId, request.TipeBayar, request.Durasi)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(pembayaran)
}

func (h PembayaranHandler) GetPembayaranById(ctx *fiber.Ctx) error {
	userId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid transaction Id",
		})
	}

	pembayaran, err := h.PembayaranUsecase.GetPembayaranByUserId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(pembayaran)
}

func (h PembayaranHandler) MidtransHandler(ctx *fiber.Ctx) error {
	var notification dto.MidtransWebhookRequest
	if err := ctx.BodyParser(&notification); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	err := h.PembayaranUsecase.UpdateStatusBayar(notification.OrderId, notification.TransactionStatus)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update payment status",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Payment status updated successfully",
	})
}

// func (h PembayaranHandler) GetPembayaranByUserID(ctx *fiber.Ctx) error {
// 	userId, err := uuid.Parse(ctx.Params("user_id"))
// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "invalid user ID",
// 		})
// 	}

// 	pembayaran, err := h.PembayaranUsecase.GetPembayaranByUserId(userId)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(pembayaran)
// }
