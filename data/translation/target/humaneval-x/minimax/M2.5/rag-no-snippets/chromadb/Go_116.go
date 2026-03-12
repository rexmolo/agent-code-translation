package main

import (
	"math/bits"
	"sort"
)

func SortArray(arr []int) []int {
	// Create a copy to avoid modifying the original slice
	result := make([]int, len(arr))
	copy(result, arr)

	// Use SliceStable to maintain relative order for elements with same number of 1s
	// This mimics Python's stable sort behavior
	sort.SliceStable(result, func(i, j int) bool {
		countI := bits.OnesCount(uint(result[i]))
		countJ := bits.OnesCount(uint(result[j]))

		// Primary sort: by number of 1s in binary representation
		if countI != countJ {
			return countI < countJ
		}
		// Secondary sort: by decimal value
		return result[i] < result[j]
	})

	return result
}