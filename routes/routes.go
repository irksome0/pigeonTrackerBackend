package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/irksome0/blog/controllers"
	"github.com/irksome0/blog/middlewares"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Post("/api/post", controllers.CreatePost)
}
