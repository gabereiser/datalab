package auth

import (
	"errors"

	"github.com/gabereiser/datalab/log"
	"github.com/gabereiser/datalab/security"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RefreshHandler(c *fiber.Ctx) error {
	if token, ok := c.Locals("_auth").(*jwt.Token); ok {
		claims := token.Claims.(jwt.MapClaims)
		log.Info("[INFO] jwt key %v %v", token, claims)

		t := security.RefreshToken(token)
		return c.JSON(fiber.Map{
			"token": t,
		})

	} else {
		c.Render("views/403", fiber.Map{
			"Status":  403,
			"Message": "Unauthorized",
		})
		return errors.New("unauthorized")
	}
}

func LoginHandler(c *fiber.Ctx) error {
	return nil
}
