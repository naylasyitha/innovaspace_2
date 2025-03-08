package rest

import (
	"innovaspace/internal/app/comment/usecase"
	"innovaspace/internal/domain/dto"

	// "innovaspace/internal/middleware"

	"github.com/gofiber/fiber/v2"

	"github.com/google/uuid"
)

type CommentHandler struct {
	CommentUsecase usecase.CommentUsecaseItf
	// validator      middleware.MiddlewareItf
}

func NewCommentHandler(routerGroup fiber.Router, commentUsecase usecase.CommentUsecaseItf) {
	CommentHandler := CommentHandler{
		CommentUsecase: commentUsecase,
		// validator:      authMiddleware,
	}

	routerGroup = routerGroup.Group("/comments")
	routerGroup.Post("/create-comment", CommentHandler.CreateComment)
	routerGroup.Get("/get-comments/:threadId", CommentHandler.GetCommentsByThread)
	routerGroup.Patch("/update-comment:threadId", CommentHandler.UpdateComment)
	routerGroup.Delete("/delete-comment:threadId", CommentHandler.DeleteComment)
}

func (h *CommentHandler) CreateComment(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uuid.UUID)

	var input dto.CreateCommentRequest
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	comment, err := h.CommentUsecase.CreateComment(userId, input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create comment"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Comment created successfully",
		"data":    comment,
	})
}

func (h *CommentHandler) GetCommentsByThread(ctx *fiber.Ctx) error {
	threadId, err := uuid.Parse(ctx.Params("threadId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	comments, err := h.CommentUsecase.GetCommentsByThread(threadId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(comments)
}

func (h *CommentHandler) UpdateComment(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uuid.UUID)
	commentId, err := uuid.Parse(ctx.Params("commentId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid ID",
		})
	}

	var input dto.UpdateCommentRequest
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	updatedComment, err := h.CommentUsecase.UpdateComment(userId, commentId, input)
	if err != nil {
		if err.Error() == "unauthorized" {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You can only update your own comment"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}

	return ctx.JSON(updatedComment)
}

func (h *CommentHandler) DeleteComment(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uuid.UUID)
	commentId, err := uuid.Parse(ctx.Params("commentId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if err := h.CommentUsecase.DeleteComment(userId, commentId); err != nil {
		if err.Error() == "unauthorized" {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You can only delete your own comment"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Comment deleted successfully"})
}
