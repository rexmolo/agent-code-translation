package main

import "fmt"

func GenerateIntegers(a, b int) []int {
	// Find min and max of a, b
	minVal := a
	if b < minVal {
		minVal = b
	}
	maxVal := a
	if b > maxVal {
		maxVal = b
	}

	// Apply clamping: lower = max(2, min(a, b)), upper = min(8, max(a, b))
	lower := minVal
	if lower < 2 {
		lower = 2
	}
	if lower > 8 {
		lower = 8
	}

	upper := maxVal
	if upper > 8 {
		upper = 8
	}
	if upper < 2 {
		upper = 2
	}

	// Collect even numbers in the range [lower, upper]
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
	fmt.Println(GenerateIntegers(2, 8))   // [2 4 6 8]
	fmt.Println(GenerateIntegers(8, 2))   // [2 4 6 8]
	fmt.Println(GenerateIntegers(10, 14)) // []
}
