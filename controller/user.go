package controller

import (
	"fmt"

	"github.com/chirag3003/go-backend-template/models"
	"github.com/chirag3003/go-backend-template/repository"
	"github.com/gofiber/fiber/v3"
)

type UserController interface {
	GetMe(ctx fiber.Ctx) error
}

type userController struct {
	user repository.UserRepository
}

func newUserController() *userController {
	return &userController{user: repo.User}
}

func (u *userController) GetMe(ctx fiber.Ctx) error {
	user, ok := ctx.Locals("user").(*models.User)
	if !ok {
		fmt.Print("Not OK")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return ctx.JSON(user)
}
