package rest

import (
	"innovaspace/internal/app/mentor/usecase"
	"innovaspace/internal/app/user/repository"
	"innovaspace/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MentorHandler struct {
	mentorUsecase usecase.MentorUsecaseItf
	userRepo      repository.UserMySQLItf
	middleware    middleware.MiddlewareItf
}

func NewMentorHandler(routerGroup fiber.Router, mentorUsecase usecase.MentorUsecaseItf, userRepo repository.UserMySQLItf, middleware middleware.MiddlewareItf) {
	MentorHandler := MentorHandler{
		mentorUsecase: mentorUsecase,
		userRepo:      userRepo,
		middleware:    middleware,
	}

	routerGroup = routerGroup.Group("/mentors")
	routerGroup.Get("/mentor-details/:id", MentorHandler.middleware.Authentication, MentorHandler.GetMentorDetails)
	routerGroup.Post("/by-preferensi", MentorHandler.GetMentorsByUserPreferensi)
	routerGroup.Get("/", MentorHandler.GetAllMentors)
}

func (h MentorHandler) GetMentorDetails(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "missing ID",
		})
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	mentor, err := h.mentorUsecase.GetMentorDetails(parsedId)
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
