package sessionController

import (
	"github.com/Grafiters/archive/app/controller/helpers"
	"github.com/Grafiters/archive/app/payload"
	"github.com/Grafiters/archive/configs"
	"github.com/gofiber/fiber/v2"
)

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
		configs.Logger.Error("Error :", errors.Errors)
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

	response := &helpers.Response{
		Code:   201,
		Status: true,
		Data:   "auth.berhasil",
	}

	return c.Status(200).JSON(response)
}
