package controller

import "github.com/gofiber/fiber/v3"

type MediaController interface {
	Upload(ctx fiber.Ctx) error
}

type mediaController struct {
}

func newMediaController() MediaController {
	return &mediaController{}
}

func (m *mediaController) Upload(ctx fiber.Ctx) error {
	files, ok := ctx.Locals("files").([]string)
	if !ok {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Upload Successful",
		"files":   files,
	})
}
