package public

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// PublicTestingIndex default implementation.
//	@Summary		Say hello to Buffalo
//	@Description	This is the hello endpoint.
//	@Param			params	query	interfaces.OrderingParam	false	"User parameters"
//	@Param			params	query	interfaces.PaginationParam	false	"User parameters"
//	@Tags			Public
//	@Produce		json
//	@Success		200	{object}	interfaces.PaginationParam
//	@Router			/public [get]
func PublicTestingIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Public Testing"}))
}
