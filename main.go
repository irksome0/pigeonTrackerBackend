package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/irksome0/blog/database"
	"github.com/irksome0/blog/routes"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file!")
	}

	port := os.Getenv("PORT")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3001",
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length",
	}))

	routes.Setup(app)

	app.Listen(":" + port)
}
