package middleware

import (
	"archive/actions/utils/identify"
	"archive/actions/utils/interfaces"
	"archive/actions/utils/services"
	"archive/models"
	"log"

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

func CurrentUser(c buffalo.Context) (*models.User, bool) {
	userData := c.Value("userData").(interfaces.JwtAuth)
	mustUID := userData.UID

	user, err := identify.UserDataDecode(mustUID)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	return user, true
}
