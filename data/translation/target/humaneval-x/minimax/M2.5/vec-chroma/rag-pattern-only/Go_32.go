package main

import (
	"math"
)

func Poly(xs []int, x float64) float64 {
	result := 0.0
	for i, coeff := range xs {
		result += float64(coeff) * math.Pow(x, float64(i))
	}
	return result
}

func FindZero(xs []int) float64 {
	begin, end := -1.0, 1.0
	for Poly(xs, begin)*Poly(xs, end) > 0 {
		begin *= 2.0
		end *= 2.0
	}
	for end-begin > 1e-10 {
		center := (begin + end) / 2.0
		if Poly(xs, center)*Poly(xs, begin) > 0 {
			begin = center
		} else {
			end = center
		}
	}
	return begin
}