package middleware

import (
	"strings"

	"github.com/Grafiters/archive/app/models"
	"github.com/Grafiters/archive/configs"
	"github.com/Grafiters/archive/configs/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	IDF string `json:"idf"`
	jwt.StandardClaims
}

func Authenticate(c *fiber.Ctx) error {
	var (
		jwtMember Auth
		member    *models.User
	)
	token := c.Get("Authorization")
	if len(token) == 0 {
		return c.Status(401).JSON(fiber.Map{
			"errors": []string{response.AuthzInvalidSession},
		})
	}

	token = strings.Replace(token, configs.Prefix, "", -1)

	err := configs.JwtConfig.DecodeTokenSession(token, jwtMember)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"errors": []string{response.JwtDecodeAndVerify},
		})
	}

	configs.DataBase.Where("uid = ?", jwtMember.IDF).First(&member)
	if member == nil {
		return c.Status(401).JSON(&response.Errors{
			Code:   401,
			Status: false,
			Errors: []string{response.JwtDecodeAndVerify},
		})
	}

	c.Locals("CurrentUser", member)

	return c.Next()
}
