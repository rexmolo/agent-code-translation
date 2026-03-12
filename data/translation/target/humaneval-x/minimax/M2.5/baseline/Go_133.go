package main

import (
	"math"
)

func SumSquares(lst []float64) int {
	squared := 0
	for _, i := range lst {
		ceiled := math.Ceil(i)
		squared += int(ceiled) * int(ceiled)
	}
	return squared
}