package main

import (
	"sort"
)

func SortArray(array []int) []int {
	// Return empty slice if input is empty
	if len(array) == 0 {
		return []int{}
	}

	// Create a copy to avoid modifying the original array
	result := make([]int, len(array))
	copy(result, array)

	// Calculate sum of first and last elements
	sum := array[0] + array[len(array)-1]

	// Sort ascending if sum is odd, descending if even
	if sum%2 == 1 {
		// Odd sum → ascending order
		sort.Slice(result, func(i, j int) bool {
			return result[i] < result[j]
		})
	} else {
		// Even sum → descending order
		sort.Sort(sort.Reverse(sort.IntSlice(result)))
	}

	return result
}
