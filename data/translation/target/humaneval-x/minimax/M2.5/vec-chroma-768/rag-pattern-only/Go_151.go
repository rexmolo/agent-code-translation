package main

import (
	"math"
)

func DoubleTheDifference(lst []float64) int {
	sum := 0
	for _, v := range lst {
		// Check if number is positive, odd, and an integer
		// v > 0 for positive
		// math.Floor(v) == v checks if it's a whole number (not a float with decimal)
		// int(v) % 2 != 0 checks if it's odd
		if v > 0 && math.Floor(v) == v && int(v)%2 != 0 {
			sum += int(v) * int(v)
		}
	}
	return sum
}

func main() {
	// Example usage
	// DoubleTheDifference([]float64{1, 3, 2, 0}) // returns 10
	// DoubleTheDifference([]float64{-1, -2, 0}) // returns 0
	// DoubleTheDifference([]float64{9, -2})     // returns 81
	// DoubleTheDifference([]float64{0})         // returns 0
}