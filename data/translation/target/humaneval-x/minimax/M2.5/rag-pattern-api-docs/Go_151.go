package main

import (
	"fmt"
	"math"
)

func DoubleTheDifference(lst []float64) int {
	sum := 0
	for _, i := range lst {
		// Check if the number is positive, odd, and an integer (no fractional part)
		if i > 0 && math.Floor(i) == i && int(i)%2 != 0 {
			sum += int(i * i)
		}
	}
	return sum
}

func main() {
	// Test cases
	fmt.Println(DoubleTheDifference([]float64{1, 3, 2, 0}))   // 10
	fmt.Println(DoubleTheDifference([]float64{-1, -2, 0}))    // 0
	fmt.Println(DoubleTheDifference([]float64{9, -2}))       // 81
	fmt.Println(DoubleTheDifference([]float64{0}))           // 0
	fmt.Println(DoubleTheDifference([]float64{}))            // 0
}
