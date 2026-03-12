package main

import (
	"math"
)

func SumSquares(lst []float64) int {
	var sum float64
	for _, v := range lst {
		c := math.Ceil(v)
		sum += c * c
	}
	return int(sum)
}
