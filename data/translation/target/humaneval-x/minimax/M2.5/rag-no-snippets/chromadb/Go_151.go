package main

import (
	"math"
)

func DoubleTheDifference(lst []float64) int {
	sum := 0
	for _, f := range lst {
		// Check if positive, is an integer (no decimal part), and is odd
		if f > 0 && f == math.Floor(f) {
			i := int(f)
			if i%2 != 0 {
				sum += i * i
			}
		}
	}
	return sum
}

func main() {
	// Test cases
	println(DoubleTheDifference([]float64{1, 3, 2, 0}))    // 10
	println(DoubleTheDifference([]float64{-1, -2, 0}))     // 0
	println(DoubleTheDifference([]float64{9, -2}))        // 81
	println(DoubleTheDifference([]float64{0}))            // 0
	println(DoubleTheDifference([]float64{}))            // 0
}