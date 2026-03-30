package main

import (
	"slices"
)

func RescaleToUnit(numbers []float64) []float64 {
	minNumber := slices.Min(numbers)
	maxNumber := slices.Max(numbers)

	// Avoid division by zero when all elements are equal
	if maxNumber == minNumber {
		return make([]float64, len(numbers))
	}

	result := make([]float64, len(numbers))
	for i, n := range numbers {
		result[i] = (n - minNumber) / (maxNumber - minNumber)
	}

	return result
}