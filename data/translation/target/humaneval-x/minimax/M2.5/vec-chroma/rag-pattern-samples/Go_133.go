package main

import (
	"fmt"
	"math"
)

func SumSquares(lst []float64) int {
	squared := 0
	for _, i := range lst {
		ceiled := int(math.Ceil(i))
		squared += ceiled * ceiled
	}
	return squared
}

func main() {
	// Test cases from examples
	fmt.Println(SumSquares([]float64{1, 2, 3}))          // 14
	fmt.Println(SumSquares([]float64{1, 4, 9}))          // 98
	fmt.Println(SumSquares([]float64{1, 3, 5, 7}))       // 84
	fmt.Println(SumSquares([]float64{1.4, 4.2, 0}))      // 29
	fmt.Println(SumSquares([]float64{-2.4, 1, 1}))       // 6
}