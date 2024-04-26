package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/irksome0/blog/database"
	"github.com/irksome0/blog/models"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		log.Fatal(("Unable to parse body!(Post)"))
	}
	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload!",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Post has been created!",
	})
}
