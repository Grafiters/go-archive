package user

import (
	"github.com/gobuffalo/buffalo"
)

func Configuration(app *buffalo.App) {
	public := app.Group("/users")

	public.GET("/me", UserUserIndex)
}
