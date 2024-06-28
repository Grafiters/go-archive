package configs

import (
	"log"

	"gorm.io/gorm"
)

var (
	DataBase  *gorm.DB
	JwtConfig *JwtService
	Logger    *LoggerFormat
)

func Initialize() error {
	Logger = NewLogger()

	// db, err := NewDatabase()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return err
	// }
	DataBase = nil

	jwt, err := NewJwtConfig()
	if err != nil {
		log.Fatal(err)
		return err
	}
	JwtConfig = jwt

	return nil
}
