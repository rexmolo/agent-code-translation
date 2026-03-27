package main

import "math"

// MeanAbsoluteDeviation calculates the Mean Absolute Deviation around the mean
// of the given dataset.
// MAD = average | x - x_mean |
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

	// Calculate the average absolute deviation from the mean
	var madSum float64
	for _, n := range numbers {
		madSum += math.Abs(n - mean)
	}

	return madSum / float64(len(numbers))
}

func main() {
	// Example usage
	result := MeanAbsoluteDeviation([]float64{1.0, 2.0, 3.0, 4.0})
	println(result) // Output: 1
}