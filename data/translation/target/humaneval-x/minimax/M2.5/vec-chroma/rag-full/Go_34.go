package main

import (
	"slices"
)

func Unique(l []int) []int {
	// Use a map to track unique elements (set implementation)
	seen := make(map[int]struct{}, len(l))
	for _, v := range l {
		seen[v] = struct{}{}
	}

	// Extract unique keys into a slice
	unique := make([]int, 0, len(seen))
	for v := range seen {
		unique = append(unique, v)
	}

	// Sort the slice in ascending order
	slices.Sort(unique)

	return unique
}