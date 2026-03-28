package main

import (
	"fmt"
	"math"
)

func MeanAbsoluteDeviation(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	// Calculate mean
	sum := 0.0
	for _, n := range numbers {
		sum += n
	}
	mean := sum / float64(len(numbers))

	// Calculate MAD
	madSum := 0.0
	for _, n := range numbers {
		madSum += math.Abs(n - mean)
	}

	return madSum / float64(len(numbers))
}

func main() {
	// Test case from docstring
	result := MeanAbsoluteDeviation([]float64{1.0, 2.0, 3.0, 4.0})
	fmt.Println(result)
}
