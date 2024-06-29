package controllers

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/irksome0/pigeonTracker/database"
	"github.com/irksome0/pigeonTracker/models"
	"github.com/irksome0/pigeonTracker/utils"
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
		Username: strings.TrimSpace(data["username"].(string)),
		Email:    strings.TrimSpace(data["email"].(string)),
		Admin:    false,
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
		SameSite: "None",
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	c.Status(200)
	return c.JSON(fiber.Map{
		"message":       "Successful login!",
		"access_token":  token,
		"expires_token": cookie.Expires,
		"user":          user,
	})
}

func GetUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	claims, err := utils.ParseJwt(cookie)
	if err != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"message": "Authentication failed!",
			"error":   err,
		})
	}

	var user models.User
	database.DB.Where("id=?", claims).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		SameSite: "None",
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Successful logout!",
	})
}

func IsAdmin(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		log.Fatal(("Unable to parse body!(Login)"))
	}
	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)
	if user.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User with such email does not exist!",
		})
	}

	if !user.Admin {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Not admin",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Admin",
	})
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	result := make([]models.UserBasic, 0, len(users))
	var user models.User
	for _, user = range users {
		var temp = models.UserBasic{
			Id:    user.Id,
			Name:  user.Username,
			Email: user.Email,
		}
		result = append(result, temp)
	}
	return c.JSON(users)
}

func DeleteUser(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		log.Fatal(("Unable to parse body!(Login)"))
	}
	var user models.User
	database.DB.Where("email=?", data["targetEmail"]).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "User with such email does not exist!",
		})
	}

	database.DB.Where("email=?", user.Email).Delete(&user)
	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "The record has been deleted",
	})
}
