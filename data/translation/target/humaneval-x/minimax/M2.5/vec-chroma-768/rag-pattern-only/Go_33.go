package main

import (
	"sort"
)

func SortThird(l []int) []int {
	// Create a copy of the input slice to avoid modifying the original
	result := make([]int, len(l))
	copy(result, l)

	// Extract elements at positions divisible by 3 (0-indexed: 0, 3, 6, ...)
	var third []int
	for i := 0; i < len(l); i += 3 {
		third = append(third, l[i])
	}

	// Sort the extracted elements
	sort.Ints(third)

	// Place sorted elements back in their original positions
	for i := 0; i < len(third); i++ {
		result[i*3] = third[i]
	}

	return result
}

// Example tests (can be removed for production)
func main() {
	// Test cases from docstring
	input1 := []int{1, 2, 3}
	output1 := SortThird(input1)
	println(output1) // [1, 2, 3]

	input2 := []int{5, 6, 3, 4, 8, 9, 2}
	output2 := SortThird(input2)
	println(output2) // [2, 6, 3, 4, 8, 9, 5]
}
