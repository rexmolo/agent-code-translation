package main

import "sort"

func SortThird(l []int) []int {
	// Create a copy of the slice to avoid modifying the original
	result := make([]int, len(l))
	copy(result, l)

	// Extract elements at indices divisible by 3 (0, 3, 6, ...)
	var thirdIndices []int
	for i := 0; i < len(l); i += 3 {
		thirdIndices = append(thirdIndices, l[i])
	}

	// Sort the extracted elements
	sort.Ints(thirdIndices)

	// Place sorted elements back at indices divisible by 3
	for i := range thirdIndices {
		result[i*3] = thirdIndices[i]
	}

	return result
}