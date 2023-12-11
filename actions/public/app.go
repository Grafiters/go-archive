package public

import "github.com/gobuffalo/buffalo"

func Configuration(app *buffalo.App) {
	public := app.Group("/public")

	public.GET("/", PublicTestingIndex)
}
