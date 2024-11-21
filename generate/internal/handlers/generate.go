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

func GenerateHandler(c *fiber.Ctx) error {
	// TODO: middleware
	var data Data
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	lessons, errs := service.GenerateSchedule(data.Teachers, data.Groups, data.Curriculums)

	errsStr := make([]string, len(errs))
	for i, err := range errs {
		errsStr[i] = err.Error()
	}
	return c.JSON(fiber.Map{
		"errors":  errsStr,
		"lessons": lessons,
	})
}
