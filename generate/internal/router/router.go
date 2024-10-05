package router

import (
	handler "generate/internal/handlers"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, log *slog.Logger) {
	app.Post("/generate", handler.GenerateHandler)
	app.Get("/metrics", func(c *fiber.Ctx) error {
		return c.Redirect("http://localhost:3000", fiber.StatusFound)
	})

}
