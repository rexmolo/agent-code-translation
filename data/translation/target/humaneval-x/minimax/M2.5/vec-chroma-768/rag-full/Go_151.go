package main

import (
	"fmt"
	"math"
)

func DoubleTheDifference(lst []float64) int {
	sum := 0
	for _, v := range lst {
		// Check if positive
		if v <= 0 {
			continue
		}
		// Check if it's an integer (no decimal part)
		// Converting float64 to int64 and back reveals if there's a fractional part
		if v != math.Floor(v) {
			continue
		}
		// Check if odd
		if int64(v)%2 == 0 {
			continue
		}
		// Square and add to sum
		sum += int(v * v)
	}
	return sum
}

func main() {
	// Test cases
	fmt.Println(DoubleTheDifference([]float64{1, 3, 2, 0}))   // 10
	fmt.Println(DoubleTheDifference([]float64{-1, -2, 0}))   // 0
	fmt.Println(DoubleTheDifference([]float64{9, -2}))       // 81
	fmt.Println(DoubleTheDifference([]float64{0}))           // 0
	fmt.Println(DoubleTheDifference([]float64{}))            // 0
}