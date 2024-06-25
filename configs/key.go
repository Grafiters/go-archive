package configs

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/dgrijalva/jwt-go"
)

type JwtService struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

func LoadKeys() (*JwtService, error) {
	privateKey, err := LoadPrevKey()
	if err != nil {
		fmt.Printf("Failed to load Private Key : %s", err)
		return nil, err
	}
	publicKey, err := LoadPublicKey()
	if err != nil {
		fmt.Printf("Failed to load Private Key : %s", err)
		return nil, err
	}

	ks := &JwtService{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}

	return ks, nil

}

func LoadPublicKey() (*rsa.PublicKey, error) {
	privKeyPEM := filepath.Join("config", "secret", "public_key.pem")
	pem, err := ioutil.ReadFile(privKeyPEM)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pem)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

func LoadPrevKey() (*rsa.PrivateKey, error) {
	privKeyPEM := filepath.Join("config", "secret", "private_key.pem")
	pem, err := ioutil.ReadFile(privKeyPEM)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pem)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}
