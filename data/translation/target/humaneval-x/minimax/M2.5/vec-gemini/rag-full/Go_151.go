package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DoubleTheDifference(lst []float64) int {
	sum := 0
	for _, i := range lst {
		// Check if positive, odd, and integer (no decimal point)
		strI := strconv.FormatFloat(i, 'f', -1, 64)
		if i > 0 && int(i)%2 != 0 && !strings.Contains(strI, ".") {
			square := i * i
			sum += int(square)
		}
	}
	return sum
}

func main() {
	// Test cases
	fmt.Println(DoubleTheDifference([]float64{1, 3, 2, 0})) // 10
	fmt.Println(DoubleTheDifference([]float64{-1, -2, 0}))   // 0
	fmt.Println(DoubleTheDifference([]float64{9, -2}))       // 81
	fmt.Println(DoubleTheDifference([]float64{0}))           // 0
	fmt.Println(DoubleTheDifference([]float64{}))            // 0
}