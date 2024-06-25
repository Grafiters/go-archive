package configs

import (
	"log"

	"gorm.io/gorm"
)

var (
	DataBase  *gorm.DB
	JwtConfig *JwtService
)

func Initialize() error {
	db, err := NewDatabase()
	if err != nil {
		log.Fatal(err)
		return err
	}
	DataBase = db

	jwt, err := NewJwtConfig()
	if err != nil {
		log.Fatal(err)
		return err
	}
	JwtConfig = jwt

	return nil
}
