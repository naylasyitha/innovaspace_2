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
	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		return nil
	}

	bearerToken := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if bearerToken == "" {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
		return nil
	}

	token := bearerToken
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
		AllowOrigins:     "https://innovaspace-intern.vercel.app",
		AllowCredentials: true,
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS, PATCH",
	}))
}
