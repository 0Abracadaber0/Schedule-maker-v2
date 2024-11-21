package service

import (
	"generate/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindCurriculum(t *testing.T) {
	curriculums := []models.Curriculum{
		{
			Name: "curriculum1",
			Subjects: []models.Plan{
				{Name: "Math", Lectures: 30, Practices: 20, Laboratories: 10},
				{Name: "Science", Lectures: 25, Practices: 15, Laboratories: 5},
			},
		},
		{
			Name: "curriculum2",
			Subjects: []models.Plan{
				{Name: "History", Lectures: 20, Practices: 10, Laboratories: 0},
				{Name: "Literature", Lectures: 15, Practices: 10, Laboratories: 5},
			},
		},
	}
	group1 := models.Group{
		Name:       "group1",
		Curriculum: "curriculum1",
	}
	group2 := models.Group{
		Name:       "group2",
		Curriculum: "curriculum2",
	}
	tests := []struct {
		group       models.Group
		curriculums []models.Curriculum
		expected    models.Curriculum
	}{
		{
			group:       group1,
			curriculums: curriculums,
			expected:    curriculums[0],
		},
		{
			group:       group2,
			curriculums: curriculums,
			expected:    curriculums[1],
		},
	}
	for _, test := range tests {
		t.Run(test.group.Name, func(t *testing.T) {
			result := findCurriculum(test.group, test.curriculums)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestFindTeacher(t *testing.T) {
	teachers := []models.Teacher{
		{
			Name:     "teacher1",
			Load:     2,
			Subjects: []string{"subject1", "subject2"},
		},
		{
			Name:     "teacher2",
			Load:     4,
			Subjects: []string{"subject1", "subject2"},
		},
	}
	plan1 := models.Plan{
		Name:         "subject1",
		Lectures:     0,
		Practices:    1,
		Laboratories: 0,
		Flow:         "1",
	}
	plan2 := models.Plan{
		Name:         "subject2",
		Lectures:     1,
		Practices:    0,
		Laboratories: 1,
		Flow:         "1",
	}
	tests := []struct {
		name       string
		teachers   *[]models.Teacher
		plan       models.Plan
		lessonType models.LessonType
		expected   string
	}{
		{
			name:       "Lecture for plan2",
			teachers:   &teachers,
			plan:       plan2,
			lessonType: models.Lecture,
			expected:   "teacher1",
		},
		{
			name:       "Lab for plan2",
			teachers:   &teachers,
			plan:       plan2,
			lessonType: models.Laboratory,
			expected:   "teacher2",
		},
		{
			name:       "Prac for plan1",
			teachers:   &teachers,
			plan:       plan1,
			lessonType: models.Practice,
			expected:   "teacher2",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := findTeacher(&teachers, test.plan, test.lessonType)
			assert.Equal(t, test.expected, result)
		})
	}
}
