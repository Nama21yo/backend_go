package main

import (
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name     string
		grades   []float64
		expected float64
		hasError bool
	}{
		{"Normal case", []float64{80, 90, 100}, 90.0, false},
		{"Single grade", []float64{75}, 75.0, false},
		{"Empty grades", []float64{}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateAverage(tt.grades)
			if (err != nil) != tt.hasError {
				t.Errorf("Expected error: %v, got: %v", tt.hasError, err)
			}
			if !tt.hasError && got != tt.expected {
				t.Errorf("Expected average: %.2f, got: %.2f", tt.expected, got)
			}
		})
	}
}
