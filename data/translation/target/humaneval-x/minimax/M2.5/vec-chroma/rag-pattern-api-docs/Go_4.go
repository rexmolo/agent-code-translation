package main

import (
	"fmt"
	"math"
)

// MeanAbsoluteDeviation calculates the Mean Absolute Deviation around the mean
// for a given slice of numbers.
// MAD = average | x - x_mean |
func MeanAbsoluteDeviation(numbers []float64) float64 {
	// Calculate the mean
	var sum float64
	for _, n := range numbers {
		sum += n
	}
	mean := sum / float64(len(numbers))

	// Calculate MAD: average of absolute differences from mean
	var madSum float64
	for _, n := range numbers {
		madSum += math.Abs(n - mean)
	}

	return madSum / float64(len(numbers))
}

func main() {
	// Test with the example from the docstring
	numbers := []float64{1.0, 2.0, 3.0, 4.0}
	result := MeanAbsoluteDeviation(numbers)
	fmt.Println(result)
}
