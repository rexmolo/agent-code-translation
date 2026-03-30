package main

import (
	"math"
)

func SumSquares(lst []float64) int {
	squared := 0
	for _, v := range lst {
		ceiled := int(math.Ceil(v))
		squared += ceiled * ceiled
	}
	return squared
}
