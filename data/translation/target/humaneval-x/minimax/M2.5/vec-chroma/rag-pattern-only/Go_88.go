package main

import "sort"

func SortArray(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	// Create a copy to not modify the original array
	copyArray := make([]int, len(array))
	copy(copyArray, array)

	// Determine sort order based on sum of first and last elements
	// Ascending if sum is odd, descending if sum is even
	isEven := (array[0]+array[len(array)-1])%2 == 0

	if isEven {
		// Sort descending
		sort.Slice(copyArray, func(i, j int) bool {
			return copyArray[i] > copyArray[j]
		})
	} else {
		// Sort ascending
		sort.Ints(copyArray)
	}

	return copyArray
}
