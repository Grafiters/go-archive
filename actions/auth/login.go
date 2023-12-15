package auth

import (
	"archive/actions/utils/interfaces"
	"archive/actions/utils/services"
	"archive/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// AuthLoginGoogle default implementation.
//	@Summary		Auth google login
//	@Description	This is the hello endpoint.
//  @Param params formData interfaces.LoginGoogle false "Auth Parameter"
//  @Tags			Auth
//	@Produce		json
//	@Router			/auth/google [post]
//	@Success		200	{object} services.ResponseMetaData
func AuthLoginGoogle(c buffalo.Context) error {

	loginGoogleData := interfaces.LoginGoogle{
		Code:        c.Param("code"),
		RedirectUrl: c.Param("redirect_url"),
	}

	err := c.Request().ParseForm()
	if err != nil {
		return c.Error(http.StatusUnprocessableEntity, err)
	}

	authProc, err := services.GoogleCallbackHandler(loginGoogleData)
	if err != nil {
		buildMessage, _ := services.BuildResponseHandler("unauthorized", err)
		return c.Render(http.StatusNonAuthoritativeInfo, r.JSON(buildMessage))
	}

	mustHash, err := services.HashPassword("12345678")
	if err != nil {
		buildMessage, _ := services.BuildResponseHandler("invalid_encryption", err)
		return c.Render(http.StatusNonAuthoritativeInfo, r.JSON(buildMessage))
	}

	data := &models.User{
		Email:    authProc.Email,
		Username: authProc.GivenName,
		Password: mustHash,
	}

	user, err := validateAndCreate(data)
	if err != nil {
		response, _ := services.BuildResponseHandler("validation", err)
		return c.Render(http.StatusNonAuthoritativeInfo, r.JSON(response))
	}

	generateToken, err := services.EncodeToken(user)
	if err != nil {
		response, _ := services.BuildResponseHandler("validation", err)
		return c.Render(http.StatusNonAuthoritativeInfo, r.JSON(response))
	}

	response, err := services.BuildResponseHandler("inserted", generateToken)
	if err != nil {
		return c.Render(http.StatusNoContent, r.JSON(authProc))
	}

	return c.Render(http.StatusOK, r.JSON(response))
}
