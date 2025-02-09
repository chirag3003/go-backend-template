package main

import (
	"github.com/chirag3003/go-backend-template/controller"
	"github.com/chirag3003/go-backend-template/db"
	"github.com/chirag3003/go-backend-template/helpers"
	"github.com/chirag3003/go-backend-template/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Loading Environment Variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Verifying Environment Variables
	helpers.VerifyENV()

	// Connecting to MongoDB
	client := db.ConnectMongo()
	defer client.Close()

	// Creating a new Fiber app
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	// setting up routes and controllers
	routes.Setup(controller.Setup(client), app)

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":  "A go backend template written by Chirag Bhalotia!",
			"author":   "Chirag Bhalotia",
			"github":   "https://github.com/chirag3003",
			"website":  "https://chirag.codes",
			"features": []string{"JWT", "CORS", "MongoDB", "S3", "Logger", "Environment Variables"},
		})
	})
}
