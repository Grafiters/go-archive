package session

import (
	"time"

	"github.com/Grafiters/archive/app/controller/helpers"
	"github.com/Grafiters/archive/app/models"
	"github.com/Grafiters/archive/app/payload"
	"github.com/Grafiters/archive/configs"
	"github.com/gofiber/fiber/v2"
)

// Auth godoc
// @Router /api/v2/session [post]
// @Summary Authenticate
// @Param body body payload.SessionGooglePayload true "session request"
// @Description Generate session user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} helpers.Response
// @Failure 422 {object} helpers.Errors
// @Failure 500 {object} helpers.Errors
func GoogleAuth(c *fiber.Ctx) error {
	errors := new(helpers.Errors)
	payload := new(payload.SessionGooglePayload)

	if len(c.Body()) <= 0 {
		var err error
		c.Status(500).JSON(helpers.Errors{
			Code:   500,
			Status: false,
			Errors: []string{"server.method.invalid_message_body"},
		})

		return err
	}

	if err := c.BodyParser(payload); err != nil {
		c.Status(500).JSON(helpers.Errors{
			Code:   500,
			Status: false,
			Errors: []string{"server.method.invalid_message_body"},
		})

		return err
	}

	helpers.Vaildate(payload, errors)
	if errors.Size() > 0 {
		return c.Status(422).JSON(errors)
	}

	if payload.AccessToken == "" && payload.Code == "" {
		return c.Status(422).JSON(helpers.Errors{
			Code:   422,
			Status: false,
			Errors: []string{"auth.google.must_use_one_from_that_two_method"},
		})
	}

	validate := payload.AccessTokenValidate(payload.AccessToken)
	if !validate {
		return c.Status(422).JSON(helpers.Errors{
			Code:   422,
			Status: false,
			Errors: []string{"auth.google.invalid_access_token"},
		})
	}

	validate = payload.CodeValidate(payload.Code)
	if !validate {
		return c.Status(422).JSON(helpers.Errors{
			Code:   422,
			Status: false,
			Errors: []string{"auth.google.invalid_code_or_redirect_value"},
		})
	}

	validate = helpers.IsURL(payload.RedirectURL)
	if !validate {
		return c.Status(422).JSON(helpers.Errors{
			Code:   422,
			Status: false,
			Errors: []string{"auth.google.invalid_redirect_url"},
		})
	}

	generateToken, err := configs.JwtConfig.GenerateTokenSession(&models.User{
		ID:        1,
		UID:       configs.Generate("UID"),
		Email:     "alone@gmail.com",
		GoogleID:  "1234566213",
		Password:  "1234566",
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	})

	if err != nil {
		return c.Status(500).JSON(helpers.Errors{
			Code:   422,
			Status: false,
			Errors: []string{"auth.google.auth_invalid_generate"},
		})
	}

	response := &helpers.Response{
		Code:   201,
		Status: true,
		Data:   generateToken,
	}

	return c.Status(200).JSON(response)
}
