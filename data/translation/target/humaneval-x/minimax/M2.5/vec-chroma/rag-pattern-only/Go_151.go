package main

import (
	"math"
)

func DoubleTheDifference(lst []float64) int {
	sum := 0
	for _, v := range lst {
		// Check if positive, odd, and an integer (no decimal part)
		if v > 0 && int64(v)%2 != 0 && v == math.Floor(v) {
			sum += int(v * v)
		}
	}
	return sum
}
