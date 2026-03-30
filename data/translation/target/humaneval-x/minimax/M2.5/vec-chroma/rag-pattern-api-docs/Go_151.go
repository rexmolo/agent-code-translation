package main

import (
	"fmt"
	"math"
)

func DoubleTheDifference(lst []float64) int {
	sum := 0
	for _, v := range lst {
		// Check if positive, odd, and an integer (no fractional part)
		if v > 0 && math.Mod(v, 1) == 0 && int64(v)%2 != 0 {
			sum += int(v * v)
		}
	}
	return sum
}

func main() {
	// Test examples
	fmt.Println(DoubleTheDifference([]float64{1, 3, 2, 0})) // 10
	fmt.Println(DoubleTheDifference([]float64{-1, -2, 0})) // 0
	fmt.Println(DoubleTheDifference([]float64{9, -2}))     // 81
	fmt.Println(DoubleTheDifference([]float64{0}))         // 0
	fmt.Println(DoubleTheDifference([]float64{}))         // 0
}
