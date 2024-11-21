package main

import (
	"context"
	"generate/internal/config"
	"generate/internal/router"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	go func() {
		err := app.Listen(":8088")
		if err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	close(quit)
	log.Info("Shutting down...")
	err := app.Shutdown()
	if err != nil {
		panic(err)
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
