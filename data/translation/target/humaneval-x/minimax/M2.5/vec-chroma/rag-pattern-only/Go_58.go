package main

import (
	"sort"
)

func Common(l1 []int, l2 []int) []int {
	// Use a map to track common elements (for uniqueness)
	common := make(map[int]bool)

	// Find common elements using nested loops
	for _, e1 := range l1 {
		for _, e2 := range l2 {
			if e1 == e2 {
				common[e1] = true
			}
		}
	}

	// Extract keys from map into a slice
	result := make([]int, 0, len(common))
	for k := range common {
		result = append(result, k)
	}

	// Sort the result in ascending order
	sort.Ints(result)

	return result
}