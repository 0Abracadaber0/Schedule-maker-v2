package main

import (
	"generate/internal/config"
	"generate/internal/router"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("Starting application",
		slog.String("env", cfg.Env),
	)

	app := fiber.New(fiber.Config{
		AppName: "Generate service",
	})
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("logger", log)
		return c.Next()
	})

	router.SetupRoutes(app, log)

	err := app.Listen(":8088")
	if err != nil {
		return
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
