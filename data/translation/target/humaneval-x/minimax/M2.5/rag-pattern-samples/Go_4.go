package main

import (
	"fmt"
	"math"
)

func MeanAbsoluteDeviation(numbers []float64) float64 {
	// Calculate the mean
	mean := 0.0
	for _, v := range numbers {
		mean += v
	}
	mean /= float64(len(numbers))

	// Calculate the sum of absolute differences from mean
	sumAbsDiff := 0.0
	for _, x := range numbers {
		sumAbsDiff += math.Abs(x - mean)
	}

	// Return the Mean Absolute Deviation
	return sumAbsDiff / float64(len(numbers))
}

func main() {
	// Test the function
	result := MeanAbsoluteDeviation([]float64{1.0, 2.0, 3.0, 4.0})
	fmt.Printf("%.1f\n", result)
}