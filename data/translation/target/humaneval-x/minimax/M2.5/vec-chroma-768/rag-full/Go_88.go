package main

import "sort"

func SortArray(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	// Create a copy to avoid modifying the original array
	result := make([]int, len(array))
	copy(result, array)

	// Determine sort direction based on first + last element
	sum := array[0] + array[-1]

	if sum%2 == 0 {
		// Even sum → sort in descending order
		sort.Slice(result, func(i, j int) bool {
			return result[i] > result[j]
		})
	} else {
		// Odd sum → sort in ascending order
		sort.Slice(result, func(i, j int) bool {
			return result[i] < result[j]
		})
	}

	return result
}
