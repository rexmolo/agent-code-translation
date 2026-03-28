package main

import "slices"

func Unique(l []int) []int {
	// Use map to simulate a set for uniqueness
	seen := make(map[int]struct{})
	for _, v := range l {
		seen[v] = struct{}{}
	}

	// Extract keys from the map to a slice
	result := make([]int, 0, len(seen))
	for k := range seen {
		result = append(result, k)
	}

	// Sort the slice using slices.Sort (ascending order)
	slices.Sort(result)

	return result
}
