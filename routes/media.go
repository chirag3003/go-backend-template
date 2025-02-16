package routes

import (
	"github.com/chirag3003/go-backend-template/middlewares"
	"github.com/gofiber/fiber/v3"
)

func mediaRoutes(app fiber.Router) {
	app.Post("/upload", conts.Media.Upload, middlewares.UploadFiles)
}
