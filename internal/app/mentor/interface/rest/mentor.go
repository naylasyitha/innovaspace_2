package rest

import (
	"innovaspace/internal/app/mentor/usecase"
	"innovaspace/internal/app/user/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MentorHandler struct {
	mentorUsecase usecase.MentorUsecaseItf
	userRepo      repository.UserMySQLItf
}

func NewMentorHandler(routerGroup fiber.Router, mentorUsecase usecase.MentorUsecaseItf, userRepo repository.UserMySQLItf) {
	MentorHandler := MentorHandler{
		mentorUsecase: mentorUsecase,
		userRepo:      userRepo,
	}

	routerGroup = routerGroup.Group("/mentors")
	routerGroup.Get("/mentor-details/:id", MentorHandler.GetMentorDetails)
	routerGroup.Post("/by-preferensi", MentorHandler.GetMentorsByUserPreferensi)
	routerGroup.Get("/", MentorHandler.GetAllMentors)
}

func (h MentorHandler) GetMentorDetails(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Parameter tidak valid",
			"errors":  "Mentor ID wajib diisi",
		})
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Parameter tidak valid",
			"errors":  "Format ID tidak valid",
		})
	}

	mentor, err := h.mentorUsecase.GetMentorDetails(parsedId)
	if err != nil {
		if err.Error() == "mentor not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Mentor tidak ditemukan",
				"errors":  err.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Mentor tidak ditemukan",
			"errors":  err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "Mentor berhasil ditemukan",
		"data": fiber.Map{
			"mentor": mentor,
		},
	})
}

func (h MentorHandler) GetMentorsByUserPreferensi(ctx *fiber.Ctx) error {
	var request struct {
		Preferensi string `json:"preferensi"`
	}

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	mentors, err := h.mentorUsecase.GetMentorsByPreferensi(request.Preferensi)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(mentors)
}

func (h MentorHandler) GetAllMentors(ctx *fiber.Ctx) error {
	mentors, err := h.mentorUsecase.GetAllMentors()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Mentor tidak ditemukan",
			"errors":  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Mentor berhasil ditemukan",
		"data": fiber.Map{
			"mentor": mentors,
		},
	})
}
