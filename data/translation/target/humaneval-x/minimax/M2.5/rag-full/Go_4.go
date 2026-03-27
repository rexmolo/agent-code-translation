package main

import (
	"fmt"
	"math"
)

func MeanAbsoluteDeviation(numbers []float64) float64 {
	// Calculate the mean
	sum := 0.0
	for _, n := range numbers {
		sum += n
	}
	mean := sum / float64(len(numbers))

	// Calculate MAD (Mean Absolute Deviation)
	deviationSum := 0.0
	for _, n := range numbers {
		deviationSum += math.Abs(n - mean)
	}

	return deviationSum / float64(len(numbers))
}

func main() {
	// Test the function
	result := MeanAbsoluteDeviation([]float64{1.0, 2.0, 3.0, 4.0})
	fmt.Println(result) // Expected: 1.0
}
