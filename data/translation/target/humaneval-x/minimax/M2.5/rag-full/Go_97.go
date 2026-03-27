package main

import (
	"math"
)

func Multiply(a, b int) int {
	return int(math.Abs(float64(a%10))) * int(math.Abs(float64(b%10)))
}
