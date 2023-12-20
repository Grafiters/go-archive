package actions

import (
	"archive/actions/auth"
	"archive/actions/public"
	"archive/actions/public/meme"
	"archive/actions/utils/interfaces"

	"github.com/gobuffalo/buffalo"
)

type ResponseData struct {
	Meta interfaces.MetaData
	Data map[string]interface{}
}

func RouteConfiguration(app *buffalo.App) {
	public.Configuration(app)
	auth.Configuration(app)
	meme.Configuration(app)
	PrivateRouteConfiguration(app)
}
