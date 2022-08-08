package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusrylmubarok/ism-api-golang/controllers"
	"github.com/gusrylmubarok/ism-api-golang/middlewares"
)

func Setup(app *fiber.App) {
	// Public API
	app.Post("/api/register", controllers.Registration)
	app.Post("/api/login", controllers.Login)
	
	// Private API for Request has been authenticated
	app.Use(middlewares.IsAuthenticated)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}