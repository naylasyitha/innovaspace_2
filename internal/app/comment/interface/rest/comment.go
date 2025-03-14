package rest

import (
	"fmt"
	"innovaspace/internal/app/comment/usecase"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/middleware"

	"github.com/gofiber/fiber/v2"

	"github.com/google/uuid"
)

type CommentHandler struct {
	CommentUsecase usecase.CommentUsecaseItf
	middleware     middleware.MiddlewareItf
}

func NewCommentHandler(routerGroup fiber.Router, commentUsecase usecase.CommentUsecaseItf, middleware middleware.MiddlewareItf) {
	CommentHandler := CommentHandler{
		CommentUsecase: commentUsecase,
		middleware:     middleware,
	}

	routerGroup = routerGroup.Group("/comments")
	routerGroup.Post("/create-comment", CommentHandler.middleware.Authentication, CommentHandler.CreateComment)
	routerGroup.Patch("/update-comment/:id", CommentHandler.middleware.Authentication, CommentHandler.UpdateComment)
	routerGroup.Delete("/delete-comment/:id", CommentHandler.middleware.Authentication, CommentHandler.DeleteComment)
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

func (h *CommentHandler) CreateComment(ctx *fiber.Ctx) error {
	var input dto.CreateCommentRequest
	if err := ctx.BodyParser(&input); err != nil {
		return errorResponse(ctx, fiber.StatusBadRequest,
			"Permintaan tidak valid", "Format request tidak sesuai")
	}

	userId := ctx.Locals("userId").(uuid.UUID)
	input.UserId = userId

	comment, err := h.CommentUsecase.CreateComment(input)
	if err != nil {
		return errorResponse(ctx, fiber.StatusInternalServerError,
			"Gagal membuat comment", err.Error())
	}

	return successResponse(ctx, fiber.StatusCreated, "Comment berhasil dibuat",
		fiber.Map{
			"comment": fiber.Map{
				"comment_id":   comment.CommentId,
				"thread_id":    comment.ThreadId,
				"user_id":      comment.UserId,
				"isi_komentar": comment.IsiKomentar,
			},
		})
}

func (h *CommentHandler) UpdateComment(ctx *fiber.Ctx) error {
	fmt.Println("Received ID:", ctx.Params("id"))
	commentId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return errorResponse(ctx, fiber.StatusBadRequest,
			"ID comment tidak valid", "Format ID tidak valid")
	}

	userId := ctx.Locals("userId").(uuid.UUID)
	var input dto.UpdateCommentRequest
	if err := ctx.BodyParser(&input); err != nil {
		return errorResponse(ctx, fiber.StatusBadRequest,
			"Format request tidak valid", "Format JSON tidak sesuai")
	}

	updatedComment, err := h.CommentUsecase.UpdateComment(userId, commentId, input)
	if err != nil {
		switch err.Error() {
		case "not allowed to update comment":
			return errorResponse(ctx, fiber.StatusForbidden,
				"Comment gagal diperbarui", "Pengguna tidak memiliki akses")
		case "request not valid":
			return errorResponse(ctx, fiber.StatusBadRequest,
				"Comment gagal diperbarui", "Isi komentar wajib diisi")
		case "comment not found":
			return errorResponse(ctx, fiber.StatusNotFound,
				"Comment tidak ditemukan", "")
		default:
			return errorResponse(ctx, fiber.StatusInternalServerError,
				"Comment gagal diperbarui", err.Error())
		}
	}

	return successResponse(ctx, fiber.StatusOK, "Berhasil memperbarui data", updatedComment)
}

func (h *CommentHandler) DeleteComment(ctx *fiber.Ctx) error {
	commentId, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return errorResponse(ctx, fiber.StatusBadRequest,
			"ID comment tidak valid", "Format ID tidak valid")
	}

	userId := ctx.Locals("userId").(uuid.UUID)
	if err := h.CommentUsecase.DeleteComment(userId, commentId); err != nil {
		switch err.Error() {
		case "not allowed to delete comment":
			return errorResponse(ctx, fiber.StatusForbidden,
				"Comment gagal dihapus", "Pengguna tidak memiliki akses")
		case "comment not found":
			return errorResponse(ctx, fiber.StatusNotFound,
				"Comment tidak ditemukan", "")
		default:
			return errorResponse(ctx, fiber.StatusInternalServerError,
				"Comment gagal dihapus", err.Error())
		}
	}

	return successResponse(ctx, fiber.StatusOK, "Comment berhasil dihapus",
		fiber.Map{"id": commentId.String()})
}
