package meme

import "github.com/gobuffalo/buffalo"

func Configuration(app *buffalo.App) {
	memeRoute := app.Group("/meme")

	memeRoute.GET("/", PublicMemeIndex)
}
