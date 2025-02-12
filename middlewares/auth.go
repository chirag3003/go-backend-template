package middlewares

import (
	"github.com/chirag3003/go-backend-template/models"
	"github.com/gofiber/fiber/v3"
)

type AuthHeaders struct {
	Authorization string `header:"authorization"`
}

func IsAuthenticated(ctx fiber.Ctx) error {
	authHeaders := AuthHeaders{}
	if err := ctx.Bind().Header(authHeaders); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	user := &models.User{}
	user, err := user.ParseJWT(authHeaders.Authorization)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	ctx.Locals("user", user)
	return ctx.Next()
}
