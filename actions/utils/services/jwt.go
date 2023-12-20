package services

import (
	"archive/actions/utils/interfaces"
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ParseAndValidate(token string, key *rsa.PublicKey) (interfaces.JwtAuth, error) {
	auth := interfaces.JwtAuth{}

	_, err := jwt.ParseWithClaims(token, &auth, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	return auth, err
}
func appendClaims(defaultClaims, customClaims jwt.MapClaims) jwt.MapClaims {
	if defaultClaims == nil {
		return customClaims
	}

	if customClaims == nil {
		return defaultClaims
	}

	for k, v := range customClaims {
		defaultClaims[k] = v
	}

	return defaultClaims
}
func ForgeToken(uid string, customClaims jwt.MapClaims, key *rsa.PrivateKey) (string, error) {
	claims := appendClaims(jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().UTC().Add(time.Hour).Unix(),
		"sub": "session",
		"iss": "ryudelta",
		"aud": [1]string{"ryudelta"},
		"uid": uid,
	}, customClaims)

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return t.SignedString(key)
}
