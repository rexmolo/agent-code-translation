package main

import "math"

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