package main

import "math"

func Multiply(a, b int) int {
	x := int(math.Abs(float64(a % 10)))
	y := int(math.Abs(float64(b % 10)))
	return x * y
}
