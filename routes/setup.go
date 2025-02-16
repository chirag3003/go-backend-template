package routes

import (
	"github.com/chirag3003/go-backend-template/controller"
	"github.com/gofiber/fiber/v3"
)

var conts *controller.Controllers

func Setup(controllers *controller.Controllers, app *fiber.App)  {
	conts = controllers


  //Setting Routes
  userRoutes(app.Group("/user"))
  authRoutes(app.Group("/auth"))
  mediaRoutes(app.Group("/media"))  
}
