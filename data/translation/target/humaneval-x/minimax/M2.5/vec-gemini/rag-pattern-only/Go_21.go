package main

import (
	"fmt"
)

func RescaleToUnit(numbers []float64) []float64 {
	if len(numbers) == 0 {
		return numbers
	}

	minNumber := numbers[0]
	maxNumber := numbers[0]

	// Find min and max values
	for _, n := range numbers {
		if n < minNumber {
			minNumber = n
		}
		if n > maxNumber {
			maxNumber = n
		}
	}

	// Handle edge case where all numbers are the same (avoid division by zero)
	if maxNumber == minNumber {
		return make([]float64, len(numbers))
	}

	// Apply linear transformation
	result := make([]float64, len(numbers))
	for i, n := range numbers {
		result[i] = (n - minNumber) / (maxNumber - minNumber)
	}

	return result
}

func main() {
	// Example usage
	numbers := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	result := RescaleToUnit(numbers)
	fmt.Println(result)
}
