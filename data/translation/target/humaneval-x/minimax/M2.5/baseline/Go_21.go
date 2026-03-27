package main

import (
	"fmt"
	"math"
)

func RescaleToUnit(numbers []float64) []float64 {
	if len(numbers) < 2 {
		return numbers
	}

	// Find min and max values
	minNumber := numbers[0]
	maxNumber := numbers[0]

	for _, n := range numbers {
		if n < minNumber {
			minNumber = n
		}
		if n > maxNumber {
			maxNumber = n
		}
	}

	// Handle edge case: if all numbers are the same
	if maxNumber == minNumber {
		result := make([]float64, len(numbers))
		for i := range result {
			result[i] = 0.0
		}
		return result
	}

	// Apply linear transformation: (x - min) / (max - min)
	result := make([]float64, len(numbers))
	for i, x := range numbers {
		result[i] = (x - minNumber) / (maxNumber - minNumber)
	}

	return result
}

func main() {
	// Example usage
	numbers := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	result := RescaleToUnit(numbers)
	fmt.Printf("%v\n", result)
}
