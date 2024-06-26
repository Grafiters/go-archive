package route

import (
	sessionController "github.com/Grafiters/archive/app/controller/session"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouter() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	api_v2_session := app.Group("/api/v2/session")
	{
		api_v2_session.Post("/", sessionController.GoogleAuth)
	}

	return app
}
