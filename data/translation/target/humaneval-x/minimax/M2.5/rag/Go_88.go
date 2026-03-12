package main

import "sort"

func SortArray(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	// Determine sort order based on sum of first and last element
	sum := array[0] + array[len(array)-1]
	descending := sum%2 == 0

	// Create a copy to avoid modifying the original array
	result := make([]int, len(array))
	copy(result, array)

	if descending {
		// Sort in descending order (even sum)
		sort.Slice(result, func(i, j int) bool {
			return result[i] > result[j]
		})
	} else {
		// Sort in ascending order (odd sum)
		sort.Ints(result)
	}

	return result
}
