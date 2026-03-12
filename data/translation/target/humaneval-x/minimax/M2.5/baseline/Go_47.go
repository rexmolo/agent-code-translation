package main

import (
	"sort"
)

func Median(l []int) float64 {
	// Create a copy to avoid modifying the original slice
	sorted := make([]int, len(l))
	copy(sorted, l)
	sort.Ints(sorted)

	n := len(sorted)
	if n%2 == 1 {
		return float64(sorted[n/2])
	}
	return float64(sorted[n/2-1]+sorted[n/2]) / 2.0
}
