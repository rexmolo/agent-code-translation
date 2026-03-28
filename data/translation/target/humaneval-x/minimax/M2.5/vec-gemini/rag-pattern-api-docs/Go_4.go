package main

import (
	"math"
)

func MeanAbsoluteDeviation(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	// Calculate the mean
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

	// Return the Mean Absolute Deviation
	return absSum / float64(len(numbers))
}