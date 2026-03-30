package main

import (
	"math"
)

// poly evaluates a polynomial with coefficients xs at point x.
// xs[0] + xs[1] * x + xs[2] * x^2 + ... + xs[n] * x^n
func poly(xs []int, x float64) float64 {
	result := 0.0
	for i, coeff := range xs {
		result += float64(coeff) * math.Pow(x, float64(i))
	}
	return result
}

// FindZero finds x such that poly(x) = 0 using binary search.
// It takes a list of coefficients and returns a zero point.
// The function requires an even number of coefficients with the largest
// non-zero coefficient to guarantee a solution.
func FindZero(xs []int) float64 {
	begin, end := -1.0, 1.0
	for poly(xs, begin)*poly(xs, end) > 0 {
		begin *= 2.0
		end *= 2.0
	}
	for end-begin > 1e-10 {
		center := (begin + end) / 2.0
		if poly(xs, center)*poly(xs, begin) > 0 {
			begin = center
		} else {
			end = center
		}
	}
	return begin
}
