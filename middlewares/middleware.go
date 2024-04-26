package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/irksome0/blog/utils"
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
