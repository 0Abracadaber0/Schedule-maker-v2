package service

import (
	"generate/internal/models"
	"math/rand"
	"strings"
)

// Проверяет, является ли строка лабораторной работой
func isEndLab(s string) bool {
	return strings.HasSuffix(s, "lab")
}

// GenerateAllLessons генерирует уникальные уроки
func GenerateAllLessons(data models.ScheduleGenerator) []map[string][]interface{} {
	allLessons := []map[string][]interface{}{}

	// Генерация лекций
	for planName, lessons := range data.Plans {
		for lessonName, lesson := range lessons {
			teacher := data.Subjects[lessonName][rand.Intn(len(data.Subjects[lessonName]))]
			found := false

			// Проверка на уникальность лекции
			for amount := 0; amount < 6; amount++ {
				lessonEntry := map[string][]interface{}{
					lessonName: {lesson.Stream, teacher, amount},
				}
				for _, existing := range allLessons {
					if equalLessonEntries(existing, lessonEntry) {
						found = true
						break
					}
				}
				if found {
					break
				}
			}

			if !found {
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
	}

	// Генерация практик и лабораторных
	for group, plan := range data.Groups {
		lessons := data.Plans[plan]
		for lessonName, lesson := range lessons {
			teacher := data.Subjects[lessonName][rand.Intn(len(data.Subjects[lessonName]))]

			// Добавление практических занятий
			for i := 0; i < lesson.Practices; i++ {
				allLessons = append(allLessons, map[string][]interface{}{
					lessonName: {group, teacher},
				})
			}

			// Добавление лабораторных занятий
			for i := 0; i < lesson.Labs; i++ {
				allLessons = append(allLessons, map[string][]interface{}{
					lessonName: {group + "lab", teacher},
				})
			}
		}
	}

	return allLessons
}

// Сравнивает два урока на идентичность
func equalLessonEntries(a, b map[string][]interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valA := range a {
		valB, ok := b[key]
		if !ok || len(valA) != len(valB) {
			return false
		}
		for i := range valA {
			if valA[i] != valB[i] {
				return false
			}
		}
	}
	return true
}
