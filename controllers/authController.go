package controllers

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/irksome0/blog/database"
	"github.com/irksome0/blog/models"
	"github.com/irksome0/blog/utils"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}
func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User
	if err := c.BodyParser(&data); err != nil {
		log.Fatal(("Unable to parse body!(Registration)"))
	}

	// VALIDATION
	// Password length
	if len(data["password"].(string)) <= 8 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Your password must be 8 characters at least!",
		})
	}
	// Email correctness
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid email!",
		})
	}

	// Check if user is already created
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "User with such email already exists!",
		})
	}

	// Filling user data
	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
		Phone:     data["phone"].(string),
	}
	// Cyphering password
	user.SetPassword(data["password"].(string))

	// Adding user to DB
	database.DB.Create(&user)

	// Response
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "User has been created!",
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {
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
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect password!",
		})
	}
	token, err := utils.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(500)
		return nil
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Successful login!",
		"user":    user,
	})
}
