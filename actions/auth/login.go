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

	data := &models.User{
		Email:    authProc.Email,
		Username: authProc.GivenName,
		Password: "12345678",
	}

	user, err := validateAndCreate(data)
	if err != nil {
		response, _ := services.BuildResponseHandler("validation", err)
		return c.Render(http.StatusNonAuthoritativeInfo, r.JSON(response))
	}

	response, err := services.BuildResponseHandler("inserted", user)
	if err != nil {
		return c.Render(http.StatusNoContent, r.JSON(authProc))
	}

	return c.Render(http.StatusOK, r.JSON(response))
}

func validateAndCreate(user *models.User) (*models.User, error) {
	tx, _ := models.DB.NewTransaction()

	check, err := user.FindUserByEmail(tx)
	if err != nil {
		return nil, err
	}

	if check == nil {
		check, err := createUser(user)
		if err != nil {
			return nil, err
		}

		return check, nil
	}

	return check, nil

}

func createUser(data *models.User) (*models.User, error) {
	tx, err := models.DB.NewTransaction()
	if err != nil {
		return nil, err
	}

	if err := tx.Create(data); err != nil {
		tx.TX.Rollback()
		return nil, err
	}

	tx.TX.Commit()

	return data, nil
}
