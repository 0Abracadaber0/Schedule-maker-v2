package router

import (
	handler "generate/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.WelcomeHandler)
}
