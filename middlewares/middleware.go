package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/irksome0/pigeonTracker/database"
	"github.com/irksome0/pigeonTracker/models"
	"github.com/irksome0/pigeonTracker/utils"

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

// func IsAdministrator(c *fiber.Ctx) error {
// 	var data map[string]string
// 	if err := c.BodyParser(&data); err != nil {
// 		log.Fatal(("Unable to parse body!(aboba)"))
// 	}
// 	var user models.User

// 	database.DB.Where("email=?", data["email"]).First(&user)

// 	if user.Id == 0 {
// 		c.Status(404)
// 		return c.JSON(fiber.Map{
// 			"message": "User with such email does not exist!",
// 		})
// 	}

//		c.Status(200)
//		if !user.Admin {
//			c.Status(401)
//			return c.JSON(fiber.Map{
//				"message": "User has no permission!",
//			})
//		}
//		return c.Next()
//	}
func IsAdministrator(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		// Log the error and return a 400 status code
		log.Println("Unable to parse body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User with such email does not exist!",
		})
	}

	if !user.Admin {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User has no permission!",
		})
	}

	return c.Next()
}
