package main

import (
	"sort"
)

func Common(l1 []int, l2 []int) []int {
	// Use a map as a set to track unique common elements
	common := make(map[int]bool)

	// Find common elements by comparing all pairs
	for _, e1 := range l1 {
		for _, e2 := range l2 {
			if e1 == e2 {
				common[e1] = true
			}
		}
	}

	// Convert map keys to slice
	result := make([]int, 0, len(common))
	for k := range common {
		result = append(result, k)
	}

	// Sort the result slice
	sort.Ints(result)

	return result
}