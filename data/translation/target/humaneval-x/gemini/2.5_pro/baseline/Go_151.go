package main

import (
	"fmt"
	"math"
)

func DoubleTheDifference(lst []float64) int {
	sum := 0
	for _, num := range lst {
		if num <= 0 {
			continue
		}

		if num != math.Floor(num) {
			continue
		}

		val := int(num)

		if val%2 != 0 {
			sum += val * val
		}
	}
	return sum
}

func main() {
	fmt.Println(DoubleTheDifference([]float64{1, 3, 2, 0}))
	fmt.Println(DoubleTheDifference([]float64{-1, -2, 0}))
	fmt.Println(DoubleTheDifference([]float64{9, -2}))
	fmt.Println(DoubleTheDifference([]float64{0}))
	fmt.Println(DoubleTheDifference([]float64{}))
}
