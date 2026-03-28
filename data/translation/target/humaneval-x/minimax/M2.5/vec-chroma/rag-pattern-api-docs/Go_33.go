package main

import (
	"fmt"
	"slices"
)

func SortThird(l []int) []int {
	// Create a copy of the input slice
	result := make([]int, len(l))
	copy(result, l)

	// Extract elements at indices divisible by 3 (0, 3, 6, 9, ...)
	var toSort []int
	for i := 0; i < len(result); i += 3 {
		toSort = append(toSort, result[i])
	}

	// Sort the extracted elements
	slices.Sort(toSort)

	// Put sorted values back into their original positions
	for i := 0; i < len(result); i += 3 {
		result[i] = toSort[i/3]
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(SortThird([]int{1, 2, 3}))           // [1 2 3]
	fmt.Println(SortThird([]int{5, 6, 3, 4, 8, 9, 2})) // [2 6 3 4 8 9 5]
}
