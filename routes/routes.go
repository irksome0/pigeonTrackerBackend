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

	app.Get("/api/user", controllers.GetUser)
	app.Post("/api/logout", controllers.Logout)

	// app.Use(middlewares.IsAdministrator)
	app.Post("/api/users", controllers.GetAllUsers)
}
