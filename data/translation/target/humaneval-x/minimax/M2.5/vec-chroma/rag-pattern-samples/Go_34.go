package main

import (
	"sort"
)

func Unique(l []int) []int {
	// Use a map to track unique elements (Go's equivalent of Python's set)
	seen := make(map[int]bool)
	for _, v := range l {
		seen[v] = true
	}

	// Convert map keys to a slice
	result := make([]int, 0, len(seen))
	for k := range seen {
		result = append(result, k)
	}

	// Sort the slice (equivalent to Python's sorted())
	sort.Ints(result)

	return result
}
