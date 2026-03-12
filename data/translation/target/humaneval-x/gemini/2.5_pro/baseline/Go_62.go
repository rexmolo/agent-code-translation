package main

import "fmt"

// Derivative computes the derivative of a polynomial represented by its coefficients.
// xs[0] + xs[1] * x + xs[2] * x^2 + ...
// It returns the coefficients of the derivative polynomial.
func Derivative(xs []int) []int {
	// If the polynomial is constant (degree 0) or empty,
	// the derivative is 0, which is represented by an empty slice of coefficients
	// for higher-order terms.
	if len(xs) <= 1 {
		return []int{}
	}

	// The derivative polynomial has a degree one less than the original,
	// so the resulting slice of coefficients is one element shorter.
	result := make([]int, len(xs)-1)

	// The loop starts from i=1 because the derivative of the constant term (xs[0]) is zero.
	// The new coefficient for x^k is derived from the old coefficient for x^(k+1).
	// The power rule (d/dx of c*x^n = n*c*x^(n-1)) means the new coefficient
	// at index i-1 is the old coefficient at index i multiplied by its power, i.
	for i := 1; i < len(xs); i++ {
		result[i-1] = xs[i] * i
	}

	return result
}

func main() {
	// Example 1 from the docstring
	poly1 := []int{3, 1, 2, 4, 5}
	deriv1 := Derivative(poly1)
	fmt.Printf("Derivative of %v is %v\n", poly1, deriv1) // Expected: [1, 4, 12, 20]

	// Example 2 from the docstring
	poly2 := []int{1, 2, 3}
	deriv2 := Derivative(poly2)
	fmt.Printf("Derivative of %v is %v\n", poly2, deriv2) // Expected: [2, 6]

	// Edge case: constant polynomial
	poly3 := []int{10}
	deriv3 := Derivative(poly3)
	fmt.Printf("Derivative of %v is %v\n", poly3, deriv3) // Expected: []

	// Edge case: empty input
	poly4 := []int{}
	deriv4 := Derivative(poly4)
	fmt.Printf("Derivative of %v is %v\n", poly4, deriv4) // Expected: []
}
