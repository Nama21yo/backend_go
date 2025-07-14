package main

import "errors"

// calculateAverage returns the average from a slice of float64 grades.
func calculateAverage(grades []float64) (float64, error) {
	if len(grades) == 0 {
		return 0, errors.New("no grades provided")
	}
	var sum float64
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades)), nil
}
