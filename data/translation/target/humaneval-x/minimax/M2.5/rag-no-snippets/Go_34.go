package main

import "sort"

func Unique(l []int) []int {
	// Use a map to track unique elements (Go doesn't have a built-in set)
	seen := make(map[int]struct{})
	for _, v := range l {
		seen[v] = struct{}{}
	}

	// Extract keys from the map into a slice
	result := make([]int, 0, len(seen))
	for k := range seen {
		result = append(result, k)
	}

	// Sort the slice of unique elements
	sort.Ints(result)

	return result
}
