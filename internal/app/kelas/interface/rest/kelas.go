package rest

import (
	"innovaspace/internal/app/kelas/usecase"

	"github.com/gofiber/fiber/v2"
)

type KelasHandler struct {
	kelasUsecase usecase.KelasUsecaseItf
}

func NewKelasHandler(routerGroup fiber.Router, kelasUsecase usecase.KelasUsecaseItf) {
	KelasHandler := KelasHandler{
		kelasUsecase: kelasUsecase,
	}

	routerGroup = routerGroup.Group("/kelas")
	routerGroup.Get("/", KelasHandler.GetAllKelas)
	routerGroup.Get("/get-detail-kelas/:id", KelasHandler.GetKelasDetails)
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

func (h *KelasHandler) GetAllKelas(ctx *fiber.Ctx) error {
	kelasList, err := h.kelasUsecase.GetAllKelas()
	if err != nil {
		return errorResponse(ctx, fiber.StatusInternalServerError,
			"Kelas Tidak Ditemukan", err.Error())
	}

	return successResponse(ctx, fiber.StatusOK, "Berhasil mendapatkan data", kelasList)
}

func (h *KelasHandler) GetKelasDetails(ctx *fiber.Ctx) error {
	kelasId := ctx.Params("id")
	if kelasId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Kelas ID is required",
		})
	}

	kelasDetail, err := h.kelasUsecase.GetKelasDetails(kelasId)
	if err != nil {
		return errorResponse(ctx, fiber.StatusInternalServerError,
			"Gagal mendapatkan data", err.Error())
	}

	return successResponse(ctx, fiber.StatusOK, "Berhasil mendapatkan data", kelasDetail)
}
