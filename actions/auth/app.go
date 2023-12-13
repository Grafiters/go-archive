package auth

import (
	"github.com/gobuffalo/buffalo"
)

// 	Auth default implementation.
//	@Summary		Auth handler
//	@Description	This is the auth endpoint.
//  @Tags			Auth
//	@Produce		json
//	@Router			/auth [post]
func Configuration(app *buffalo.App) {
	public := app.Group("/auth")

	public.POST("/google", AuthLoginGoogle)
}
