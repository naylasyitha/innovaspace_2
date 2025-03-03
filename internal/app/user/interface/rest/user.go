package rest

import (
	"innovaspace/internal/app/user/usecase"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/validation"
	"log"
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
		log.Printf("Error parsing request body: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.validator.Validate(register); err != nil {
		log.Printf("Validation error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.usecase.Register(register); err != nil {
		return err
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

	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success login user",
		"token":   token,
	})
}
