package main

import (
	"slices"
	"sort"
)

func Common(l1 []int, l2 []int) []int {
	// Use a map to track elements from l1 for O(n+m) efficiency
	seen := make(map[int]bool)
	for _, e := range l1 {
		seen[e] = true
	}

	// Find common elements
	var result []int
	for _, e := range l2 {
		if seen[e] {
			result = append(result, e)
		}
	}

	// Sort the result using slices.Sort (equivalent to Python's sorted())
	slices.Sort(result)

	return result
}