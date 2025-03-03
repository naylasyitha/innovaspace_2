package middleware

import (
	"fmt"
	"innovaspace/internal/infra/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
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
	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		ctx.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
		return nil
	}

	bearerToken := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if bearerToken == "" {
		ctx.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
		return nil
	}

	token := bearerToken
	fmt.Println(token)

	id, err := m.jwt.ValidateToken(token)
	if err != nil {
		ctx.Status(401).JSON(fiber.Map{"message": "Unauthorized", "error": err.Error()})
		return nil
	}

	ctx.Locals("userId", id)
	return ctx.Next()
}
