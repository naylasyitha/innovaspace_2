package rest

import (
	"innovaspace/internal/app/user/usecase"
	"innovaspace/internal/domain/dto"
	"innovaspace/internal/validation"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	routerGroup.Post("/get-profile", userHandler.GetProfileByUsername)
	routerGroup.Patch("/update/:id", userHandler.Update)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	var register dto.Register

	if err := ctx.BodyParser(&register); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	file, _ := ctx.FormFile("profile_picture")

	if err := h.validator.Validate(register); err != nil {
		log.Printf("Validation error: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := h.usecase.Register(register, file)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
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

func (h *UserHandler) GetProfileByUsername(ctx *fiber.Ctx) error {
	var request struct {
		Username string `json:"username"`
	}

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.usecase.GetProfileByUsername(request.Username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(user)
}

func (h *UserHandler) Update(ctx *fiber.Ctx) error {
	userId, err := uuid.Parse(ctx.Params("user_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	var request dto.GetProfile
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err = h.usecase.UpdateUser(userId, request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}
