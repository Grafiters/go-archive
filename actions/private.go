package actions

import (
	"archive/actions/user"

	"archive/actions/utils/middleware"

	"github.com/gobuffalo/buffalo"
)

func MiddlewareAuth(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		userData, isAuthenticated := middleware.Authenticate(c)

		if !isAuthenticated {
			return c.Render(401, r.JSON(map[string]interface{}{"error": "Missing or invalid token"}))
		}

		c.Set("userData", userData)

		return next(c)
	}
}

func PrivateRouteConfiguration(app *buffalo.App) {
	private := app.Group("/")
	private.Use(MiddlewareAuth)
	user.Configuration(private)
}
