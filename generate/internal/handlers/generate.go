package handlers

import (
	model "generate/internal/models"
	"generate/internal/service"

	"github.com/gofiber/fiber/v2"
)

type RequestData struct {
	Curriculums []model.Curriculum
	Teachers    []model.Teacher
	Classrooms  []model.Classroom
}

func GenerateHandler(c *fiber.Ctx) error {

	data := new(RequestData)
	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	service.GenerateSchedule(data.Curriculums, data.Teachers, data.Classrooms)

	return c.JSON(data)
}
