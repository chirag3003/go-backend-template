package controller

import (
	"github.com/chirag3003/go-backend-template/models"
	"github.com/chirag3003/go-backend-template/repository"
	"github.com/gofiber/fiber/v3"
)

type AuthController interface {
	Login(ctx fiber.Ctx) error
	Register(ctx fiber.Ctx) error
}

type authController struct {
	user repository.UserRepository
}

func newAuthController() *authController {
	return &authController{
		user: repo.User,
	}
}

func (a *authController) Login(ctx fiber.Ctx) error {
	//Parsing Body
	loginUserBody := models.LoginUserBody{}
	if err := ctx.Bind().JSON(&loginUserBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
	}

	//Fetching User with email
	user, _ := a.user.GetUserByEmail(ctx.Context(), loginUserBody.Email)
	//Checking if user exists
	if user == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}
	//Validating Password
	if !user.VerifyPassword(loginUserBody.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	//Generating JWT Token
	token, err := user.GenerateJWT()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}

	//Returning Response
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":      user.ID,
			"name":    user.Name,
			"email":   user.Email,
			"phoneNo": user.PhoneNo,
		},
	})
}

func (a *authController) Register(ctx fiber.Ctx) error {
	//Parsing Body
	registerUserBody := models.RegisterUserBody{}
	if err := ctx.Bind().JSON(&registerUserBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
	}

	//Checking if User already exists
	existingUser, err := a.user.GetUserByEmail(ctx.Context(), registerUserBody.Email)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}
	if existingUser != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	//Populating User
	user := &models.User{
		Name:  registerUserBody.Name,
		Email: registerUserBody.Email,
	}

	//Setting Password
	if err := user.SetPassword(registerUserBody.Password); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}

	//Creating User
	if err := a.user.CreateUser(ctx.Context(), user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Register Successful",
	})
}
