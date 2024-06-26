package payload

import (
	"github.com/gookit/validate"
)

type SessionGooglePayload struct {
	AccessToken string `json:"access_token,omitempty" form:"access_token" validate:"optional|AccessTokenValidate"`
	Code        string `json:"code,omitempty" form:"code" validate:"optional|CodeValidate"`
	RedirectURL string `json:"redirect_url,omitempty" form:"redirect_url" validate:"optional|CodeValidate"`
}

type RegisterUserPayload struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type SessionLoginEmail struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (p SessionGooglePayload) Messages() map[string]string {
	invalid_message := "auth.google.invalid_{field}"

	return validate.MS{
		"required":            invalid_message,
		"AccessTokenValidate": "auth.google.invalid_access_token",
		"CodeValidate":        "auth.google.invalid_code_or_redirect_value",
	}
}

func (p RegisterUserPayload) Messages() map[string]string {
	invalid_message := "auth.register.invalid_{field}"

	return validate.MS{
		"required": invalid_message,
	}
}

func (p SessionLoginEmail) Messages() map[string]string {
	invalid_message := "auth.session.invalid_{field}"

	return validate.MS{
		"required": invalid_message,
	}
}

func (sg SessionGooglePayload) AccessTokenValidate(token string) bool {
	if sg.AccessToken == "" && sg.Code == "" {
		return false
	}

	return true
}

func (sg SessionGooglePayload) CodeValidate(token string) bool {
	if sg.Code != "" && sg.RedirectURL != "" && sg.AccessToken == "" {
		return true
	}

	return false
}
