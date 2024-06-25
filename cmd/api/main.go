package main

import (
	"log"

	"github.com/Grafiters/archive/configs"
	"github.com/Grafiters/archive/route"
)

func main() {
	if err := configs.Initialize(); err != nil {
		log.Fatal(err)
		return
	}

	r := route.SetupRouter()
	r.Listen(":3000")
}
