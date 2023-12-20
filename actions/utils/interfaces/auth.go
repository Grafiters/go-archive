package interfaces

import (
	"archive/models"
	"crypto/rsa"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type LoginGoogle struct {
	RedirectUrl string `json:"redirect_url" binding:"required" swag:"form"`
	Code        string `json:"code" binding:"required" swag:"form"`
}

type GoogleData struct {
	Email     string `json:"email" binding:"optional" swag:"form"`
	Name      string `json:"name" binding:"optional" swag:"form"`
	GivenName string `json:"given_name" binding:"optional" swag:"form"`
	Picture   string `json:"picture" binding:"optional" swag:"form"`
}

type JwtAuth struct {
	UID      string   `json:"uid"`
	Email    string   `json:"email"`
	Audiance []string `json:"aud,omitempty"`

	jwt.StandardClaims
}

type KeyStore struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

type SessionAuth struct {
	User    models.User
	Expired uint32
}

type UserMe struct {
	UID      uuid.UUID `json:"id" db:"id"`
	Email    string    `json:"email" db:"email"`
	Username string    `json:"username" db:"username"`
}
