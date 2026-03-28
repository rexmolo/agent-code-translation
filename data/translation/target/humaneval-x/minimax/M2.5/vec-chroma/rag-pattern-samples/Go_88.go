package main

import "sort"

func SortArray(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}
	// Make a copy to avoid modifying the original array
	copy := make([]int, len(array))
	copy(copy, array)

	sumFirstLast := array[0] + array[-1]
	reverse := sumFirstLast%2 == 0 // true if even, false if odd

	if reverse {
		// Sort in descending order
		sort.Slice(copy, func(i, j int) bool {
			return copy[i] > copy[j]
		})
	} else {
		// Sort in ascending order
		sort.Ints(copy)
	}

	return copy
}
