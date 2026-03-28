package main

import (
	"sort"
)

func Maximum(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	// Make a copy to avoid modifying the original slice
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)
	return sorted[len(sorted)-k:]
}
