package main

import (
	"slices"
)

func SortThird(l []int) []int {
	// Make a copy to avoid mutating the original
	result := make([]int, len(l))
	copy(result, l)

	// Extract elements at indices divisible by 3 (0, 3, 6, 9, ...)
	third := make([]int, 0, (len(l)+2)/3)
	for i := 0; i < len(l); i += 3 {
		third = append(third, l[i])
	}

	// Sort the extracted elements
	slices.Sort(third)

	// Put sorted values back at indices divisible by 3
	idx := 0
	for i := 0; i < len(l); i += 3 {
		result[i] = third[idx]
		idx++
	}

	return result
}