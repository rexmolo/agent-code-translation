package main

import (
	"fmt"
	"math"
)

func MeanAbsoluteDeviation(numbers []float64) float64 {
	// Calculate the mean of the numbers
	var sum float64
	for _, n := range numbers {
		sum += n
	}
	mean := sum / float64(len(numbers))

	// Calculate the sum of absolute differences from the mean
	var absSum float64
	for _, n := range numbers {
		absSum += math.Abs(n - mean)
	}

	// Return Mean Absolute Deviation
	return absSum / float64(len(numbers))
}

func main() {
	// Test with the example from the docstring
	numbers := []float64{1.0, 2.0, 3.0, 4.0}
	result := MeanAbsoluteDeviation(numbers)
	fmt.Println(result)
}
