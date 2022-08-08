package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusrylmubarok/ism-api-golang/configuration"
)

func main() {

	configuration.Connection()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Helllo, World!")
	})

	app.Listen(":3000")
}