package service

import (
	"generate/internal/models"
)

func GenerateSchedule(
	teachers []models.Teacher,
	groups []models.Group,
	curriculums []models.Curriculum,
	// classrooms []models.Classroom,
) ([]models.Lesson, []error) {
	var errs []error
	var lessons []models.Lesson

	for _, group := range groups {

		curriculum := findCurriculum(group, curriculums)
		for _, plan := range curriculum.Subjects {
			// TODO: учитывать flow для лекций

			// поиск преподавателя для лекций
			if plan.Lectures > 0 {
				teacher, err := findTeacher(&teachers, plan, models.Lecture)
				if err != nil {
					errs = append(errs, err)
				}
				// добавление лекций к списку пар
				for i := 0; i < plan.Lectures; i++ {
					lessons = append(lessons, models.Lesson{
						Subject: plan.Name,
						Group:   group.Name,
						Teacher: teacher,
						IsLab:   false,
					})
				}
			}

			if plan.Practices > 0 {
				// поиск преподавателя для практик
				teacher, err := findTeacher(&teachers, plan, models.Practice)
				if err != nil {
					errs = append(errs, err)
				}
				// добавление практик к списку пар
				for i := 0; i < plan.Practices; i++ {
					lessons = append(lessons, models.Lesson{
						Subject: plan.Name,
						Group:   group.Name,
						Teacher: teacher,
						IsLab:   false,
					})
				}
			}

			if plan.Laboratories > 0 {
				// поиск преподавателя для лабораторных
				teacher, err := findTeacher(&teachers, plan, models.Laboratory)
				if err != nil {
					errs = append(errs, err)
				}
				// добавление лабораторных к списку пар
				for i := 0; i < plan.Laboratories; i++ {
					lessons = append(lessons, models.Lesson{
						Subject: plan.Name,
						Group:   group.Name,
						Teacher: teacher,
						IsLab:   true,
					})
				}
			}

		}
	}

	return lessons, errs
}
