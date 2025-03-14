package middleware

import (
	"fmt"
	"innovaspace/internal/infra/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type MiddlewareItf interface {
	Authentication(ctx *fiber.Ctx) error
}

type Middleware struct {
	jwt *jwt.JWT
}

func NewMiddleware(jwt *jwt.JWT) MiddlewareItf {
	return &Middleware{
		jwt: jwt,
	}
}

func (m *Middleware) Authentication(ctx *fiber.Ctx) error {
	authHeader := ctx.GetReqHeaders()["Authorization"]

	if authHeader == nil {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Autentikasi gagal",
			"errors": fiber.Map{
				"auth": "Anda belum login",
			},
		})
		return nil
	}

	if len(authHeader) < 1 {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
		return nil
	}

	bearerToken := authHeader[0]
	if bearerToken == "" {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		return nil
	}

	token := strings.Split(bearerToken, " ")[1]
	fmt.Println(token)

	id, err := m.jwt.ValidateToken(token)
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized", "error": err.Error()})
		return nil
	}

	ctx.Locals("userId", id)
	return ctx.Next()
}

func CorsMiddleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowHeaders:     "*",
		AllowMethods:     "*",
	}))
}
