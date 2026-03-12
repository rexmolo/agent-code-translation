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
	sum := 0.0
	for i, coeff := range xs {
		sum += float64(coeff) * math.Pow(x, float64(i))
	}
	return sum
}

func FindZero(xs []int) float64 {
	/* xs are coefficients of a polynomial.
	find_zero find x such that poly(x) = 0.
	find_zero returns only only zero point, even if there are many.
	Moreover, find_zero only takes list xs having even number of coefficients
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
	// Test cases
	result1 := FindZero([]int{1, 2})
	fmt.Printf("%.2f\n", result1) // -0.5

	result2 := FindZero([]int{-6, 11, -6, 1})
	fmt.Printf("%.2f\n", result2) // 1.0
}
