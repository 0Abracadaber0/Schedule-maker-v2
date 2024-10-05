package handlers

import (
	model "generate/internal/models"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type RequestData struct {
	Curriculums []model.Curriculum
	Teachers    []model.Teacher
}

func GenerateHandler(c *fiber.Ctx) error {
	log := c.Locals("logger").(*slog.Logger)
	log.Info("Request POST /generate has been get")

	data := new(RequestData)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// data.Curriculums и data.Teachers

	return c.JSON(data)
}
