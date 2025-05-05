package middlewares

import (
	"fmt"
	"os"

	"github.com/chirag3003/go-backend-template/helpers"
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

		// If image optimization is required, use the OptimiseImage helper function
		open, err := helpers.OptimiseImage(*file)
		if err != nil {
			// Return an internal server error status if optimization fails
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		// Uncomment the following lines if image optimization is not required
		// This will directly open the file without any modifications
		// open, err := file.Open()
		// if err != nil {
		//     // Return an internal server error status if file opening fails
		//     return ctx.SendStatus(fiber.StatusInternalServerError)
		// }
		// defer open.Close() // Ensure the file reader is closed after use


		// Creating file key
		key := fmt.Sprintf("%s/%s%s", os.Getenv("S3_FOLDER"), name, file.Filename+".webp")

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
