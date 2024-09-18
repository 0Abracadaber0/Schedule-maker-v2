package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func GenerateHandler(c *fiber.Ctx) error {
	log := c.Locals("logger").(*slog.Logger)

	log.Info("Request POST /generate has been get")

	return c.SendString("Welcome here!")
}
