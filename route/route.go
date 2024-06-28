package route

import (
	"html/template"

	"github.com/Grafiters/archive/app/controller/account"
	"github.com/Grafiters/archive/app/controller/session"
	"github.com/Grafiters/archive/route/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func SetupRouter() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("openapi/*", swagger.New(swagger.Config{
		Title:  "Space Forum api",
		Layout: "BaseLayout",
		Plugins: []template.JS{
			template.JS(`SwaggerUIBundle.plugins.DownloadUrl`),
		},
		CustomStyle: template.CSS(`
			@import url('https://cdn.jsdelivr.net/npm/swagger-themes@1.4.3/themes/dark.css');
		`),
	}))

	api_v2_session := app.Group("/api/v2/session")
	{
		api_v2_session.Post("/", session.GoogleAuth)
	}

	api_v2_account := app.Group("/api/v2/account", middleware.Authenticate)
	{
		api_v2_account.Get("/", account.GetUsersMe)
	}

	return app
}
