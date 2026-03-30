package main

import (
	"slices"
)

func RescaleToUnit(numbers []float64) []float64 {
	minNumber := slices.Min(numbers)
	maxNumber := slices.Max(numbers)

	result := make([]float64, len(numbers))
	for i, x := range numbers {
		result[i] = (x - minNumber) / (maxNumber - minNumber)
	}
	return result
}
