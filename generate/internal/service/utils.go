package service

import (
	"errors"
	"generate/internal/models"
)

func findCurriculum(group models.Group, curriculums []models.Curriculum) models.Curriculum {
	for _, curriculum := range curriculums {
		if group.Curriculum == curriculum.Name {
			return curriculum
		}
	}
	return models.Curriculum{}
}

func findTeacher(
	teachers *[]models.Teacher,
	plan models.Plan,
	lessonType models.LessonType,
) (models.Teacher, error) {
	// TODO: tests
	var requiredLoad int

	switch lessonType {
	case models.Lecture:
		requiredLoad = plan.Lectures * 2
	case models.Practice:
		requiredLoad = plan.Practices * 2
	case models.Laboratory:
		requiredLoad = plan.Laboratories * 2
	default:
		return models.Teacher{}, errors.New("unknown lessonType")
	}

	for _, teacher := range *teachers {
		if teacher.Load >= requiredLoad {
			canTeach := false
			for _, subject := range teacher.Subjects {
				if subject == plan.Name {
					canTeach = true
					break
				}
			}

			if canTeach {
				return teacher, nil
			}
		}
	}

	return models.Teacher{}, errors.New("teacher not found")
}
