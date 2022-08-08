package middlewares

import  (
	"github.com/gofiber/fiber/v2"
	"github.com/gusrylmubarok/ism-api-golang/util"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	// User Account is Unauthenticated
	if _, err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated.",
		})
	}
	return c.Next()
}