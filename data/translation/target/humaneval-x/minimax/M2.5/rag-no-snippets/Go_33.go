package main

import "sort"

func SortThird(l []int) []int {
	// Create a copy of the slice
	result := make([]int, len(l))
	copy(result, l)

	// Extract elements at indices divisible by 3 (0, 3, 6, 9, ...)
	var third []int
	for i := 0; i < len(result); i += 3 {
		third = append(third, result[i])
	}

	// Sort the third elements
	sort.Ints(third)

	// Put them back at their original positions
	for i := 0; i < len(result); i += 3 {
		result[i] = third[i/3]
	}

	return result
}