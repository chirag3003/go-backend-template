package routes

import (
	"github.com/chirag3003/go-backend-template/middlewares"
	"github.com/gofiber/fiber/v3"
)

func userRoutes(app fiber.Router) {

	app.Get("/me", middlewares.IsAuthenticated, conts.User.GetMe)
}
