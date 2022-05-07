package main

import (
	"github.com/gusrylmubarok/goadminproject/database"
	"github.com/gusrylmubarok/goadminproject/routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
