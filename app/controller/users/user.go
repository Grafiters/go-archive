package users

import (
	"fmt"
	"time"

	"github.com/Grafiters/archive/app/controller/helpers"
	"github.com/Grafiters/archive/app/models"
	"github.com/Grafiters/archive/app/payload"
	"github.com/Grafiters/archive/configs"
	"github.com/Grafiters/archive/configs/response"
	"github.com/Grafiters/archive/configs/types"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var err error

// Register godoc
// @Router /api/v2/users/register [post]
// @Summary Register
// @Param body body payload.RegisterUserPayload true "session request"
// @Description Register user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Success
// @Failure 422 {object} response.Errors
// @Failure 500 {object} response.Errors
func RegisterUser(c *fiber.Ctx) error {
	errors := new(response.Errors)
	payload := new(payload.RegisterUserPayload)

	if len(c.Body()) <= 0 {
		c.Status(500).JSON(response.Errors{
			Code:   500,
			Status: false,
			Errors: []string{response.InvalidMessageBody},
		})

		return err
	}

	if err := c.BodyParser(payload); err != nil {
		c.Status(500).JSON(response.Errors{
			Code:   500,
			Status: false,
			Errors: []string{response.InvalidMessageBody},
		})

		return err
	}

	helpers.Vaildate(payload, errors)
	if errors.Size() > 0 {
		return c.Status(422).JSON(errors)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	user := &models.User{
		UID:       configs.Generate("UID"),
		Email:     payload.Email,
		Password:  string(hashedPassword),
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	}

	result := configs.DataBase.Create(user)

	if result.Error != nil {
		configs.Logger.Error("query -> ", result.Error.Error())
		c.Status(500).JSON(response.Errors{
			Code:   500,
			Status: false,
			Errors: []string{response.InternalServerError},
		})
	}

	var users int64
	errored := configs.DataBase.Table("users").Model(&models.User{}).Count(&users)
	if errored.Error != nil {
		configs.Logger.Error("query -> ", errored.Error.Error())
		c.Status(500).JSON(response.Errors{
			Code:   500,
			Status: false,
			Errors: []string{response.InternalServerError},
		})
	}

	if users == 1 {
		configs.DataBase.Table("users").Where("id = ?", user.ID).Updates(&models.User{
			Role: types.MemberRoleAdmin,
		})
	}

	generateToken, errorConfig := configs.JwtConfig.GenerateTokenSession(user)
	if errorConfig != nil {
		configs.Logger.Error("generate session -> ", errorConfig.Error())
		return c.Status(500).JSON(response.Errors{
			Code:   422,
			Status: false,
			Errors: []string{"auth.google.auth_invalid_generate"},
		})
	}

	response := &response.Success{
		Code:   201,
		Status: true,
		Data:   generateToken,
	}

	return c.Status(200).JSON(response)
}
