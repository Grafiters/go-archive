package main

import (
	"log"

	"github.com/Grafiters/archive/configs"
	"github.com/Grafiters/archive/route"

	_ "github.com/Grafiters/archive/cmd/docs"
)

// @title Fiber Swagger Example API
// @version 10.0
// @description This is a sample server for a Fiber application.
// @termsOfService http://swagger.io/terms/
// @contact.name Ryudelta Support
// @contact.url https://t.me/Grafiters
// @contact.email ryudelta7@gmail.com
// @license.name MIT
// @license.url https://github.com/Grafiters/go-archive/tree/main/license
// @host localhost:3000
// @BasePath /
func main() {
	if err := configs.Initialize(); err != nil {
		log.Fatal(err)
		return
	}

	r := route.SetupRouter()
	r.Listen(":3000")
}
