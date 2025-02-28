package rest

import (
	"innovaspace/internal/app/user/usecase"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/validation"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase   usecase.UserUsecaseItf
	validator validation.InputValidation
}

func NewUserHandler(routerGroup fiber.Router, userUsecase usecase.UserUsecaseItf, Validator validation.InputValidation) {
	userHandler := UserHandler{
		usecase:   userUsecase,
		validator: Validator,
	}

	routerGroup = routerGroup.Group("/users")
	routerGroup.Post("/register", userHandler.Register)
	routerGroup.Post("/login", userHandler.Login)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	var register dto.Register

	if err := ctx.BodyParser(&register); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if register.Email == "" || register.Password == "" || register.Username == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email, username, dan password wajib diisi",
		})
	}

	if err := h.usecase.Register(register); err != nil {
		return err
		// switch{
		// case errors.Is(err, usecase.ErrEmailExists):
		// 	return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
		// 		"error": "email sudah terdaftar",
		// 	})
		// case errors.Is(err, usecase.ErrInvalidFormat):
		// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		// 		"error": err.Error(),
		// 	})
		// default:
		// 	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 		"error": "internal server error",
		// 	})
		// }
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "success register user",
	})

}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	var login dto.Login

	if err := ctx.BodyParser(&login); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	token, err := h.usecase.Login(login)
	if err != nil {
		return err
		// switch{
		// case errors.Is(err, usecase.ErrUserNotFound):
		// 	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		// 		"error": "username atau password salah",
		// 	})
		// case errors.Is(err, usecase.ErrInvalidPassword):
		// 	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		// 		"error": "username atau password salah",
		// 	})
		// }
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success login user",
		"token":   token,
	})
}
