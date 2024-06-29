// CODE  BELOW IS NOT USED ANYWHERE WITHING THE APPLICATION
// IT'S NOT FINISHED AND DOESN'T WORK PROPERLY

package controllers

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/irksome0/pigeonTracker/database"
	"github.com/irksome0/pigeonTracker/models"
)

func PostPigeon(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		log.Fatal("Unable to parse body!(Pigeon creation)")
	}

	code := models.GeneratePigeonCode()

	if data["yearOfBirth"] == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid year of birth!",
		})
	}
	gender := strings.TrimSpace(data["gender"].(string))
	fmt.Println(gender)
	if !strings.Contains("FMfm", gender) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid gender value!",
		})
	}
	pigeonData := models.Pigeon{
		PigeonCode:  code,
		YearOfBirth: data["yearOfBirth"].(int),
		Gender:      gender,
		Colour:      strings.TrimSpace(data["colour"].(string)),
		Mother:      data["mother"].(int),
		Father:      data["father"].(int),
		OwnerEmail:  strings.TrimSpace(data["ownerEmail"].(string)),
	}

	database.DB.Create(&pigeonData)
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "User has been created!",
		"pigeon":  pigeonData,
	})
}
