package main

import (
	"fmt"
	"math"
)

func DoubleTheDifference(lst []float64) int {
	sum := 0
	for _, i := range lst {
		// Check if i is a positive integer (not negative, not float)
		// Use math.Floor to check if it's a whole number
		if i > 0 && math.Floor(i) == i {
			// Convert to int to check odd/even
			intVal := int(i)
			if intVal%2 != 0 {
				sum += intVal * intVal
			}
		}
	}
	return sum
}

func main() {
	// Test cases
	fmt.Println(DoubleTheDifference([]float64{1, 3, 2, 0}))   // Expected: 10
	fmt.Println(DoubleTheDifference([]float64{-1, -2, 0}))     // Expected: 0
	fmt.Println(DoubleTheDifference([]float64{9, -2}))          // Expected: 81
	fmt.Println(DoubleTheDifference([]float64{0}))             // Expected: 0
	fmt.Println(DoubleTheDifference([]float64{}))              // Expected: 0
}