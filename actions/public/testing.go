package public

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// PublicTestingIndex default implementation.
func PublicTestingIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Public Testing"}))
}
