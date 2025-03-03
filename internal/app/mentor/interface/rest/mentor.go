package rest

import (
	"innovaspace/internal/app/mentor/usecase"
	"innovaspace/internal/app/user/repository"
	"log"

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
	routerGroup.Post("/preferensi", MentorHandler.GetMentorsByUserPreferensi)
	routerGroup.Post("/", MentorHandler.GetMentorsByPreferensi)
}

func (h MentorHandler) GetMentorsByUserPreferensi(ctx *fiber.Ctx) error {
	userId, err := uuid.Parse(ctx.Query("user_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	mentors, err := h.mentorUsecase.GetMentorsDetails(userId)
	if err != nil {
		log.Println("Error fetching mentors:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(mentors)
}

func (h MentorHandler) GetMentorsByPreferensi(ctx *fiber.Ctx) error {
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
