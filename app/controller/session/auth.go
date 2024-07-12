package session

import (
	"time"

	"github.com/Grafiters/archive/app/controller/helpers"
	"github.com/Grafiters/archive/app/models"
	"github.com/Grafiters/archive/app/payload"
	"github.com/Grafiters/archive/configs"
	"github.com/Grafiters/archive/configs/response"
	"github.com/Grafiters/archive/lib/google"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Auth godoc
// @Router /api/v2/session [post]
// @Summary Authenticate
// @Param body body payload.SessionGooglePayload true "session request"
// @Description Generate session user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Success
// @Failure 422 {object} response.Errors
// @Failure 500 {object} response.Errors
func GoogleAuth(c *fiber.Ctx) error {
	errors := new(response.Errors)
	payload := new(payload.SessionGooglePayload)

	if len(c.Body()) <= 0 {
		var err error
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

	if payload.AccessToken == "" && payload.Code == "" {
		return c.Status(422).JSON(response.Errors{
			Code:   422,
			Status: false,
			Errors: []string{"auth.google.must_use_one_from_that_two_method"},
		})
	}

	validate := payload.AccessTokenValidate(payload.AccessToken)
	if !validate {
		return c.Status(422).JSON(response.Errors{
			Code:   422,
			Status: false,
			Errors: []string{"auth.google.invalid_access_token"},
		})
	}

	validate = payload.CodeValidate(payload.Code)
	if !validate {
		return c.Status(422).JSON(response.Errors{
			Code:   422,
			Status: false,
			Errors: []string{"auth.google.invalid_code_or_redirect_value"},
		})
	}

	validate = helpers.IsURL(payload.RedirectURL)
	if !validate {
		return c.Status(422).JSON(response.Errors{
			Code:   422,
			Status: false,
			Errors: []string{"auth.google.invalid_redirect_url"},
		})
	}

	googleCallback, err := google.GoogleCallbackHandler(*payload)
	if err != nil {
		return c.Status(422).JSON(response.Errors{
			Code:   422,
			Status: false,
			Errors: []string{err.Error()},
		})
	}

	var member *models.User
	configs.DataBase.Where("google_id = ?", googleCallback.GoogleID).FirstOrCreate(&member, &models.User{
		UID:       configs.Generate("UID"),
		Email:     googleCallback.Email,
		GoogleID:  googleCallback.GoogleID,
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	})

	generateToken, err := configs.JwtConfig.GenerateTokenSession(member)
	if err != nil {
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

	return c.Status(201).JSON(response)
}

// Auth godoc
// @Router /api/v2/session [post]
// @Summary Authenticate
// @Param body body payload.SessionLoginEmail true "session request"
// @Description Generate session user
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Success
// @Failure 422 {object} response.Errors
// @Failure 500 {object} response.Errors
func CreateSession(c *fiber.Ctx) error {
	errors := new(response.Errors)
	payload := new(payload.SessionLoginEmail)

	if len(c.Body()) <= 0 {
		var err error
		c.Status(500).JSON(response.Errors{
			Code:   500,
			Status: false,
			Errors: []string{response.InvalidMessageBody},
		})

		return err
	}

	if err := c.BodyParser(payload); err != nil {
		return c.Status(500).JSON(response.Errors{
			Code:   500,
			Status: false,
			Errors: []string{response.InvalidMessageBody},
		})
	}

	helpers.Vaildate(payload, errors)
	if errors.Size() > 0 {
		return c.Status(422).JSON(errors)
	}

	var member *models.User
	result := configs.DataBase.Table("users").Where("email = ?", payload.Email).First(&member)
	if result.Error != nil {
		c.Status(422).JSON(response.Errors{
			Code:   422,
			Status: false,
			Errors: []string{"auth.session.email_or_password_invalid"},
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(payload.Password)); err != nil {
		return c.Status(422).JSON(response.Errors{
			Code:   422,
			Status: false,
			Errors: []string{"auth.session.email_or_password_invalid"},
		})
	}

	generateToken, err := configs.JwtConfig.GenerateTokenSession(member)
	if err != nil {
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

	return c.Status(201).JSON(response)
}
