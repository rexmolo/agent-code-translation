package main

import "math"

func MeanAbsoluteDeviation(numbers []float64) float64 {
	n := float64(len(numbers))
	if n == 0 {
		return 0
	}

	// Calculate the mean
	sum := 0.0
	for _, x := range numbers {
		sum += x
	}
	mean := sum / n

	// Calculate the sum of absolute differences from the mean
	absSum := 0.0
	for _, x := range numbers {
		absSum += math.Abs(x - mean)
	}

	return absSum / n
}