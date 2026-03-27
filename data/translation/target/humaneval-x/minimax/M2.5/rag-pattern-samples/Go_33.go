package main

import (
	"sort"
)

func SortThird(l []int) []int {
	// Create a copy of the input slice to avoid modifying the original
	result := make([]int, len(l))
	copy(result, l)
	
	// Extract indices and values at positions divisible by 3 (0, 3, 6, ...)
	var indices []int
	for i := 0; i < len(l); i += 3 {
		indices = append(indices, i)
	}
	
	// Extract values at those indices
	values := make([]int, len(indices))
	for i, idx := range indices {
		values[i] = result[idx]
	}
	
	// Sort the extracted values
	sort.Ints(values)
	
	// Place sorted values back at their original positions
	for i, idx := range indices {
		result[idx] = values[i]
	}
	
	return result
}