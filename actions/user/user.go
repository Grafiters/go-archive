package user

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// 	UserUserIndex default implementation.
//	@Tags		User
// @Param Authorization header string true "Bearer YOUR_ACCESS_TOKEN"
//	@Router		/users/me [get]
func UserUserIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Public Testing"}))
}
