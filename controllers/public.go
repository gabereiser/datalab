package controllers

import (
	"github.com/gabereiser/datalab/controllers/auth"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {
	return auth.RegisterHandler(c)
}

func IndexScreen(c *fiber.Ctx) error {
	return c.Render("views/index", fiber.Map{
		"Title": "Datalab",
		"User":  nil,
	}, "layouts/default")
}

func PrivacyScreen(c *fiber.Ctx) error {
	return c.Render("views/index", fiber.Map{
		"Title": "Datalab",
		"User":  nil,
	}, "layouts/default")
}

func AboutScreen(c *fiber.Ctx) error {
	return c.Render("views/index", fiber.Map{
		"Title": "Datalab",
		"User":  nil,
	}, "layouts/default")
}

func TermsScreen(c *fiber.Ctx) error {
	return c.Render("views/index", fiber.Map{
		"Title": "Datalab",
		"User":  nil,
	}, "layouts/default")
}

func LoginScreen(c *fiber.Ctx) error {
	return c.Render("views/index", fiber.Map{
		"Title": "Datalab",
		"User":  nil,
	}, "layouts/default")
}
