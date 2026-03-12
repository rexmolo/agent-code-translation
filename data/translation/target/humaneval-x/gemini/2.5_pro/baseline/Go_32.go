package main

import (
	"fmt"
	"math"
)

// poly evaluates a polynomial with coefficients xs at point x.
// The formula is: xs[0] + xs[1]*x + xs[2]*x^2 + ...
func poly(xs []int, x float64) float64 {
	sum := 0.0
	for i, coeff := range xs {
		sum += float64(coeff) * math.Pow(x, float64(i))
	}
	return sum
}

// FindZero finds a root for the polynomial defined by the coefficients xs.
// It uses the bisection method to find x such that poly(x) = 0.
// FindZero returns only one zero point, even if there are many.
// The method assumes that the coefficients guarantee a solution can be found
// by expanding the initial interval [-1, 1].
func FindZero(xs []int) float64 {
	begin, end := -1.0, 1.0
	// Expand the interval [begin, end] until a root is bracketed.
	// This happens when poly(begin) and poly(end) have opposite signs.
	for poly(xs, begin)*poly(xs, end) > 0 {
		begin *= 2.0
		end *= 2.0
	}

	// Bisection method to find the root within the interval.
	for end-begin > 1e-10 {
		center := (begin + end) / 2.0
		if poly(xs, center)*poly(xs, begin) > 0 {
			// Root is in the [center, end] half.
			begin = center
		} else {
			// Root is in the [begin, center] half.
			end = center
		}
	}
	return begin
}

func main() {
	// Example 1 from Python docstring: f(x) = 1 + 2x
	result1 := FindZero([]int{1, 2})
	fmt.Printf("find_zero([1, 2]) -> %.2f\n", result1)

	// Example 2 from Python docstring: (x - 1)*(x - 2)*(x - 3) = -6 + 11x - 6x^2 + x^3
	result2 := FindZero([]int{-6, 11, -6, 1})
	fmt.Printf("find_zero([-6, 11, -6, 1]) -> %.2f\n", result2)
}
