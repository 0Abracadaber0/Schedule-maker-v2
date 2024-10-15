package service_test

import (
	"errors"
	model "generate/internal/models"
	"generate/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTeacher(t *testing.T) {
	teachers := []model.Teacher{
		{
			Name:            "Alice",
			Subjects:        []string{"math", "physics"},
			CanGiveLectures: true,
			Load:            4,
		},
		{
			Name:            "Bob",
			Subjects:        []string{"chemistry"},
			CanGiveLectures: false,
			Load:            2,
		},
		{
			Name:            "Charlie",
			Subjects:        []string{"biology", "math"},
			CanGiveLectures: true,
			Load:            2,
		},
		{
			Name:            "Max",
			Subjects:        []string{"chess"},
			CanGiveLectures: true,
			Load:            0,
		},
	}

	tests := []struct {
		name        string
		subject     string
		lessonType  model.LessonType
		expectedErr error
		expected    *model.Teacher
	}{
		{
			name:        "Find teacher for lecture",
			subject:     "math",
			lessonType:  model.Lecture,
			expectedErr: nil,
			expected:    &teachers[0],
		},
		{
			name:        "Find teacher for practice",
			subject:     "chemistry",
			lessonType:  model.Practice,
			expectedErr: nil,
			expected:    &teachers[1],
		},
		{
			name:        "Find teacher for laboratory",
			subject:     "biology",
			lessonType:  model.Laboratory,
			expectedErr: nil,
			expected:    &teachers[2],
		},
		{
			name:        "Teacher cannot give lectures",
			subject:     "chemistry",
			lessonType:  model.Lecture,
			expectedErr: errors.New("no teacher available for subject: chemistry, type: Lecture"),
			expected:    nil,
		},
		{
			name:        "No teacher found for the subject",
			subject:     "computer science",
			lessonType:  model.Laboratory,
			expectedErr: errors.New("no teacher available for subject: computer science, type: Laboratory"),
			expected:    nil,
		},
		{
			name:        "The teacher ran out of hours",
			subject:     "chess",
			lessonType:  model.Practice,
			expectedErr: errors.New("no teacher available for subject: chess, type: Practice"),
			expected:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := service.FindTeacher(test.subject, test.lessonType, teachers)

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		item     string
		expected bool
	}{
		{
			name:     "Item present in the slice",
			slice:    []string{"math", "physics", "biology"},
			item:     "math",
			expected: true,
		},
		{
			name:     "Item not present in the slice",
			slice:    []string{"math", "physics", "biology"},
			item:     "chemistry",
			expected: false,
		},
		{
			name:     "Empty slice",
			slice:    []string{},
			item:     "math",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := service.Contains(test.slice, test.item)
			assert.Equal(t, test.expected, result)
		})
	}
}
