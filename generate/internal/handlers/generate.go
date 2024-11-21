package handlers

import (
	"generate/internal/models"
	"generate/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Data struct {
	Teachers    []models.Teacher    `json:"teachers"`
	Groups      []models.Group      `json:"groups"`
	Curriculums []models.Curriculum `json:"curriculums"`
	Classrooms  []models.Classroom  `json:"classrooms"`
}

// TODO: presenter.go (https://github.com/gofiber/recipes/blob/master/clean-architecture/api/presenter/book.go)

func GenerateHandler(c *fiber.Ctx) error {
	// TODO: take logger
	var data Data
	if err := c.BodyParser(data); err != nil {
		return err
	}

	service.GenerateSchedule(data.Teachers, data.Groups, data.Curriculums)

	return c.SendStatus(fiber.StatusNoContent)
}
