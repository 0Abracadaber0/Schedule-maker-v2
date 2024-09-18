package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func WelcomeHandler(c *fiber.Ctx) error {
	return c.SendString("Welcome here!")
}
