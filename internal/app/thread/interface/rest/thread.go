package rest

import (
	"innovaspace/internal/app/thread/usecase"
	"innovaspace/internal/domain/dto"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ThreadHandler struct {
	threadUsecase usecase.ThreadUsecaseItf
}

func NewThreadHandler(routerGroup fiber.Router, threadUsecase usecase.ThreadUsecaseItf) {
	ThreadHandler := ThreadHandler{
		threadUsecase: threadUsecase,
	}

	routerGroup = routerGroup.Group("/threads")
	routerGroup.Post("/create-thread/", ThreadHandler.CreateThread)
	routerGroup.Get("/", ThreadHandler.GetAllThreads)
	routerGroup.Patch("/:id", ThreadHandler.UpdateThread)
	routerGroup.Delete("/:id", ThreadHandler.DeleteThread)
}

func (h *ThreadHandler) CreateThread(ctx *fiber.Ctx) error {
	var input dto.CreateThreadRequest
	if err := ctx.BodyParser(&input); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})

	}

	// userId := ctx.Locals("userId").(uuid.UUID)

	thread, err := h.threadUsecase.CreateThread(input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create thread",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(thread)
}

func (h ThreadHandler) GetAllThreads(ctx *fiber.Ctx) error {
	threads, err := h.threadUsecase.GetAllThreads()
	if err != nil {
		log.Println("Error fetching threads:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch threads",
		})
	}
	return ctx.JSON(threads)
}

func (h ThreadHandler) UpdateThread(ctx *fiber.Ctx) error {
	threadId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid thread ID"})
	}

	userId, err := uuid.Parse(ctx.Locals("userId").(string))
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var input dto.UpdateThreadRequest
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	err = h.threadUsecase.UpdateThread(threadId, userId, input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update thread",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Thread updated successfully",
	})
}

func (h ThreadHandler) DeleteThread(ctx *fiber.Ctx) error {
	threadId, err := uuid.Parse(ctx.Params("thread_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	userId, err := uuid.Parse(ctx.Locals("userId").(string))
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.threadUsecase.DeleteThread(threadId, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete thread",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "Thread deleted successfully",
	})
}
