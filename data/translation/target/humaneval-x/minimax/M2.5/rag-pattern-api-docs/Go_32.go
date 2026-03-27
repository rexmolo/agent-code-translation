package main

import (
	"fmt"
	"math"
)

// Poly evaluates polynomial with coefficients xs at point x.
// xs[0] + xs[1] * x + xs[2] * x^2 + .... xs[n] * x^n
func Poly(xs []float64, x float64) float64 {
	result := 0.0
	for i, coeff := range xs {
		result += coeff * math.Pow(x, float64(i))
	}
	return result
}

// FindZero finds x such that poly(x) = 0.
// FindZero returns only one zero point, even if there are many.
// Moreover, FindZero only takes list xs having even number of coefficients
// and largest non zero coefficient as it guarantees a solution.
func FindZero(xs []int) float64 {
	// Convert int slice to float64 slice for mathematical operations
	floatXs := make([]float64, len(xs))
	for i, v := range xs {
		floatXs[i] = float64(v)
	}

	begin, end := -1.0, 1.0
	for Poly(floatXs, begin)*Poly(floatXs, end) > 0 {
		begin *= 2.0
		end *= 2.0
	}
	for end-begin > 1e-10 {
		center := (begin + end) / 2.0
		if Poly(floatXs, center)*Poly(floatXs, begin) > 0 {
			begin = center
		} else {
			end = center
		}
	}
	return begin
}

func main() {
	// Test cases
	fmt.Printf("%.2f\n", FindZero([]int{1, 2}))         // f(x) = 1 + 2x
	fmt.Printf("%.2f\n", FindZero([]int{-6, 11, -6, 1})) // (x - 1) * (x - 2) * (x - 3)
}