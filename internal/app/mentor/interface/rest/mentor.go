package rest

import (
	"innovaspace/internal/app/mentor/usecase"
	"innovaspace/internal/app/user/repository"

	"github.com/gofiber/fiber/v2"
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
	routerGroup.Post("/mentor-details", MentorHandler.GetMentorDetails)
	routerGroup.Post("/by-preferensi", MentorHandler.GetMentorsByUserPreferensi)
	routerGroup.Get("/", MentorHandler.GetAllMentors)
}

func (h MentorHandler) GetMentorDetails(ctx *fiber.Ctx) error {
	var request struct {
		Username string `json:"username"`
	}

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	mentor, err := h.mentorUsecase.GetMentorDetails(request.Username)
	if err != nil {
		if err.Error() == "mentor not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Mentor not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(mentor)
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
			"error": "Failed to fetch mentors",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"mentors": mentors,
	})
}
