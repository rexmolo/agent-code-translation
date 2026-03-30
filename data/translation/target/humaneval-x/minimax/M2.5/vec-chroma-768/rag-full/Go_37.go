package main

import (
	"fmt"
	"sort"
)

func SortEven(l []int) []int {
	// Extract elements at even indices (0, 2, 4, ...) and odd indices (1, 3, 5, ...)
	var evens []int
	var odds []int

	for i := 0; i < len(l); i++ {
		if i%2 == 0 {
			evens = append(evens, l[i])
		} else {
			odds = append(odds, l[i])
		}
	}

	// Sort the even-indexed elements
	sort.Ints(evens)

	// Interleave evens and odds back together
	var result []int
	for i := 0; i < len(evens); i++ {
		result = append(result, evens[i])
		if i < len(odds) {
			result = append(result, odds[i])
		}
	}

	return result
}

// Example usage and tests
func main() {
	// Test case 1
	input1 := []int{1, 2, 3}
	result1 := SortEven(input1)
	fmt.Printf("%v\n", result1)

	// Test case 2
	input2 := []int{5, 6, 3, 4}
	result2 := SortEven(input2)
	fmt.Printf("%v\n", result2)
}
