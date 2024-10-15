package service

import (
	"fmt"
	model "generate/internal/models"
)

// func GenerateSchedule(
// 	curriculums []model.Curriculum,
// 	teachers []model.Teacher,
// 	classrooms []model.Classroom,
// 	log *slog.Logger,
// ) ([]model.Lesson, error) {

// }

func FindTeacher(
	subjectName string,
	lessonType model.LessonType,
	teachers []model.Teacher,
) (*model.Teacher, error) {
	for _, teacher := range teachers {

		// If the teacher cannot teach this subject
		if !Contains(teacher.Subjects, subjectName) {
			continue
		}

		// If the teacher cannot give lectures and this pair is a lecture
		if lessonType == model.Lecture && !teacher.CanGiveLectures {
			continue
		}

		// If the teacher is already busy
		if teacher.Load == 0 {
			continue
		}

		// If it fits all the parameters
		return &teacher, nil
	}

	return nil, fmt.Errorf("no teacher available for subject: %s, type: %s", subjectName, lessonType)
}

func Contains(slice []string, item string) bool {
	for _, elem := range slice {
		if elem == item {
			return true
		}
	}
	return false
}
