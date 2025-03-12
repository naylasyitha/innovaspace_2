package rest

import (
	"innovaspace/internal/app/pembayaran/usecase"
	"innovaspace/internal/domain/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PembayaranHandler struct {
	PembayaranUsecase usecase.PembayaranUsecaseItf
}

func NewPembayaranHandler(routerGroup fiber.Router, pembayaranUsecase usecase.PembayaranUsecaseItf) {
	PembayaranHandler := PembayaranHandler{
		PembayaranUsecase: pembayaranUsecase,
	}

	routerGroup = routerGroup.Group("/pembayaran")
	routerGroup.Post("/create-pembayaran", PembayaranHandler.CreatePembayaran)
	routerGroup.Get("/:id", PembayaranHandler.GetPembayaranById)
	// routerGroup.Get("/pembayaran-user", PembayaranHandler.GetPembayaranByUserID)
}

func (h PembayaranHandler) CreatePembayaran(ctx *fiber.Ctx) error {
	var request dto.PaymentRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userId, err := uuid.Parse(ctx.Locals("userId").(string))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user Id",
		})
	}

	pembayaran, err := h.PembayaranUsecase.CreatePembayaran(userId, request.Total, request.TipeBayar, request.Durasi)
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
