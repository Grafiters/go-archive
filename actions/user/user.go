package user

import (
	"archive/actions/utils/middleware"
	"archive/actions/utils/services"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// 	UserUserIndex default implementation.
//	@Tags			User
// 	@Param Authorization header string true "Bearer YOUR_ACCESS_TOKEN"
//  @Tags			Auth
//	@Produce		json
//	@Router			/users/me [get]
func UserUserIndex(c buffalo.Context) error {
	currentUser, err := middleware.CurrentUser(c)
	if !err {
		buildMessage, _ := services.BuildResponseHandler("unauthorized", "")
		return c.Render(http.StatusNonAuthoritativeInfo, r.JSON(buildMessage))
	}

	response, errRes := services.BuildResponseHandler("success", currentUser)
	if errRes != nil {
		return c.Render(http.StatusNoContent, r.JSON(err))
	}
	return c.Render(http.StatusOK, r.JSON(response))
}
