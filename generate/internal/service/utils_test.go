package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEndsWithLab(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"математикаlab", true},
		{"математика", false},
		{"m", false},
		{"lab", true},
		{"", false},
		{"mat", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := endsWithLab(tt.input)
			assert.Equal(t, tt.expected, result, "Expected %v for input %s but got %v", tt.expected, tt.input, result)
		})
	}
}
