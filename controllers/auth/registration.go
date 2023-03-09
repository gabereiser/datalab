package auth

import "github.com/gofiber/fiber/v2"

func RegisterScreen(c *fiber.Ctx) error {
	return c.Render("views/index", fiber.Map{
		"Title": "Datalab",
		"User":  nil,
	}, "layouts/default")
}
func RegisterHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}
