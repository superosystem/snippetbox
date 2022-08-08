package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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

	// JWT Provider
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})
	// Generate Token
	token, err := claims.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	// Set Key into Cookie
	cookie := fiber.Cookie {
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	
	return c.JSON(fiber.Map {
		"message": "Login Succsses.",
	})
}

type Claims struct{
	jwt.RegisteredClaims
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func (token *jwt.Token) (interface{}, error) {
		return 	[]byte("secret"), nil
	})
	// User Account is Unauthenticated
	if err != nil || !token.Valid{
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated.",
		})
	}
	// Get all value from claims
	claims := token.Claims.(*Claims)
	// Get data user from Id or Issuer
	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	// Set value cookie and expiration time to unvalid 
	cookie := fiber.Cookie {
		Name:		"jwt",
		Value: 		"",
		Expires:	time.Now().Add(-time.Hour),
		HTTPOnly: 	true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logout is success.",
	})
}