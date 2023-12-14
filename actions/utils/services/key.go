package services

import (
	"archive/actions/utils/interfaces"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
}

func LoadKeys() (*interfaces.KeyStore, error) {
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

	ks := &interfaces.KeyStore{
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
