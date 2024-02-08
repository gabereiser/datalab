package security

import (
	"strings"

	"github.com/gabereiser/datalab/config"
	"github.com/gabereiser/datalab/data"
	"github.com/gabereiser/datalab/log"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func password_equals(dbpass string, inpass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbpass), []byte(inpass)) == nil
}

func hash(inpass string) string {
	d, _ := bcrypt.GenerateFromPassword([]byte(inpass), 20)
	return string(d)
}

func NewAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if token, ok := c.Locals("_auth").(*jwt.Token); ok {
			claims := token.Claims.(jwt.MapClaims)
			log.Info("jwt key %v %v", token, claims)
			return c.Next()
		} else {
			log.Info("Unauthorized view")
			return c.Render("views/401", fiber.Map{
				"Status":  401,
				"Message": "Unauthorized",
			}, "index")
		}
	}
}

func Login(email string, password string) *jwt.Token {
	a := data.FindAccount(email)
	if a != nil && password_equals(a.Password, password) {
		return makeToken(a)
	}
	return nil
}

func RefreshToken(token *jwt.Token) *jwt.Token {
	if verify(token) {
		claims := token.Claims.(jwt.MapClaims)
		id := claims["id"].(string)
		email := claims["email"].(string)
		a := data.GetAccount(data.NewID(&id))
		if a != nil {
			if a.Email == email {
				return makeToken(a)
			}
		}
	}
	return nil
}

func verify(token *jwt.Token) bool {
	m := token.Method
	return m.Verify("", token.Signature, []byte(config.Config.SecretKey)) == nil
}

func makeToken(account *data.AccountModel) *jwt.Token {
	claims := jwt.MapClaims{
		"id":    idString(account.ID),
		"email": account.Email,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
}

func idString(id data.ID) string {
	return strings.ReplaceAll(id.String(), "-", "")
}
