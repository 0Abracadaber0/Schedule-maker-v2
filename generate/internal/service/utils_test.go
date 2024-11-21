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
