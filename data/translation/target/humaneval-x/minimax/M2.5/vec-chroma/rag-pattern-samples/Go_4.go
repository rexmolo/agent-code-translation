package main

import (
	"fmt"
	"math"
)

// MeanAbsoluteDeviation calculates the Mean Absolute Deviation around the mean
// of a given dataset. MAD = average | x - x_mean |
func MeanAbsoluteDeviation(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	// Calculate the mean
	sum := 0.0
	for _, n := range numbers {
		sum += n
	}
	mean := sum / float64(len(numbers))

	// Calculate the average of absolute differences from mean
	absSum := 0.0
	for _, n := range numbers {
		absSum += math.Abs(n - mean)
	}

	return absSum / float64(len(numbers))
}

func main() {
	// Test the function
	numbers := []float64{1.0, 2.0, 3.0, 4.0}
	result := MeanAbsoluteDeviation(numbers)
	fmt.Println(result)
}