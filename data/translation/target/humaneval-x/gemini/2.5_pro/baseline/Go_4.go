package main

import (
	"fmt"
	"math"
)

// MeanAbsoluteDeviation calculates the Mean Absolute Deviation for a slice of float64s.
// It is the average absolute difference between each element and the mean.
// MAD = average | x - x_mean |
func MeanAbsoluteDeviation(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}

	n := float64(len(numbers))

	// Calculate the mean
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	mean := sum / n

	// Calculate the sum of absolute deviations from the mean
	var deviationSum float64
	for _, num := range numbers {
		deviationSum += math.Abs(num - mean)
	}

	// Return the mean absolute deviation
	return deviationSum / n
}

// main function to demonstrate the usage of MeanAbsoluteDeviation
func main() {
	// Example from the Python doctest
	data := []float64{1.0, 2.0, 3.0, 4.0}
	result := MeanAbsoluteDeviation(data)
	fmt.Println(result) // Expected output: 1.0

	// Test with an empty slice
	emptyData := []float64{}
	resultEmpty := MeanAbsoluteDeviation(emptyData)
	fmt.Println(resultEmpty) // Expected output: 0.0
}