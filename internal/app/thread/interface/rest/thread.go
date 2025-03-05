package rest

import (
	"innovaspace/internal/app/thread/usecase"
	"innovaspace/internal/domain/dto"

	// "github.com/gofiber/fiber"
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
	routerGroup.Post("/", ThreadHandler.CreateThread)
	routerGroup.Get("/", ThreadHandler.GetAllThreads)
	routerGroup.Get("/:id", ThreadHandler.GetThreadById)
	routerGroup.Patch("/:id", ThreadHandler.UpdateThread)
	routerGroup.Delete("/:id", ThreadHandler.DeleteThread)
}

func (h ThreadHandler) CreateThread(ctx *fiber.Ctx) error {
	var input dto.CreateThreadRequest
	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch threads",
		})
	}
	return ctx.JSON(threads)
}

func (h ThreadHandler) GetThreadById(ctx *fiber.Ctx) error {
	threadId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	thread, err := h.threadUsecase.GetThreadById(threadId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Thread not found",
		})
	}

	return ctx.JSON(thread)
}

func (h ThreadHandler) UpdateThread(ctx *fiber.Ctx) error {
	threadId, _ := uuid.Parse(ctx.Params("id"))
	var input dto.UpdateThreadRequest
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errror": "Invalid input",
		})
	}
	err := h.threadUsecase.UpdateThread(threadId, input)
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
	threadId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	err = h.threadUsecase.DeleteThread(threadId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete thread",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "Thread deleted successfully",
	})
}
