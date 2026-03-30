package main

import "fmt"

// Derivative computes the derivative of a polynomial
// represented by its coefficients.
// xs[0] + xs[1] * x + xs[2] * x^2 + ...
// Returns the derivative coefficients in the same form.
// Example: derivative([3, 1, 2, 4, 5]) = [1, 4, 12, 20]
func Derivative(xs []int) []int {
	if len(xs) <= 1 {
		return []int{}
	}

	result := make([]int, 0, len(xs)-1)
	// Start from index 1 since derivative of constant (index 0) is always 0
	for i := 1; i < len(xs); i++ {
		result = append(result, i*xs[i])
	}
	return result
}

func main() {
	fmt.Println(Derivative([]int{3, 1, 2, 4, 5})) // [1, 4, 12, 20]
	fmt.Println(Derivative([]int{1, 2, 3}))       // [2, 6]
}