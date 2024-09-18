package router

import (
	handler "generate/internal/handlers"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, log *slog.Logger) {
	app.Post("/generate", handler.GenerateHandler)
}
