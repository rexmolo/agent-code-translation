package main

import "fmt"

func Derivative(xs []int) []int {
	// Calculate derivative: for coefficient at index i with power i,
	// the derivative coefficient at index i-1 is i * coefficient
	// Skip the first element (index 0) since derivative of constant is 0
	result := make([]int, 0, len(xs)-1)
	for i, x := range xs {
		if i > 0 {
			result = append(result, i*x)
		}
	}
	return result
}

func main() {
	// Test cases
	fmt.Println(Derivative([]int{3, 1, 2, 4, 5})) // [1, 4, 12, 20]
	fmt.Println(Derivative([]int{1, 2, 3}))       // [2, 6]
}
