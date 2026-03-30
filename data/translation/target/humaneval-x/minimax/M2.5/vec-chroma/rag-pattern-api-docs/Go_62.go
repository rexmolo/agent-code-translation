package main

func Derivative(xs []int) []int {
	// Create a slice to hold the derivative coefficients
	// Pre-allocate with len(xs)-1 since the constant term (index 0) becomes 0
	result := make([]int, 0, len(xs)-1)
	for i, x := range xs {
		// Multiply each coefficient by its index (power of x)
		// This is the derivative: d/dx(x^i) = i * x^(i-1)
		result = append(result, i*x)
	}
	// Return from index 1 onwards to skip the constant term (derivative of constant is 0)
	return result[1:]
}
