package main

import (
	"math"
)

// Poly evaluates a polynomial with coefficients xs at point x
func Poly(xs []int, x float64) float64 {
	result := 0.0
	for i, coeff := range xs {
		result += float64(coeff) * math.Pow(x, float64(i))
	}
	return result
}

// FindZero finds x such that poly(x) = 0 using binary search
func FindZero(xs []int) float64 {
	begin := -1.0
	end := 1.0

	// Expand the search range until we find values with opposite signs
	for Poly(xs, begin)*Poly(xs, end) > 0 {
		begin *= 2.0
		end *= 2.0
	}

	// Binary search to find the root
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
