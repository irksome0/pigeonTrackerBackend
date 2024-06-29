package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/irksome0/pigeonTracker/controllers"
	"github.com/irksome0/pigeonTracker/middlewares"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/user", controllers.GetUser)
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/createPigeon", controllers.PostPigeon)

	app.Use(middlewares.IsAdministrator)
	app.Use("/api/checkIfAdmin", controllers.IsAdmin)
	app.Post("/api/users", controllers.GetAllUsers)
	app.Post("/api/delete", controllers.DeleteUser)
}
