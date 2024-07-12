package account

import (
	"github.com/Grafiters/archive/app/models"
	"github.com/Grafiters/archive/configs/response"
	"github.com/gofiber/fiber/v2"
)

// Auth godoc
// @Router /api/v2/account [get]
// @Summary Account
// @Description Get account detail
// @Tags Account
// @Accept  json
// @Produce  json
// @Success 200 {object} helpers.Response{data=models.User}
// @Failure 422 {object} helpers.Errors
// @Failure 500 {object} helpers.Errors
func GetUsersMe(c *fiber.Ctx) error {
	CurrentUser := c.Locals("CurrentUser").(*models.User)

	return c.Status(200).JSON(&response.Success{
		Code:   200,
		Status: true,
		Data:   CurrentUser,
	})
}
