package auth

import (
	"archive/actions/utils/interfaces"
	"archive/actions/utils/services"
	"fmt"
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
		return c.Error(http.StatusInternalServerError, err)
	}
	fmt.Println(c.Request())
	authProc, err := services.GoogleCallbackHandler(loginGoogleData)
	if err != nil {
		return c.Render(http.StatusNoContent, r.JSON(authProc))
	}

	return c.Render(http.StatusOK, r.JSON(authProc))
}
