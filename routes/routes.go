package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gusrylmubarok/ism-api-golang/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Registration)
	app.Post("/api/login", controllers.Login)
}