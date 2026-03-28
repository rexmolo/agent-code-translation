package main

func Derivative(xs []int) []int {
	result := make([]int, 0, len(xs)-1)
	for i, x := range xs {
		if i > 0 {
			result = append(result, i*x)
		}
	}
	return result
}

// Example usage:
func main() {
	// Test cases
	println(Derivative([]int{3, 1, 2, 4, 5})) // [1 4 12 20]
	println(Derivative([]int{1, 2, 3}))       // [2 6]
}
