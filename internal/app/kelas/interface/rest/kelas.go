package rest

import (
	"innovaspace/internal/app/kelas/usecase"
	"innovaspace/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type KelasHandler struct {
	kelasUsecase usecase.KelasUsecaseItf
	middleware   middleware.MiddlewareItf
}

func NewKelasHandler(routerGroup fiber.Router, kelasUsecase usecase.KelasUsecaseItf, middleware middleware.MiddlewareItf) {
	KelasHandler := KelasHandler{
		kelasUsecase: kelasUsecase,
		middleware:   middleware,
	}

	routerGroup = routerGroup.Group("/kelas")
	routerGroup.Get("/", KelasHandler.middleware.Authentication, KelasHandler.GetAllKelas)
	routerGroup.Get("/get-detail-kelas/:id", KelasHandler.middleware.Authentication, KelasHandler.GetKelasDetails)
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
	kelasId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return errorResponse(ctx, fiber.StatusBadRequest,
			"ID kelas tidak valid", "Format ID tidak valid")
	}

	kelasDetail, err := h.kelasUsecase.GetKelasDetails(kelasId)
	if err != nil {
		return errorResponse(ctx, fiber.StatusInternalServerError,
			"Gagal mendapatkan data", err.Error())
	}

	return successResponse(ctx, fiber.StatusOK, "Berhasil mendapatkan data", kelasDetail)
}
