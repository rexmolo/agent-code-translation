package main

import "sort"

func SortArray(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	// Create a copy to not modify the original array
	result := make([]int, len(array))
	copy(result, array)

	// Check if sum of first and last is odd (ascending) or even (descending)
	ascending := (array[0]+array[len(array)-1])%2 != 0

	if ascending {
		// Sort in ascending order
		sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	} else {
		// Sort in descending order
		sort.Slice(result, func(i, j int) bool { return result[i] > result[j] })
	}

	return result
}
