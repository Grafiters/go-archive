package identify

import (
	"archive/actions/utils/interfaces"
	"archive/actions/utils/services"
	"archive/models"
	"fmt"
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

type Token struct {
	token string
}

var (
	interfaceJwt interfaces.JwtAuth
	tx           *pop.Connection
	err          error
)

func init() {
	tx, err = models.DB.NewTransaction()
	if err != nil {
		log.Fatal("Error Auth JWT Connection : ", err)
	}
}

func UserIdentify(c buffalo.Context) (*models.User, error) {
	token := services.ParsingTokenHeader(c)
	err := services.DecodeToken(token, &interfaceJwt)
	if err != nil {
		log.Fatal("Error decode jwt :", err)
		return nil, err
	}

	mustDetailUserAuth, err := userData(interfaceJwt)
	if err != nil {
		log.Fatal("Error validate jwt client :", err)
		return nil, err
	}

	return mustDetailUserAuth, nil
}

func userData(jwtDecode interfaces.JwtAuth) (*models.User, error) {
	mustUid := jwtDecode.UID
	if mustUid == "" {
		err := fmt.Errorf("Jwt auth not recognized")
		return nil, err
	}

	var user *models.User
	err := models.DB.Where("id = ?", mustUid).First(user)
	if err != nil {
		err := fmt.Errorf("User is not registered on system or Jwt Auth is invalid value")
		return nil, err
	}

	return user, nil
}

func UserDataDecode(email string) (*interfaces.UserMe, error) {
	user := &models.User{}
	err := models.DB.Where("users.email = ?", email).First(user)
	if err != nil {
		fmt.Println(err)
		err := fmt.Errorf("User is not registered on system or Jwt Auth is invalid value")
		return nil, err
	}

	convert := ConvertToUserMe(user)
	return convert, nil
}

func ConvertToUserMe(user *models.User) *interfaces.UserMe {
	return &interfaces.UserMe{
		UID:      user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}
