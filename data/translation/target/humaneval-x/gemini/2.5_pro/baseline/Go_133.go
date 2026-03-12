package main

import (
	"fmt"
	"math"
)

// SumSquares calculates the sum of squared numbers after rounding each up to the nearest integer.
// It mirrors the logic of the Python function.
func SumSquares(lst []float64) int {
	var squaredSum float64 = 0
	for _, i := range lst {
		ceiledValue := math.Ceil(i)
		squaredSum += ceiledValue * ceiledValue
	}
	return int(squaredSum)
}

func main() {
	// Examples from the Python docstring
	fmt.Println(SumSquares([]float64{1, 2, 3}))
	fmt.Println(SumSquares([]float64{1, 4, 9}))
	fmt.Println(SumSquares([]float64{1, 3, 5, 7}))
	fmt.Println(SumSquares([]float64{1.4, 4.2, 0}))
	fmt.Println(SumSquares([]float64{-2.4, 1, 1}))
}
