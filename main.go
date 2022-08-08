package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gusrylmubarok/ism-api-golang/database"
	"github.com/gusrylmubarok/ism-api-golang/routes"
)

func main() {
	// Load Database Connection 
	database.Connection()
	// Use Fiber
	app := fiber.New()
	// User CORS
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	// Instance Router
	routes.Setup(app)
	// Setup Server Port
	app.Listen(":3000")	
}