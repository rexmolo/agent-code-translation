package mad

import "math"

// MeanAbsoluteDeviation calculates the Mean Absolute Deviation around the mean
// of a given dataset. MAD = average | x - x_mean |
func MeanAbsoluteDeviation(numbers []float64) float64 {
	n := float64(len(numbers))
	mean := 0.0

	// Calculate mean
	for _, x := range numbers {
		mean += x
	}
	mean /= n

	// Calculate MAD
	totalDiff := 0.0
	for _, x := range numbers {
		totalDiff += math.Abs(x - mean)
	}
	return totalDiff / n
}
