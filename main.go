package main

import (
	"fmt"
	"os"

	"github.com/chirag3003/go-backend-template/controller"
	"github.com/chirag3003/go-backend-template/db"
	"github.com/chirag3003/go-backend-template/helpers"
	"github.com/chirag3003/go-backend-template/helpers/aws"
	"github.com/chirag3003/go-backend-template/middlewares"
	"github.com/chirag3003/go-backend-template/repository"
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

	// Setting up AWS
	aws.SetupAWS()

	// Creating a new Fiber app
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	// setting up repository
	repo := repository.Setup(client)

	// setting up middlewares
	middlewares.Setup(repo)

	// setting up routes and controllers
	routes.Setup(controller.Setup(repo), app)

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":  "A go backend template written by Chirag Bhalotia!",
			"author":   "Chirag Bhaloti",
			"github":   "https://github.com/chirag3003",
			"website":  "https://chirag.codes",
			"features": []string{"JWT", "CORS", "MongoDB", "S3", "Logger", "Environment Variables", "Go Modules", "Mongodb", "Docker", "Air", "Auto Reload", "Compose", "Deploy-script"},
		})
	})

	fmt.Println("Listening on port 5000")
	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
