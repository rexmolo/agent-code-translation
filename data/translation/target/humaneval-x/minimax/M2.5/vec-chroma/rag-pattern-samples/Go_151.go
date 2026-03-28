package main

import (
	"fmt"
	"math"
)

func DoubleTheDifference(lst []float64) int {
	sum := 0
	for _, i := range lst {
		// Check if positive, odd, and an integer
		if i > 0 {
			// Check if it's an integer (no decimal part)
			if math.Floor(i) == i {
				intVal := int(i)
				if intVal%2 != 0 {
					sum += intVal * intVal
				}
			}
		}
	}
	return sum
}

func main() {
	// Test cases from docstring
	fmt.Println(DoubleTheDifference([]float64{1, 3, 2, 0}))  // 10
	fmt.Println(DoubleTheDifference([]float64{-1, -2, 0}))   // 0
	fmt.Println(DoubleTheDifference([]float64{9, -2}))      // 81
	fmt.Println(DoubleTheDifference([]float64{0}))           // 0
	fmt.Println(DoubleTheDifference([]float64{}))            // 0
}
