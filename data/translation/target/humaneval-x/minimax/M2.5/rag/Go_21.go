package main

import (
	"fmt"
	"slices"
)

func RescaleToUnit(numbers []float64) []float64 {
	if len(numbers) < 2 {
		return numbers
	}

	minNumber := slices.Min(numbers)
	maxNumber := slices.Max(numbers)

	// Avoid division by zero if all numbers are the same
	if maxNumber == minNumber {
		return make([]float64, len(numbers))
	}

	result := make([]float64, len(numbers))
	for i, x := range numbers {
		result[i] = (x - minNumber) / (maxNumber - minNumber)
	}

	return result
}

func main() {
	// Test
	numbers := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(RescaleToUnit(numbers))
}
