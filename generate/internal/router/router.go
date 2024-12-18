package router

import (
	handler "generate/internal/handlers"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PostGenerateHandler(ctx *fiber.Ctx) error {
	return handler.GenerateHandler(ctx)
}

func SetupRoutes(app *fiber.App, log *slog.Logger) {
	app.Use(func(ctx *fiber.Ctx) error {
		start := time.Now()

		err := ctx.Next()

		log.Info("HTTP request",
			slog.String("method", ctx.Method()),
			slog.String("route", ctx.Path()),
			slog.Int("status", ctx.Response().StatusCode()),
			slog.Duration("latency", time.Since(start)),
		)

		return err
	})

	app.Post("/generate", PostGenerateHandler)
	// TODO: tests
}
