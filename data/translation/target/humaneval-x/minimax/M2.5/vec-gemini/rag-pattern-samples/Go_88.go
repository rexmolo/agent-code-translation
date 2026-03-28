package main

import (
	"sort"
)

func SortArray(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	// Create a copy to avoid modifying the original array
	result := make([]int, len(array))
	copy(result, array)

	// Sort descending if first + last is even, otherwise ascending
	if (array[0]+array[len(array)-1])%2 == 0 {
		sort.Slice(result, func(i, j int) bool {
			return result[i] > result[j]
		})
	} else {
		sort.Ints(result)
	}

	return result
}
