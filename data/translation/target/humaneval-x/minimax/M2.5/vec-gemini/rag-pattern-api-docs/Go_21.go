package main

import (
	"fmt"
	"slices"
)

func RescaleToUnit(numbers []float64) []float64 {
	minNumber := slices.Min(numbers)
	maxNumber := slices.Max(numbers)

	// Handle edge case where all numbers are the same to avoid division by zero
	if maxNumber == minNumber {
		result := make([]float64, len(numbers))
		return result
	}

	result := make([]float64, len(numbers))
	for i, x := range numbers {
		result[i] = (x - minNumber) / (maxNumber - minNumber)
	}
	return result
}

func main() {
	// Test the function with the example from docstring
	input := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	result := RescaleToUnit(input)
	fmt.Println(result)
}