package main

import "sort"

func SortArray(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	// Create a copy to avoid modifying the original array
	copyArray := make([]int, len(array))
	copy(copyArray, array)

	// If sum of first and last element is even → descending order
	// If sum of first and last element is odd → ascending order
	descending := (array[0]+array[len(array)-1])%2 == 0

	if descending {
		sort.Slice(copyArray, func(i, j int) bool {
			return copyArray[i] > copyArray[j]
		})
	} else {
		sort.Slice(copyArray, func(i, j int) bool {
			return copyArray[i] < copyArray[j]
		})
	}

	return copyArray
}
