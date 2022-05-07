package controllers

import (
	"github.com/gusrylmubarok/goadminproject/database"
	"github.com/gusrylmubarok/goadminproject/models"
	"github.com/gofiber/fiber"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
