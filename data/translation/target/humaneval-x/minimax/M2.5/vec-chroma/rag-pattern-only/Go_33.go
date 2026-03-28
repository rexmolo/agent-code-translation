package main

import (
	"sort"
)

func SortThird(l []int) []int {
	// Create a copy of the slice
	result := make([]int, len(l))
	copy(result, l)

	// Extract elements at indices divisible by 3 (0, 3, 6, 9, ...)
	var thirdElements []int
	for i := 0; i < len(result); i += 3 {
		thirdElements = append(thirdElements, result[i])
	}

	// Sort the extracted elements
	sort.Ints(thirdElements)

	// Put sorted elements back at positions 0, 3, 6, 9, ...
	for i := 0; i < len(thirdElements); i++ {
		result[i*3] = thirdElements[i]
	}

	return result
}