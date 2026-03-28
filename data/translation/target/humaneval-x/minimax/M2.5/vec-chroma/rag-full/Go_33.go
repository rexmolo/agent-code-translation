package main

import "sort"

func SortThird(l []int) []int {
	// Create a copy to avoid modifying the original slice
	result := make([]int, len(l))
	copy(result, l)

	// Extract elements at indices divisible by 3 (0, 3, 6, ...)
	thirdElements := make([]int, 0, (len(l)+2)/3)
	for i := 0; i < len(l); i += 3 {
		thirdElements = append(thirdElements, result[i])
	}

	// Sort the extracted elements
	sort.Ints(thirdElements)

	// Place sorted elements back at indices 0, 3, 6, ...
	for i := 0; i < len(thirdElements); i++ {
		result[i*3] = thirdElements[i]
	}

	return result
}
