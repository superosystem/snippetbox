package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusrylmubarok/ism-api-golang/configuration"
	"github.com/gusrylmubarok/ism-api-golang/routes"
)

func main() {

	configuration.Connection()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}