package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusrylmubarok/ism-api-golang/database"
	"github.com/gusrylmubarok/ism-api-golang/routes"
)

func main() {
	database.Connection()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}