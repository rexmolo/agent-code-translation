package main

import "fmt"

func RescaleToUnit(numbers []float64) []float64 {
	// Handle empty slice
	if len(numbers) == 0 {
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

	// Handle case where all numbers are the same (avoid division by zero)
	if minNumber == maxNumber {
		return make([]float64, len(numbers))
	}

	// Apply linear transform: (x - min) / (max - min)
	result := make([]float64, len(numbers))
	for i, n := range numbers {
		result[i] = (n - minNumber) / (maxNumber - minNumber)
	}

	return result
}

func main() {
	// Test example from docstring
	result := RescaleToUnit([]float64{1.0, 2.0, 3.0, 4.0, 5.0})
	fmt.Println(result)
}