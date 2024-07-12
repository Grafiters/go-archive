package role

import (
	"strings"

	"github.com/Grafiters/archive/app/models"
	"github.com/Grafiters/archive/configs"
	"github.com/Grafiters/archive/configs/response"
	"github.com/gofiber/fiber/v2"
)

func AdminVaildator(c *fiber.Ctx) error {
	CurrentUser := c.Locals("CurrentUser").(*models.User)
	var permissions []*models.Permission

	configs.DataBase.Table("permissions").Where("role = ?", CurrentUser.Role).Find(&permissions)
	if len(permissions) > 0 {
		for _, i := range permissions {
			if i.Path == c.Path() && i.Action == c.Method() && strings.ToLower(i.Verb) != "accept" {
				return c.Status(401).JSON(&response.Errors{
					Code:   401,
					Status: false,
					Errors: []string{response.AuthzInvalidPermission},
				})
			}
		}
	}

	return c.Next()
}
