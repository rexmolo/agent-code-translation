package main

import "sort"

func SortArray(array []int) []int {
	n := len(array)
	if n == 0 {
		return []int{}
	}

	// Create a copy of the array to avoid modifying the original
	result := make([]int, n)
	copy(result, array)

	// Determine sort order: odd sum -> ascending, even sum -> descending
	isEven := (array[0] + array[n-1]) % 2 == 0

	if isEven {
		// Sort in descending order
		sort.Slice(result, func(i, j int) bool {
			return result[i] > result[j]
		})
	} else {
		// Sort in ascending order
		sort.Ints(result)
	}

	return result
}
