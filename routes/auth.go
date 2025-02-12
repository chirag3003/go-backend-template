package routes

import "github.com/gofiber/fiber/v3"

func authRoutes(app fiber.Router) {
  app.Post("/login", conts.Auth.Login)
  app.Post("/register", conts.Auth.Register)
}
