package main

import (
	"math"
)

func DoubleTheDifference(lst []float64) int {
	sum := 0
	for _, i := range lst {
		// Check if positive and odd
		if i > 0 && math.Mod(i, 2) != 0 {
			// Check if it's an integer (no decimal part)
			if i == math.Floor(i) {
				sum += int(i * i)
			}
		}
	}
	return sum
}

func main() {
	// Test cases
	println(DoubleTheDifference([]float64{1, 3, 2, 0})) // 10
	println(DoubleTheDifference([]float64{-1, -2, 0}))  // 0
	println(DoubleTheDifference([]float64{9, -2}))      // 81
	println(DoubleTheDifference([]float64{0}))          // 0
	println(DoubleTheDifference([]float64{}))          // 0
}
