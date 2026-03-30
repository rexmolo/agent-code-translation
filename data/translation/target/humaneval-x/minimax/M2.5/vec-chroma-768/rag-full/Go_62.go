package main

func Derivative(xs []int) []int {
	// Compute derivative of polynomial: xs[0] + xs[1]*x + xs[2]*x^2 + ...
	// Derivative: xs[1] + 2*xs[2]*x + 3*xs[3]*x^2 + ...
	// In coefficient form: [xs[1], 2*xs[2], 3*xs[3], ...]
	result := make([]int, 0, len(xs)-1)
	for i, x := range xs {
		if i > 0 {
			result = append(result, i*x)
		}
	}
	return result
}