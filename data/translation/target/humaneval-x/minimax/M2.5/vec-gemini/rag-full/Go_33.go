package main

import (
	"slices"
)

func SortThird(l []int) []int {
	// Make a copy of the input slice
	result := make([]int, len(l))
	copy(result, l)
	
	// Extract elements at indices divisible by 3 (0, 3, 6, ...)
	third := make([]int, 0)
	for i := 0; i < len(l); i += 3 {
		third = append(third, l[i])
	}
	
	// Sort the extracted elements
	slices.Sort(third)
	
	// Put sorted elements back at indices divisible by 3
	for i := 0; i < len(third); i++ {
		result[i*3] = third[i]
	}
	
	return result
}