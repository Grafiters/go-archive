package services

import (
	"archive/actions/utils/interfaces"
	"archive/models"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	key *interfaces.KeyStore
	err error
)

func init() {
	key, err = LoadKeys()
	if err != nil {
		log.Fatalf("Failed to load Key Store")
	}
}

func EncodeToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().UTC().Add(time.Hour).Unix(),
		"sub": "session",
		"iss": "ryudelta",
		"aud": [1]string{"ryudelta"},
		"uid": user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(key.PrivateKey)
	if err != nil {
		log.Fatal("Error Signed : %s", err)
		return "", err
	}

	return signedToken, nil
}

func DecodeToken(tokenString string, target interface{}) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key.PublicKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		recordBytes := claims["record"].(map[string]interface{})
		recordJSON, err := json.Marshal(recordBytes)
		if err != nil {
			return err
		}
		err = json.Unmarshal(recordJSON, &target)
		if err != nil {
			return err
		}

		return nil
	} else {
		return err
	}
}
