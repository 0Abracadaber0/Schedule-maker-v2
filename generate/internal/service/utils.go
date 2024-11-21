package service

import (
	"fmt"
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
) (string, error) {
	var requiredLoad int

	switch lessonType {
	case models.Lecture:
		requiredLoad = plan.Lectures * 2
	case models.Practice:
		requiredLoad = plan.Practices * 2
	case models.Laboratory:
		requiredLoad = plan.Laboratories * 2
	default:
		fmt.Println("error")
		return "", fmt.Errorf("lesson type not supported")
	}

	for i := range *teachers {
		teacher := &(*teachers)[i]
		if teacher.Load >= requiredLoad {
			canTeach := false
			for _, subject := range teacher.Subjects {
				if subject == plan.Name {
					canTeach = true
					break
				}
			}

			if canTeach {
				teacher.Load -= requiredLoad
				return teacher.Name, nil
			}
		}
	}

	fmt.Println("error")
	return "", fmt.Errorf("teacher for %s %s not found", lessonType, plan.Name)
}
