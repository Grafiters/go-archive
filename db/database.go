package database

import (
	"archive/db/config"
	"log"
)

func Create() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func Drop() {
	db, err := config.ConfigDB()
	if err != nil {
		log.Fatal(err)
	}

	error := config.DeleteDatabase(db)
	if err != nil {
		log.Fatal(error)
	}

	defer db.Close()
}
