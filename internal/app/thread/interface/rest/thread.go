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
	routerGroup.Get("/show-all-thread/", ThreadHandler.GetAllThreads)
	routerGroup.Patch("/update-thread/:thread_id", ThreadHandler.UpdateThread)
	routerGroup.Delete("/delete-thread/:thread_id", ThreadHandler.DeleteThread)
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
	threadId, err := uuid.Parse(ctx.Params("thread_id"))
	// fmt.Println(threadId)
	// fmt.Println(err)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid thread ID"})
	}

	var input dto.UpdateThreadRequest
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// userId, err := input.UserId.Value()
	// if err != nil {
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	// }

	// err = h.threadUsecase.UpdateThread(threadId, userId, input)
	err = h.threadUsecase.UpdateThread(threadId, input)
	if err != nil {
		if err.Error() == "Unauthorized" {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "unauthorized"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
		// return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 	"error": "Failed to update thread",
		// })
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
