package middlewares

import (
	"fmt"
	"os"

	"github.com/chirag3003/go-backend-template/models"
	"github.com/gofiber/fiber/v3"
	"github.com/matoous/go-nanoid/v2"
)

func UploadFiles(ctx fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["files"]
	var fileUrls []string

	for _, file := range files {
		//Generating unique file name
		name, err := gonanoid.New(10)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		// open file reader
		open, err := file.Open()
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		defer open.Close()

		// Creating file key
		key := fmt.Sprintf("%s/%s%s", os.Getenv("S3_FOLDER"), name, file.Filename)

		// upload file to s3
		res, err := repo.S3.Upload(ctx.Context(), key, open)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		// append file url to fileUrls
		fileUrls = append(fileUrls, res.Location)

		// uploding file details to mongodb
		err = repo.Media.CreateMedia(ctx.Context(), &models.Media{
			Key:  key,
			Url:  res.Location,
			Etag: *res.ETag,
			Mime: file.Header.Get("Content-Type"),
			Size: file.Size,
		})
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
	}

	ctx.Locals("files", fileUrls)
	return ctx.Next()
}
