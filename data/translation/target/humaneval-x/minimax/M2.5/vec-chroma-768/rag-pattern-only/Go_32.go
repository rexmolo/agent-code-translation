package main

import (
	"fmt"
	"math"
)

func poly(xs []int, x float64) float64 {
	/*
	Evaluates polynomial with coefficients xs at point x.
	return xs[0] + xs[1] * x + xs[1] * x^2 + .... xs[n] * x^n
	*/
	result := 0.0
	for i, coeff := range xs {
		result += float64(coeff) * math.Pow(x, float64(i))
	}
	return result
}

func FindZero(xs []int) float64 {
	/*
	xs are coefficients of a polynomial.
	findZero find x such that poly(x) = 0.
	findZero returns only only zero point, even if there are many.
	Moreover, findZero only takes list xs having even number of coefficients
	and largest non zero coefficient as it guarantees
	a solution.
	*/
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

func main() {
	// Test examples from the original Python docstrings
	result1 := FindZero([]int{1, 2}) // f(x) = 1 + 2x
	fmt.Printf("find_zero([1, 2]) = %.2f (expected -0.5)\n", result1)

	result2 := FindZero([]int{-6, 11, -6, 1}) // (x - 1) * (x - 2) * (x - 3) = -6 + 11x - 6x^2 + x^3
	fmt.Printf("find_zero([-6, 11, -6, 1]) = %.2f (expected 1.0)\n", result2)
}