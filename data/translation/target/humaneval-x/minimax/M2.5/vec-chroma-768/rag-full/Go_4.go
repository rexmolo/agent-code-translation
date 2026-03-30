package main

import (
	"fmt"
	"math"
)

func MeanAbsoluteDeviation(numbers []float64) float64 {
	// Calculate the mean of the numbers
	sum := 0.0
	for _, n := range numbers {
		sum += n
	}
	mean := sum / float64(len(numbers))

	// Calculate the sum of absolute differences from the mean
	absSum := 0.0
	for _, n := range numbers {
		absSum += math.Abs(n - mean)
	}

	// Return the Mean Absolute Deviation
	return absSum / float64(len(numbers))
}

func main() {
	// Test with example from docstring: [1.0, 2.0, 3.0, 4.0] should return 1.0
	numbers := []float64{1.0, 2.0, 3.0, 4.0}
	result := MeanAbsoluteDeviation(numbers)
	fmt.Println(result)
}
