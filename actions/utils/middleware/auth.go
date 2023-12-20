package middleware

import (
	"archive/actions/utils/identify"
	"archive/actions/utils/interfaces"
	"archive/actions/utils/services"

	"github.com/gobuffalo/buffalo"
)

func Authenticate(c buffalo.Context) (any, bool) {
	token := services.ParsingTokenHeader(c)

	var jwtAuth interfaces.JwtAuth
	authUser := services.DecodeToken(token, &jwtAuth)

	if authUser != nil {
		return authUser, false
	}
	return jwtAuth, true
}

func CurrentUser(c buffalo.Context) (*interfaces.UserMe, bool) {
	userData := c.Value("userData").(interfaces.JwtAuth)
	mustUID := userData.Email

	user, err := identify.UserDataDecode(mustUID)
	if err != nil {
		return nil, false
	}
	return user, true
}
