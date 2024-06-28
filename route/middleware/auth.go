package middleware

import (
	"strings"
	"time"

	"github.com/Grafiters/archive/app/models"
	"github.com/Grafiters/archive/configs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	IDF string `json:"idf"`
	jwt.StandardClaims
}

var (
	AuthzInvalidSession = "authz.invalid_session"
	JwtDecodeAndVerify  = "jwt.decode_and_verify"
	ServerInternalError = "server.internal_error"
)

func Authenticate(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if len(token) == 0 {
		return c.Status(401).JSON(fiber.Map{
			"errors": []string{AuthzInvalidSession},
		})
	}

	token = strings.Replace(token, configs.Prefix, "", -1)

	data := &models.User{
		ID:        1,
		UID:       configs.Generate("UID"),
		Email:     "alone@gmail.com",
		GoogleID:  "1234566213",
		Password:  "1234566",
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	}

	err := configs.JwtConfig.DecodeTokenSession(token, data)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"errors": []string{JwtDecodeAndVerify},
		})
	}
	c.Locals("CurrentUser", data)

	return c.Next()
}
