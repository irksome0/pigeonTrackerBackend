package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/irksome0/pigeonTracker/database"
	"github.com/irksome0/pigeonTracker/routes"
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
		AllowOrigins:     "http://localhost:3002, https://pigeon-tracker-2mb7v9d5k-irksome0s-projects.vercel.app, https://pigeon-tracker.vercel.app, https://pigeon-tracker-irksome0s-projects.vercel.app/",
		AllowHeaders:     "*",
	}))

	routes.Setup(app)

	app.Listen(":" + port)
}
