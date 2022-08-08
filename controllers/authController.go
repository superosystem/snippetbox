package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusrylmubarok/ism-api-golang/database"
	"github.com/gusrylmubarok/ism-api-golang/models"
	"golang.org/x/crypto/bcrypt"
)

func Registration(c *fiber.Ctx) error {
	// Catch data from request
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	// Check Password Request have to match
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password do not match.",
		})
	}
	// Hash password with BCrypt Method
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	// Set Request to Model
	user := models.User{
		FirstName:	data["first_name"],
		LastName: 	data["last_name"],
		Email: 		data["email"],
		Password: 	password,
	}
	// Save to database
	database.DB.Create(&user)
	// Send Response
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	// Catch the Request
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}
	// Check email account has been registration
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Account is not found.",
		})
	}
	// Compare password account
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password is incorrect",
		})
	}

	return c.JSON(user)
}