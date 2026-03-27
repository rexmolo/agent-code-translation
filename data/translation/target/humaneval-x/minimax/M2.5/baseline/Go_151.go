package main

import (
	"math"
)

func DoubleTheDifference(lst []float64) int {
	sum := 0
	for _, v := range lst {
		// Check if positive, an integer (no fractional part), and odd
		if v > 0 && v == math.Trunc(v) && int64(v)%2 != 0 {
			sum += int(v * v)
		}
	}
	return sum
}

func main() {
	// Test cases
	println(DoubleTheDifference([]float64{1, 3, 2, 0}))     // Expected: 10
	println(DoubleTheDifference([]float64{-1, -2, 0}))     // Expected: 0
	println(DoubleTheDifference([]float64{9, -2}))         // Expected: 81
	println(DoubleTheDifference([]float64{0}))             // Expected: 0
	println(DoubleTheDifference([]float64{}))              // Expected: 0
}
