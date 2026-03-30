package main

import (
	"slices"
)

func Maximum(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	// Create a copy to avoid modifying the original slice
	sorted := slices.Clone(arr)
	slices.Sort(sorted)
	return sorted[len(sorted)-k:]
}