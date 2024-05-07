package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/irksome0/blog/database"
	"github.com/irksome0/blog/models"
	"github.com/irksome0/blog/utils"

	"log"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := utils.ParseJwt(cookie); err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Authentication failed!",
		})
	}
	return c.Next()
}

func IsAdministrator(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		log.Fatal(("Unable to parse body!(Login)"))
	}
	var user models.User

	database.DB.Where("email=?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "User with such email does not exist!",
		})
	}

	c.Status(200)
	if user.Admin {
		c.Next()
		return c.JSON(fiber.Map{
			"message": "redirecting",
		})
	}
	return c.JSON(fiber.Map{
		"message": "User is not administrator",
	})
}
