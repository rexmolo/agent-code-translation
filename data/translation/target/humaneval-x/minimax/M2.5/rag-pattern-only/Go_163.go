package main

import (
	"fmt"
)

func GenerateIntegers(a, b int) []int {
	// Determine the lower and upper bounds, clamped between 2 and 8
	lower := max(2, min(a, b))
	upper := min(8, max(a, b))

	// Collect all even numbers in the range
	var result []int
	for i := lower; i <= upper; i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(GenerateIntegers(2, 8))  // Expected: [2 4 6 8]
	fmt.Println(GenerateIntegers(8, 2))  // Expected: [2 4 6 8]
	fmt.Println(GenerateIntegers(10, 14)) // Expected: []
}
