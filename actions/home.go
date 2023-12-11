package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
//	@Summary		Welcome to Buffalo
//	@Description	This is the root endpoint.
//	@Tags			Default
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/ [get]
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Welcome to Buffalo!"}))
}
