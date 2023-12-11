package public

import (
	"archive/actions/public/meme"

	"github.com/gobuffalo/buffalo"
)

func Configuration(app *buffalo.App) {
	public := app.Group("/public")

	public.GET("/", PublicTestingIndex)
	meme.Configuration(public)
}
