package service

import (
	"generate/internal/models"
	"math/rand"
	"strings"
)

func isEndLab(s string) bool {
	return strings.HasPrefix(s, "lab")
}

func GenerateAllLessons(data models.ScheduleGenerator) []map[string][]interface{} {
	allLessons := []map[string][]interface{}{}

	// Генерация лекций
	for planName, lessons := range data.Plans {
		for lessonName, lesson := range lessons {
			teacher := data.Subjects[lessonName][rand.Intn(len(data.Subjects[lessonName]))]
			for i := 0; i < lesson.Lectures; i++ {
				streamCount := 0
				for _, groupPlan := range data.Groups {
					if groupPlan == planName {
						streamCount++
					}
				}
				allLessons = append(allLessons, map[string][]interface{}{
					lessonName: {lesson.Stream, teacher, streamCount},
				})
			}
		}
	}

	// Генерация практик и лабораторных
	for group, plan := range data.Groups {
		lessons := data.Plans[plan]
		for lessonName, lesson := range lessons {
			teacher := data.Subjects[lessonName][rand.Intn(len(data.Subjects[lessonName]))]
			for i := 0; i < lesson.Practices; i++ {
				allLessons = append(allLessons, map[string][]interface{}{
					lessonName: {group, teacher},
				})
			}
			for i := 0; i < lesson.Labs; i++ {
				allLessons = append(allLessons, map[string][]interface{}{
					lessonName: {group + "lab", teacher},
				})
			}
		}
	}

	return allLessons
}
